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

			return fs
		}(),
		Use:   "[--selector=<selector>]",
		Short: "List the lights",
	}
}

func LsCmd(ctx Context) (int, error) {
	var p Printer

	c := ctx.Client
	selector := ctx.Flags.String("selector")
	format := ctx.Flags.String("output-format")

	if format == "" && ctx.Config.OutputFormat != "" {
		format = ctx.Config.OutputFormat
	}

	lights, err := c.ListLights(selector)

	Debugf("%+v\n", lights)

	if err != nil {
		return ExitFailure, err
	}

	p = NewPrinter(format)
	p.Lights(lights)

	return ExitSuccess, nil
}
