package lumecmd

import (
	"flag"
	"fmt"
)

var (
	idWidth, locationWidth, groupWidth, labelWidth, lastSeenWidth, powerWidth int
)

func init() {
	var cmdName string = "ls"
	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)
	selector := fs.String("selector", defaultSelector, "Set the selector")
	fs.StringVar(selector, "s", defaultSelector, "Set the selector")

	RegisterCommand(cmdName, Command{
		Func:  LsCmd,
		Flags: fs,
	})
}

func LsCmd(args CmdArgs) int {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		return 1
	}
	PrintLights(lights)
	return 0
}
