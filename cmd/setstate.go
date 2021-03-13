package lumecmd

import (
	"flag"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdSetState() Command {
	return Command{
		Name: "set-state",
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

			fs.String("format", defaultOutputFormat, "Set the output format")

			return fs
		}(),
		Use:   "[--selector <selector>] [--power (on|off)] [--color <color>] [--brightness <brightness>] [--duration <sec>] [--infrared <infrared>] [--fast]",
		Short: "Set various state attributes",
	}
}

func SetStateCmd(args CmdArgs) (int, error) {
	var p Printer

	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")
	format := args.Flags.String("format")

	if format == "" && args.Config.OutputFormat != "" {
		format = args.Config.OutputFormat
	}

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

	if power == "" && color == "" && brightnessFlag == "" && infraredFlag == "" {
		printCmdHelp(args.Name)
		return ExitFailure, nil
	}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}

	if !fast {
		p = NewPrinter(format)
		p.Results(r.Results)
	}

	return ExitSuccess, nil
}
