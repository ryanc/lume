package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	lifx "git.kill0.net/chill9/lume"
	lumecmd "git.kill0.net/chill9/lume/cmd"
	"github.com/BurntSushi/toml"

	"golang.org/x/sys/windows"
)

const lumercFile = ".lumerc"

func main() {
	var originalMode uint32
	stdout := windows.Handle(os.Stdout.Fd())

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	defer windows.SetConsoleMode(stdout, originalMode)

	var config lumecmd.Config
	config = loadConfig()

	if config.AccessToken == "" {
		config.AccessToken = os.Getenv("LIFX_ACCESS_TOKEN")
	}

	if config.AccessToken == "" {
		fmt.Println("access token is not set")
		os.Exit(1)
	}

	flag.Parse()

	command := flag.Arg(0)

	c := lifx.NewClient(config.AccessToken)

	cmdArgs := lumecmd.CmdArgs{
		Client: c,
		Config: config,
	}

	cmd, ok := lumecmd.GetCommand(command)
	if !ok {
		fmt.Printf("lume: '%s' is not lume command. See 'lume help'\n", command)
		os.Exit(1)
	}
	fs := cmd.Flags
	fs.Parse(os.Args[2:])

	cmdArgs.Flags = lumecmd.Flags{fs}
	exitCode, err := cmd.Func(cmdArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
	}
	os.Exit(exitCode)
}

func loadConfig() lumecmd.Config {
	var config lumecmd.Config
	var tryPath, configPath string

	homeDir, err := os.UserHomeDir()
	if err == nil {
		tryPath = path.Join(homeDir, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	cwd, err := os.Getwd()
	if err == nil {
		tryPath = path.Join(cwd, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	if configPath != "" {
		toml.DecodeFile(configPath, &config)
	}

	return config
}
