package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	lifx "git.kill0.net/chill9/lume"
	lumecmd "git.kill0.net/chill9/lume/cmd"
	"github.com/BurntSushi/toml"
)

const lumercFile string = ".lumerc"

func main() {
	var config lumecmd.Config

	configPath := getConfigPath()
	if configPath == "" {
		fmt.Println("fatal: ~/.lumerc was not found")
		os.Exit(1)
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Printf("fatal: failed to parse %s\n", configPath)
		fmt.Println(err)
		os.Exit(1)
	}

	envAccessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if envAccessToken != "" {
		config.AccessToken = envAccessToken
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

	cmdArgs.Flags = lumecmd.Flags{FlagSet: fs}
	exitCode, err := cmd.Func(cmdArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
	}
	os.Exit(exitCode)
}

func getConfigPath() string {
	var tryPath, configPath string

	// ~/.lumerc
	homeDir, err := os.UserHomeDir()
	if err == nil {
		tryPath = path.Join(homeDir, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	// ./.lumerc
	cwd, err := os.Getwd()
	if err == nil {
		tryPath = path.Join(cwd, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	return configPath
}
