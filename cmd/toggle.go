package lumecmd

import (
	"flag"
)

func NewCmdToggle() Command {
	return Command{
		Name: "toggle",
		Func: ToggleCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("toggle", flag.ExitOnError)

			duration := fs.Float64("duration", defaultDuration, "Set the duration")
			fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			fs.String("format", defaultOutputFormat, "Set the output format")

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Toggle the power on/off",
	}
}

func ToggleCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	duration := ctx.Flags.Float64("duration")
	selector := ctx.Flags.String("selector")
	format := ctx.Flags.String("format")

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	r, err := c.Toggle(selector, duration)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Results(r.Results)

	return ExitSuccess, nil
}
