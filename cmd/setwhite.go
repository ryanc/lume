package lumecmd

import (
	"flag"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdSetWhite() Command {
	return Command{
		Name: "set-white",
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
	}
}

func SetWhiteCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	state := lifx.State{}
	selector := ctx.Flags.String("selector")
	format, err := getOutputFormatFromFlags(ctx.Flags)
	if err != nil {
		return ExitFailure, err
	}

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	power := ctx.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	kelvinFlag := ctx.Flags.String("kelvin")
	if kelvinFlag != "" {
		kelvin := ctx.Flags.Int16("kelvin")
		color, err := lifx.NewWhite(kelvin)
		if err != nil {
			return ExitFailure, err
		}
		state.Color = color
	}

	name := ctx.Flags.String("name")
	if name != "" {
		name := ctx.Flags.String("name")
		color, err := lifx.NewWhiteString(name)
		if err != nil {
			return ExitFailure, err
		}
		state.Color = color
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

	if power == "" && kelvinFlag == "" && name == "" && brightnessFlag == "" && infraredFlag == "" {
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
