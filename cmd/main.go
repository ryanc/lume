package lumecmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"git.kill0.net/chill9/lifx-go"
	"github.com/BurntSushi/toml"
)

var userAgent string

func init() {
	userAgent = initUserAgent()

	RegisterCommand("help", Command{
		Func: HelpCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("help", flag.ExitOnError)

			return fs
		}(),
		Use:   "<command>",
		Short: "Show help for a command",
	})
	RegisterCommand("ls", Command{
		Func: LsCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("ls", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector=<selector>]",
		Short: "List the lights",
	})
	RegisterCommand("poweroff", Command{
		Func: PoweroffCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("poweroff", flag.ExitOnError)

			duration := fs.Float64("duration", defaultDuration, "Set the duration")
			fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Power on",
	})
	RegisterCommand("poweron", Command{
		Func: PoweronCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("poweron", flag.ExitOnError)

			duration := fs.Float64("duration", defaultDuration, "Set the duration")
			fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Power on",
	})
	RegisterCommand("set-color", Command{
		Func: SetColorCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("set-color", flag.ExitOnError)

			selector := fs.String("selector", "all", "the selector")
			fs.StringVar(selector, "s", "all", "the selector")

			power := fs.String("power", defaultPower, "power state")
			fs.StringVar(power, "p", defaultPower, "power state")

			hue := fs.String("hue", defaultHue, "hue level")
			fs.StringVar(hue, "H", defaultHue, "hue level")

			saturation := fs.String("saturation", defaultSaturation, "saturation level")
			fs.StringVar(saturation, "S", defaultSaturation, "saturation level")

			rgb := fs.String("rgb", defaultRGB, "RGB value")
			fs.StringVar(rgb, "r", defaultRGB, "RGB value")

			name := fs.String("name", defaultName, "named color")
			fs.StringVar(name, "n", defaultName, "named color")

			brightness := fs.String("brightness", defaultBrightness, "brightness state")
			fs.StringVar(brightness, "b", defaultBrightness, "brightness state")

			duration := fs.Float64("duration", defaultDuration, "duration state")
			fs.Float64Var(duration, "d", defaultDuration, "duration state")

			fast := fs.Bool("fast", defaultFast, "fast state")
			fs.BoolVar(fast, "f", defaultFast, "fast state")

			return fs
		}(),
		Use:   "[--selector <selector>] [--power (on|off)] [--hue <hue>] [--saturation <saturation>] [--rgb <rbg>] [--name <color>] [--brightness <brightness>] [--duration <sec>] [--fast]",
		Short: "Set the color",
	})
	RegisterCommand("set-state", Command{
		Func: SetStateCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("set-state", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			power := fs.String("power", defaultPower, "power state")
			fs.StringVar(power, "p", defaultPower, "power state")

			color := fs.String("color", defaultColor, "color state")
			fs.StringVar(color, "c", defaultColor, "color state")

			brightness := fs.String("brightness", defaultBrightness, "brightness state")
			fs.StringVar(brightness, "b", defaultBrightness, "brightness state")

			duration := fs.Float64("duration", defaultDuration, "duration state")
			fs.Float64Var(duration, "d", defaultDuration, "duration state")

			infrared := fs.String("infrared", defaultInfrared, "infrared state")
			fs.StringVar(infrared, "i", defaultInfrared, "infrared state")

			fast := fs.Bool("fast", defaultFast, "fast state")
			fs.BoolVar(fast, "f", defaultFast, "fast state")

			return fs
		}(),
		Use:   "[--selector <selector>] [--power (on|off)] [--color <color>] [--brightness <brightness>] [--duration <sec>] [--infrared <infrared>] [--fast]",
		Short: "Set various state attributes",
	})
	RegisterCommand("set-white", Command{
		Func: SetWhiteCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("set-white", flag.ExitOnError)

			selector := fs.String("selector", "all", "the selector")
			fs.StringVar(selector, "s", "all", "the selector")

			power := fs.String("power", defaultPower, "power state")
			fs.StringVar(power, "p", defaultPower, "power state")

			kelvin := fs.String("kelvin", defaultWhiteKelvin, "kelvin level")
			fs.StringVar(kelvin, "k", defaultWhiteKelvin, "kelvin level")

			name := fs.String("name", defaultWhiteName, "named white level")
			fs.StringVar(name, "n", defaultWhiteName, "named white level")

			brightness := fs.String("brightness", defaultBrightness, "brightness state")
			fs.StringVar(brightness, "b", defaultBrightness, "brightness state")

			duration := fs.Float64("duration", defaultDuration, "duration state")
			fs.Float64Var(duration, "d", defaultDuration, "duration state")

			infrared := fs.String("infrared", defaultInfrared, "infrared state")
			fs.StringVar(infrared, "i", defaultInfrared, "infrared state")

			fast := fs.Bool("fast", defaultFast, "fast state")
			fs.BoolVar(fast, "f", defaultFast, "fast state")

			return fs
		}(),
		Use:   "[--selector <selector>] [--power (on|off)] [--kelvin <kelvin>] [--name <color>] [--brightness <brightness>] [--duration <sec>] [--infrared] [--fast]",
		Short: "Set the white level",
	})
	RegisterCommand("toggle", Command{
		Func: ToggleCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("toggle", flag.ExitOnError)

			duration := fs.Float64("duration", defaultDuration, "Set the duration")
			fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Toggle the power on/off",
	})
}

const lumercFile string = ".lumerc"

func Main(args []string) (int, error) {
	var config Config
	var err error

	if len(args) == 1 {
		args = append(args, "help")
	}

	configPath := getConfigPath()
	if configPath == "" {
		err = errors.New("fatal: ~/.lumerc was not found")
		return ExitFailure, err
	}

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		err = fmt.Errorf("fatal: failed to parse %s; %w", configPath, err)
		return ExitFailure, err
	}

	envAccessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if envAccessToken != "" {
		config.AccessToken = envAccessToken
	}

	if err = config.Validate(); err != nil {
		return ExitFailure, fmt.Errorf("fatal: %s", err)
	}

	flag.Parse()

	command := args[1]

	c := lifx.NewClient(
		config.AccessToken,
		lifx.WithUserAgent(userAgent),
	)

	cmdArgs := CmdArgs{
		Client: c,
		Config: config,
	}

	cmd, ok := GetCommand(command)
	if !ok {
		err = fmt.Errorf("lume: '%s' is not lume command. See 'lume help'", command)
		return ExitFailure, err
	}
	fs := cmd.Flags
	fs.Parse(args[2:])

	cmdArgs.Flags = Flags{FlagSet: fs}
	cmdArgs.Name = command
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

func initUserAgent() string {
	var b strings.Builder

	b.WriteString("lume")
	b.WriteRune('/')
	b.WriteString(Version)
	return b.String()
}
