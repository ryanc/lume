package lumecmd

import (
	"fmt"

	"git.kill0.net/chill9/lifx-go"
)

func SetColorCmd(args CmdArgs) (int, error) {
	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")

	power := args.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	hueFlag := args.Flags.String("hue")
	saturationFlag := args.Flags.String("saturation")
	rgbFlag := args.Flags.String("rgb")
	name := args.Flags.String("name")

	if (hueFlag == "" || saturationFlag == "") && rgbFlag == "" && name == "" {
		printCmdHelp(args.Name)
		return ExitFailure, nil
	}

	if hueFlag != "" || saturationFlag != "" {
		color := lifx.HSBKColor{}

		if hueFlag != "" {
			hue := args.Flags.Float32("hue")
			color.H = lifx.Float32Ptr(hue)
		}

		if saturationFlag != "" {
			saturation := args.Flags.Float32("saturation")
			color.S = lifx.Float32Ptr(saturation)
		}
		state.Color = color

	} else if rgbFlag != "" {
		color, err := parseRGB(rgbFlag)
		if err != nil {
			return ExitFailure, err
		}
		state.Color = color
	} else if name != "" {
		hsb, ok := args.Config.Colors[name]
		if !ok {
			return ExitFailure, fmt.Errorf("%s is not a defined color", name)
		}
		color, err := lifx.NewHSBColor(hsb[0], hsb[1], hsb[2])
		if err != nil {
			return ExitFailure, err
		}
		state.Color = color
	}

	brightnessFlag := args.Flags.String("brightness")
	if brightnessFlag != "" {
		brightness := args.Flags.Float64("brightness")
		state.Brightness = brightness
	}

	duration := args.Flags.Float64("duration")
	state.Duration = duration

	fast := args.Flags.Bool("fast")
	state.Fast = fast

	r, err := c.SetState(selector, state)
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		return ExitFailure, err
	}

	if !fast {
		PrintResults(r.Results)
	}

	return ExitSuccess, nil
}
