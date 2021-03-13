package lumecmd

import (
	"flag"

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

			fs.String("format", defaultOutputFormat, "Set the output format")

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Power on",
	}
}

func PoweroffCmd(args CmdArgs) (int, error) {
	var p Printer

	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	format := args.Flags.String("format")
	state := lifx.State{Power: "off", Duration: duration}

	if format == "" && args.Config.OutputFormat != "" {
		format = args.Config.OutputFormat
	}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Results(r.Results)

	return ExitSuccess, nil
}
