package lumecmd

import (
	"flag"
	"fmt"
)

var (
	idWidth, locationWidth, groupWidth, labelWidth, lastSeenWidth, powerWidth int
)

func init() {
	fs := flag.NewFlagSet("toggle", flag.ExitOnError)
	selector := fs.String("selector", "all", "Set the selector")
	fs.StringVar(selector, "s", "all", "Set the selector")

	RegisterCommand("ls", Command{
		Func:  LsCmd,
		Flags: fs,
	})
}

func LsCmd(args CmdArgs) int {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	PrintLights(lights)
	return 0
}
