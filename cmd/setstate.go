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

func SetStateCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	state := lifx.State{}
	selector := ctx.Flags.String("selector")
	format := ctx.Flags.String("format")

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	power := ctx.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	color := ctx.Flags.String("color")
	if color != "" {
		state.Color = lifx.NamedColor(color)
	}

	brightnessFlag := ctx.Flags.String("brightness")
	if brightnessFlag != "" {
		brightness := ctx.Flags.Float64("brightness")
		state.Brightness = brightness
	}

	duration := ctx.Flags.Float64("duration")
	state.Duration = duration

	infraredFlag := ctx.Flags.String("infrared")
	if infraredFlag != "" {
		infrared := ctx.Flags.Float64("infrared")
		state.Infrared = infrared
	}

	fast := ctx.Flags.Bool("fast")
	state.Fast = fast

	if power == "" && color == "" && brightnessFlag == "" && infraredFlag == "" {
		printCmdHelp(ctx.Name)
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
