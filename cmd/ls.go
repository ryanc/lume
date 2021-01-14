package lumecmd

import (
	"flag"
)

func init() {
	var cmdName string = "ls"
	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)
	selector := fs.String("selector", defaultSelector, "Set the selector")
	fs.StringVar(selector, "s", defaultSelector, "Set the selector")

	RegisterCommand(cmdName, Command{
		Func:  LsCmd,
		Flags: fs,
		Use:   "[--selector=<selector>]",
		Short: "List the lights",
	})
}

func LsCmd(args CmdArgs) (int, error) {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		return 1, err
	}
	PrintLights(lights)
	return 0, nil
}
