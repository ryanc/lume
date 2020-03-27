package lumecmd

import (
	"flag"
	"fmt"
)

func init() {
	fs := flag.NewFlagSet("toggle", flag.ExitOnError)
	duration := fs.Float64("duration", 1.0, "Set the duration")
	fs.Float64Var(duration, "d", 1.0, "Set the duration")
	selector := fs.String("selector", "all", "Set the selector")
	fs.StringVar(selector, "s", "all", "Set the selector")

	RegisterCommand("toggle", Command{
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
		fmt.Println(err)
		return 1
	}
	PrintResults(r.Results)
	return 0
}
