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
	homeDir, err := os.UserHomeDir()
	_, err = toml.DecodeFile(path.Join(homeDir, lumercFile), &config)
	if os.IsNotExist(err) {
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
		fmt.Printf("lume: '%s' is not lume command. See 'lume' --help.'\n", command)
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
