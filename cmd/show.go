package lumecmd

import (
	"flag"
	"fmt"
	"strings"
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
	var b strings.Builder

	c := ctx.Client
	selector := ctx.Flags.String("selector")
	lights, err := c.ListLights(selector)

	if err != nil {
		return ExitFailure, err
	}

	sortLights(lights)

	for i, l := range lights {
		indent = 0
		fmt.Fprintf(
			&b,
			"%s Light ID: %s, %s, Power: %s\n",
			ColorizeIndicator(l.Power),
			l.Id,
			connected(l.Connected),
			ColorizePower(l.Power),
		)
		indent += Tabstop + 2
		FprintfWithIndent(&b, indent, "Label: %s, ID: %s\n", l.Label, l.Id)
		FprintfWithIndent(&b, indent, "UUID: %s\n", l.UUID)
		FprintfWithIndent(&b, indent, "Location: %s, ID: %s\n", l.Location.Name, l.Location.Id)
		FprintfWithIndent(&b, indent, "Group: %s, ID: %s\n", l.Group.Name, l.Group.Id)
		FprintfWithIndent(&b, indent, "Color: Hue: %.1f, Saturation: %.1f%%, Kelvin: %d\n",
			*l.Color.H, *l.Color.S, *l.Color.K)
		FprintfWithIndent(&b, indent, "Brightness: %.1f%%\n", l.Brightness*100)
		if l.Effect != "" {
			FprintfWithIndent(&b, indent, "Effect: %s\n", l.Effect)
		}
		FprintfWithIndent(&b, indent, "Product: %s\n", l.Product.Name)
		FprintfWithIndent(&b, indent, "Capabilities: ")
		fmt.Fprintf(&b, "Color: %s, ", YesNo(l.Product.Capabilities.HasColor))
		fmt.Fprintf(&b, "Variable Color Temp: %s, ", YesNo(l.Product.Capabilities.HasVariableColorTemp))
		fmt.Fprintf(&b, "IR: %s, ", YesNo(l.Product.Capabilities.HasIR))
		fmt.Fprintf(&b, "Chain: %s, ", YesNo(l.Product.Capabilities.HasChain))
		fmt.Fprintf(&b, "Multizone: %s, ", YesNo(l.Product.Capabilities.HasMultizone))
		fmt.Fprintf(&b, "Min Kelvin: %.1f, ", l.Product.Capabilities.MinKelvin)
		fmt.Fprintf(&b, "Max Kelvin: %.1f ", l.Product.Capabilities.MaxKelvin)
		fmt.Fprintln(&b)
		// List applicable selectors (most to least specific)
		FprintfWithIndent(&b, indent, "Selectors:\n")
		indent += Tabstop
		FprintfWithIndent(&b, indent, "id:%s\n", l.Id)
		FprintfWithIndent(&b, indent, "label:%s\n", l.Label)
		FprintfWithIndent(&b, indent, "group_id:%s\n", l.Group.Id)
		FprintfWithIndent(&b, indent, "group:%s\n", l.Group.Name)
		FprintfWithIndent(&b, indent, "location_id:%s\n", l.Location.Id)
		FprintfWithIndent(&b, indent, "location:%s\n", l.Location.Name)
		indent -= Tabstop
		FprintfWithIndent(&b, indent, "Last Seen: %s (%.1fs ago)\n", l.LastSeen, l.SecondsLastSeen)

		if i < len(lights)-1 {
			fmt.Fprintln(&b)
		}

		fmt.Print(b.String())
	}
	return ExitSuccess, nil
}

func connected(c bool) string {
	if c {
		return "Connected"
	}
	return "Disconnected"
}
