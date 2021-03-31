package lumecmd

import (
	"flag"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdPoweron() Command {
	return Command{
		Name: "poweron",
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
	}
}

func PoweronCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	duration := ctx.Flags.Float64("duration")
	selector := ctx.Flags.String("selector")
	format := ctx.Flags.String("output-format")
	state := lifx.State{Power: "on", Duration: duration}

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Results(r.Results)

	return ExitSuccess, nil
}
