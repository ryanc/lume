package lumecmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"

	lifx "git.kill0.net/chill9/lume"
	"github.com/BurntSushi/toml"
)

const lumercFile string = ".lumerc"

func Main(args []string) (int, error) {
	var config Config
	var err error

	configPath := getConfigPath()
	if configPath == "" {
		err = errors.New("fatal: ~/.lumerc was not found")
		return ExitError, err
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		err = fmt.Errorf("fatal: failed to parse %s", configPath)
		return ExitError, err
	}

	envAccessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if envAccessToken != "" {
		config.AccessToken = envAccessToken
	}

	if err = config.Validate(); err != nil {
		return ExitError, fmt.Errorf("fatal: %s", err)
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
		err = fmt.Errorf("lume: '%s' is not lume command. See 'lume help'", command)
		return ExitError, err
	}
	fs := cmd.Flags
	fs.Parse(args[2:])

	cmdArgs.Flags = Flags{FlagSet: fs}
	exitCode, err := cmd.Func(cmdArgs)
	if err != nil {
		err = fmt.Errorf("fatal: %s", err)
	}

	return exitCode, err
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
