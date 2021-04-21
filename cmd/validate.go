package lumecmd

import (
	"errors"
	"flag"
	"fmt"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdValidate() Command {
	return Command{
		Name: "validate",
		Func: ValidateCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("validate", flag.ExitOnError)

			return fs
		}(),
		Use:   "<command>",
		Short: "Validate a color string",
	}
}

func ValidateCmd(ctx Context) (int, error) {
	c := ctx.Client

	if len(ctx.Args) != 1 {
		printCmdHelp(ctx.Name)
		return ExitFailure, nil
	}

	color := lifx.NamedColor(ctx.Args[0])

	i, err := c.ValidateColor(color)
	if err != nil {
		return ExitFailure, err
	}

	if validColor, ok := i.(*lifx.HSBKColor); ok {
		fmt.Print(validColor)
	} else {
		return ExitFailure, errors.New("go type %T but wanted *HSBKColor")
	}
	fmt.Println()

	return ExitSuccess, nil
}
