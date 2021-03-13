package lumecmd

import (
	"flag"
)

func NewCmdLs() Command {
	return Command{
		Name: "ls",
		Func: LsCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("ls", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			fs.String("format", defaultOutputFormat, "Set the output format")

			return fs
		}(),
		Use:   "[--selector=<selector>]",
		Short: "List the lights",
	}
}

func LsCmd(args CmdArgs) (int, error) {
	var p Printer

	c := args.Client
	selector := args.Flags.String("selector")
	format := args.Flags.String("format")

	if format == "" && args.Config.OutputFormat != "" {
		format = args.Config.OutputFormat
	}

	lights, err := c.ListLights(selector)
	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Lights(lights)

	return ExitSuccess, nil
}
