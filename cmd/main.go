package lumecmd

import (
	"flag"
	"fmt"
	"os"
	"path"

	lifx "git.kill0.net/chill9/lume"
	"github.com/BurntSushi/toml"
)

const lumercFile string = ".lumerc"

func Main(args []string) int {
	var config Config

	configPath := getConfigPath()
	if configPath == "" {
		fmt.Println("fatal: ~/.lumerc was not found")
		return 1
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		fmt.Printf("fatal: failed to parse %s\n", configPath)
		fmt.Println(err)
		return 1
	}

	envAccessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if envAccessToken != "" {
		config.AccessToken = envAccessToken
	}

	if config.AccessToken == "" {
		fmt.Println("access token is not set")
		return 1
	}

	flag.Parse()

	command := flag.Arg(0)

	c := lifx.NewClient(config.AccessToken)

	cmdArgs := CmdArgs{
		Client: c,
		Config: config,
	}

	cmd, ok := GetCommand(command)
	if !ok {
		fmt.Printf("lume: '%s' is not lume command. See 'lume help'\n", command)
		return 1
	}
	fs := cmd.Flags
	fs.Parse(args[2:])

	cmdArgs.Flags = Flags{FlagSet: fs}
	exitCode, err := cmd.Func(cmdArgs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err)
	}
	return exitCode
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
