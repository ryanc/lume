package lumecmd

import (
	"flag"
	"fmt"
)

func init() {
	var cmdName string = "set-state"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)
	duration := fs.Float64("duration", 1.0, "Set the duration")
	fs.Float64Var(duration, "d", 1.0, "Set the duration")
	selector := fs.String("selector", "all", "Set the selector")
	fs.StringVar(selector, "s", "all", "Set the selector")

	RegisterCommand(cmdName, Command{
		Func:  ToggleCmd,
		Flags: fs,
	})
}

func ToggleCmd(args CmdArgs) int {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	r, err := c.Toggle(selector, duration)
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		return 1
	}
	PrintResults(r.Results)
	return 0
}
