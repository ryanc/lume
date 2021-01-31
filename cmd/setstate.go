package lumecmd

import (
	"flag"

	lifx "git.kill0.net/chill9/lume"
)

func init() {
	var cmdName string = "set-state"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

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

	RegisterCommand(cmdName, Command{
		Func:  SetStateCmd,
		Flags: fs,
		Use:   "[--selector <selector>] [--power (on|off)] [--color <color>] [--brightness <brightness>] [--duration <sec>] [--infrared <infrared>] [--fast]",
		Short: "Set various state attributes",
	})
}

func SetStateCmd(args CmdArgs) (int, error) {
	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")

	power := args.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	color := args.Flags.String("color")
	if color != "" {
		state.Color = lifx.NamedColor(color)
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
		return ExitFailure, err
	}

	if !fast {
		PrintResults(r.Results)
	}

	return ExitSuccess, nil
}
