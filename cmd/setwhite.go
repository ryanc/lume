package lumecmd

import (
	"flag"

	lifx "git.kill0.net/chill9/lume"
)

func init() {
	var cmdName string = "set-white"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

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

	RegisterCommand(cmdName, Command{
		Func:  SetWhiteCmd,
		Flags: fs,
		Use:   "[--selector <selector>] [--power (on|off)] [--kelvin <kelvin>] [--name <color>] [--brightness <brightness>] [--duration <sec>] [--infrared] [--fast]",
		Short: "Set the white level",
	})
}

func SetWhiteCmd(args CmdArgs) (int, error) {
	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")

	power := args.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	kelvinFlag := args.Flags.String("kelvin")
	if kelvinFlag != "" {
		kelvin := args.Flags.Int16("kelvin")
		color, err := lifx.NewWhite(kelvin)
		if err != nil {
			return ExitError, err
		}
		state.Color = color
	}

	name := args.Flags.String("name")
	if name != "" {
		name := args.Flags.String("name")
		color, err := lifx.NewWhiteString(name)
		if err != nil {
			return ExitError, err
		}
		state.Color = color
	}

	brightnessFlag := args.Flags.String("brightness")
	if brightnessFlag != "" {
		brightness := args.Flags.Float64("brightness")
		state.Brightness = brightness
	}

	duration := args.Flags.Float64("duration")
	state.Duration = duration

	infraredFlag := args.Flags.String("infrared")
	if infraredFlag != "" {
		infrared := args.Flags.Float64("infrared")
		state.Infrared = infrared
	}

	fast := args.Flags.Bool("fast")
	state.Fast = fast

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitError, err
	}

	if !fast {
		PrintResults(r.Results)
	}

	return ExitSuccess, nil
}
