package lumecmd

import (
	"flag"
	"fmt"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdPoweroff() Command {
	return Command{
		Name: "poweroff",
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
	}
}

func PoweroffCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	duration := ctx.Flags.Float64("duration")
	selector := ctx.Flags.String("selector")
	state := lifx.State{Power: "off", Duration: duration}
	format, err := getOutputFormatFromFlags(ctx.Flags)
	if err != nil {
		return ExitFailure, err
	}

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	fmt.Print(p.Results(r.Results))

	return ExitSuccess, nil
}
