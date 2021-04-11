package lumecmd

import (
	"flag"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdBreathe() Command {
	return Command{
		Name: "breathe",
		Func: BreatheCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("breathe", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			fs.String("color", defaultColor, "The color to use for the breathe effect")

			fs.String("from-color", defaultColor, "The color to start the effect from")

			fs.Float64("period", lifx.DefaultBreathePeriod, "The time in seconds for one cycle of the effect")

			fs.Float64("cycles", lifx.DefaultBreatheCycles, "The number of times to repeat the effect")

			fs.Bool("persist", lifx.DefaultBreathePersist, "If false set the light back to its previous value when effect ends, if true leave the last effect color")

			fs.Bool("power-on", lifx.DefaultBreathePowerOn, "If true, turn the bulb on if it is not already on")

			fs.Float64("peak", lifx.DefaultBreathePeak, "Defines where in a period the target color is at its maximum (min: 0.0, max: 1.0)")

			return fs
		}(),
		Use:   "[--selector <selector>] --color <color> [--from-color <color>] [--period <period>] [--cycles <cycles>] [--persist <persist>] [--power-on] [--peak <peak>]",
		Short: "The breathe effect",
	}
}

func BreatheCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	breathe := lifx.NewBreathe()
	selector := ctx.Flags.String("selector")
	format, err := getOutputFormatFromFlags(ctx.Flags)
	if err != nil {
		return ExitFailure, err
	}

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	color := ctx.Flags.String("color")
	if color != "" {
		breathe.Color = lifx.NamedColor(color)
	}

	from_color := ctx.Flags.String("from-color")
	if from_color != "" {
		breathe.FromColor = lifx.NamedColor(from_color)
	}

	periodFlag := ctx.Flags.String("period")
	if periodFlag != "" {
		period := ctx.Flags.Float64("period")
		breathe.Period = period
	}

	cyclesFlag := ctx.Flags.String("cycles")
	if cyclesFlag != "" {
		cycles := ctx.Flags.Float64("cycles")
		breathe.Cycles = cycles
	}

	persist := ctx.Flags.Bool("persist")
	breathe.Persist = persist

	power_on := ctx.Flags.Bool("power-on")
	breathe.PowerOn = power_on

	peakFlag := ctx.Flags.String("peak")
	if peakFlag != "" {
		peak := ctx.Flags.Float64("peak")
		breathe.Peak = peak
	}

	if color == "" {
		printCmdHelp(ctx.Name)
		return ExitFailure, nil
	}

	if err := breathe.Valid(); err != nil {
		return ExitFailure, err
	}

	r, err := c.Breathe(selector, breathe)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Results(r.Results)

	return ExitSuccess, nil
}
