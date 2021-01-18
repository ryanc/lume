package lumecmd

import (
	"flag"
)

func init() {
	var cmdName string = "toggle"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	duration := fs.Float64("duration", defaultDuration, "Set the duration")
	fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

	selector := fs.String("selector", defaultSelector, "Set the selector")
	fs.StringVar(selector, "s", defaultSelector, "Set the selector")

	RegisterCommand(cmdName, Command{
		Func:  ToggleCmd,
		Flags: fs,
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Toggle the power on/off",
	})
}

func ToggleCmd(args CmdArgs) (int, error) {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	r, err := c.Toggle(selector, duration)
	if err != nil {
		return ExitError, err
	}
	PrintResults(r.Results)
	return ExitSuccess, nil
}
