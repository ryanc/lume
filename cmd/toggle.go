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

			return fs
		}(),
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Toggle the power on/off",
	}
}

func ToggleCmd(args CmdArgs) (int, error) {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	r, err := c.Toggle(selector, duration)
	if err != nil {
		return ExitFailure, err
	}
	PrintResults(r.Results)
	return ExitSuccess, nil
}
