package lumecmd

import (
	"flag"
	"fmt"

	"git.kill0.net/chill9/lifx-go"
)

func NewCmdSetColor() Command {
	return Command{
		Name: "set-color",
		Func: SetColorCmd,
		Flags: func() *flag.FlagSet {
			fs := flag.NewFlagSet("set-color", flag.ExitOnError)

			selector := fs.String("selector", "all", "the selector")
			fs.StringVar(selector, "s", "all", "the selector")

			power := fs.String("power", defaultPower, "power state")
			fs.StringVar(power, "p", defaultPower, "power state")

			hue := fs.String("hue", defaultHue, "hue level")
			fs.StringVar(hue, "H", defaultHue, "hue level")

			saturation := fs.String("saturation", defaultSaturation, "saturation level")
			fs.StringVar(saturation, "S", defaultSaturation, "saturation level")

			rgb := fs.String("rgb", defaultRGB, "RGB value")
			fs.StringVar(rgb, "r", defaultRGB, "RGB value")

			name := fs.String("name", defaultName, "named color")
			fs.StringVar(name, "n", defaultName, "named color")

			brightness := fs.String("brightness", defaultBrightness, "brightness state")
			fs.StringVar(brightness, "b", defaultBrightness, "brightness state")

			duration := fs.Float64("duration", defaultDuration, "duration state")
			fs.Float64Var(duration, "d", defaultDuration, "duration state")

			fast := fs.Bool("fast", defaultFast, "fast state")
			fs.BoolVar(fast, "f", defaultFast, "fast state")

			return fs
		}(),
		Use:   "[--selector <selector>] [--power (on|off)] [--hue <hue>] [--saturation <saturation>] [--rgb <rbg>] [--name <color>] [--brightness <brightness>] [--duration <sec>] [--fast]",
		Short: "Set the color",
	}
}

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
