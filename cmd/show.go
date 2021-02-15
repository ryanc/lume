package lumecmd

import "fmt"

func ShowCmd(args CmdArgs) (int, error) {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)

	if err != nil {
		return ExitFailure, err
	}

	sortLights(lights)

	for _, l := range lights {
		fmt.Printf(
			"Light Id: %s, Label: %s, %s, Power: %s\n",
			l.Id,
			l.Label,
			connected(l.Connected),
			powerColor(l.Power),
		)
		fmt.Printf("  Label: %s\n", l.Label)
		fmt.Printf("  UUID: %s\n", l.UUID)
		fmt.Printf("  Location: %s, ID: %s\n", l.Location.Name, l.Location.Id)
		fmt.Printf("  Group: %s, ID: %s\n", l.Group.Name, l.Group.Id)
		fmt.Printf("  Color: Hue: %.1f, Saturation: %.1f%%, Kelvin: %d\n",
			*l.Color.H, *l.Color.S, *l.Color.K)
		fmt.Printf("  Brightness: %.1f%%\n", l.Brightness*100)
		if l.Effect != "" {
			fmt.Printf("  Effect: %s\n", l.Effect)
		}
		fmt.Printf("  Product: %s\n", l.Product.Name)
		fmt.Printf("  Capabilities: ")
		fmt.Printf("Color: %s, ", YesNo(l.Product.Capabilities.HasColor))
		fmt.Printf("Variable Color Temp: %s, ", YesNo(l.Product.Capabilities.HasVariableColorTemp))
		fmt.Printf("IR: %s, ", YesNo(l.Product.Capabilities.HasIR))
		fmt.Printf("Chain: %s, ", YesNo(l.Product.Capabilities.HasChain))
		fmt.Printf("Multizone: %s, ", YesNo(l.Product.Capabilities.HasMultizone))
		fmt.Printf("Min Kelvin: %.1f, ", l.Product.Capabilities.MinKelvin)
		fmt.Printf("Max Kelvin: %.1f ", l.Product.Capabilities.MaxKelvin)
		fmt.Println()
		// List applicable selectors (most to least specific)
		fmt.Printf("  Selectors:\n")
		fmt.Printf("    id:%s\n", l.Id)
		fmt.Printf("    label:%s\n", l.Label)
		fmt.Printf("    group_id:%s\n", l.Group.Id)
		fmt.Printf("    group:%s\n", l.Group.Name)
		fmt.Printf("    location_id:%s\n", l.Location.Id)
		fmt.Printf("    location:%s\n", l.Location.Name)
		fmt.Printf("  Last Seen: %s (%.1fs ago)\n", l.LastSeen, l.SecondsLastSeen)
		fmt.Println()
	}
	return ExitSuccess, nil
}

func connected(c bool) string {
	if c {
		return "Connected"
	}
	return "Disconnected"
}
