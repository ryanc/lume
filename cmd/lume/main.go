package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"git.kill0.net/chill9/lume"
	lumecmd "git.kill0.net/chill9/lume/cmd"
	"github.com/BurntSushi/toml"
)

const lumercFile = ".lumerc"

type Config struct {
	AccessToken string
}

func main() {
	var config Config
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
	}

	cmd, ok := lumecmd.GetCommand(command)
	if !ok {
		fmt.Printf("lume: '%s' is not lume command. See 'lume' --help.'\n", command)
		os.Exit(1)
	}
	fs := cmd.Flags
	fs.Parse(os.Args[2:])

	cmdArgs.Flags = lumecmd.Flags{fs}
	os.Exit(cmd.Func(cmdArgs))
}
