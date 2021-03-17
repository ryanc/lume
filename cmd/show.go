package lumecmd

import (
	"flag"
	"fmt"
)

const Tabstop int = 2

func NewCmdShow() Command {
	return Command{
		Name: "show",
		Func: ShowCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("show", flag.ExitOnError)

			selector := fs.String("selector", defaultSelector, "Set the selector")
			fs.StringVar(selector, "s", defaultSelector, "Set the selector")

			return fs
		}(),
		Use:   "[--selector=<selector>]",
		Short: "Show details about the lights",
	}
}

func ShowCmd(ctx Context) (int, error) {
	var indent int
	c := ctx.Client
	selector := ctx.Flags.String("selector")
	lights, err := c.ListLights(selector)

	if err != nil {
		return ExitFailure, err
	}

	sortLights(lights)

	for i, l := range lights {
		indent = 0
		fmt.Printf(
			"Light ID: %s, %s, Power: %s\n",
			l.Id,
			connected(l.Connected),
			ColorizePower(l.Power),
		)
		indent += Tabstop
		PrintfWithIndent(indent, "Label: %s, ID: %s\n", l.Label, l.Id)
		PrintfWithIndent(indent, "UUID: %s\n", l.UUID)
		PrintfWithIndent(indent, "Location: %s, ID: %s\n", l.Location.Name, l.Location.Id)
		PrintfWithIndent(indent, "Group: %s, ID: %s\n", l.Group.Name, l.Group.Id)
		PrintfWithIndent(indent, "Color: Hue: %.1f, Saturation: %.1f%%, Kelvin: %d\n",
			*l.Color.H, *l.Color.S, *l.Color.K)
		PrintfWithIndent(indent, "Brightness: %.1f%%\n", l.Brightness*100)
		if l.Effect != "" {
			PrintfWithIndent(indent, "Effect: %s\n", l.Effect)
		}
		PrintfWithIndent(indent, "Product: %s\n", l.Product.Name)
		PrintfWithIndent(indent, "Capabilities: ")
		fmt.Printf("Color: %s, ", YesNo(l.Product.Capabilities.HasColor))
		fmt.Printf("Variable Color Temp: %s, ", YesNo(l.Product.Capabilities.HasVariableColorTemp))
		fmt.Printf("IR: %s, ", YesNo(l.Product.Capabilities.HasIR))
		fmt.Printf("Chain: %s, ", YesNo(l.Product.Capabilities.HasChain))
		fmt.Printf("Multizone: %s, ", YesNo(l.Product.Capabilities.HasMultizone))
		fmt.Printf("Min Kelvin: %.1f, ", l.Product.Capabilities.MinKelvin)
		fmt.Printf("Max Kelvin: %.1f ", l.Product.Capabilities.MaxKelvin)
		fmt.Println()
		// List applicable selectors (most to least specific)
		PrintfWithIndent(indent, "Selectors:\n")
		indent += Tabstop
		PrintfWithIndent(indent, "id:%s\n", l.Id)
		PrintfWithIndent(indent, "label:%s\n", l.Label)
		PrintfWithIndent(indent, "group_id:%s\n", l.Group.Id)
		PrintfWithIndent(indent, "group:%s\n", l.Group.Name)
		PrintfWithIndent(indent, "location_id:%s\n", l.Location.Id)
		PrintfWithIndent(indent, "location:%s\n", l.Location.Name)
		indent -= Tabstop
		PrintfWithIndent(indent, "Last Seen: %s (%.1fs ago)\n", l.LastSeen, l.SecondsLastSeen)

		if i < len(lights)-1 {
			fmt.Println()
		}
	}
	return ExitSuccess, nil
}

func connected(c bool) string {
	if c {
		return "Connected"
	}
	return "Disconnected"
}
