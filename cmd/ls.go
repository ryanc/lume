package lumecmd

import "flag"

func NewCmdLs() Command {
	return Command{
		Name: "ls",
		Func: LsCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("ls", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector=<selector>]",
		Short: "List the lights",
	}
}

func LsCmd(args CmdArgs) (int, error) {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		return ExitFailure, err
	}
	PrintLights(lights)
	return ExitSuccess, nil
}
