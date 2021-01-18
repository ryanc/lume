package lumecmd

import (
	"flag"

	lifx "git.kill0.net/chill9/lume"
)

func init() {
	var cmdName string = "poweron"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	duration := fs.Float64("duration", defaultDuration, "Set the duration")
	fs.Float64Var(duration, "d", defaultDuration, "Set the duration")

	selector := fs.String("selector", defaultSelector, "Set the selector")
	fs.StringVar(selector, "s", defaultSelector, "Set the selector")

	RegisterCommand(cmdName, Command{
		Func:  PoweronCmd,
		Flags: fs,
		Use:   "[--selector <selector>] [--duration <sec>]",
		Short: "Power on",
	})
}

func PoweronCmd(args CmdArgs) (int, error) {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	state := lifx.State{Power: "on", Duration: duration}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitError, err
	}
	PrintResults(r.Results)
	return ExitSuccess, nil
}
