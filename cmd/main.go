package lumecmd

import (
	"errors"
	"flag"
	"fmt"

	"git.kill0.net/chill9/lifx-go"
)

func init() {
	RegisterCommand(NewCmdHelp())
	RegisterCommand(NewCmdLs())
	RegisterCommand(NewCmdPoweroff())
	RegisterCommand(NewCmdPoweron())
	RegisterCommand(NewCmdSetColor())
	RegisterCommand(NewCmdSetState())
	RegisterCommand(NewCmdSetWhite())
	RegisterCommand(NewCmdShow())
	RegisterCommand(NewCmdToggle())
	RegisterCommand(NewCmdVersion())
	RegisterCommand(NewCmdBreathe())
	RegisterCommand(NewCmdValidate())
}

var Version string
var BuildDate string
var GitCommit string

func Main(args []string) (int, error) {
	var config *Config = GetConfig()
	var err error

	flag.Parse()

	if len(args) == 1 {
		args = append(args, "help")
	}

	configPath := getConfigPath()
	if configPath == "" {
		err = errors.New("fatal: ~/.lumerc was not found")
		return ExitFailure, err
	}

	if config, err = LoadConfigFile(configPath); err != nil {
		return ExitFailure, err
	}
	config.MergeWithEnv()

	if err = config.Validate(); err != nil {
		return ExitFailure, fmt.Errorf("fatal: %s", err)
	}

	command := args[1]

	c := lifx.NewClient(
		config.AccessToken,
		lifx.WithUserAgent(config.userAgent),
		lifx.WithDebug(config.Debug),
	)

	ctx := Context{
		Client: c,
		Config: *config,
		Args:   args[2:],
	}

	cmd, ok := GetCommand(command)
	if !ok {
		err = fmt.Errorf("lume: '%s' is not lume command. See 'lume help'", command)
		return ExitFailure, err
	}

	fs := cmd.Flags
	if fs != nil {
		fs.Parse(args[2:])
		ctx.Flags = Flags{FlagSet: fs}
	}
	ctx.Name = command

	exitCode, err := cmd.Func(ctx)
	if err != nil {
		err = fmt.Errorf("fatal: %s", err)
	}

	return exitCode, err
}
