package lumecmd

import (
	lifx "git.kill0.net/chill9/lume"
)

func SetWhiteCmd(args CmdArgs) (int, error) {
	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")

	power := args.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	kelvinFlag := args.Flags.String("kelvin")
	if kelvinFlag != "" {
		kelvin := args.Flags.Int16("kelvin")
		color, err := lifx.NewWhite(kelvin)
		if err != nil {
			return ExitFailure, err
		}
		state.Color = color
	}

	name := args.Flags.String("name")
	if name != "" {
		name := args.Flags.String("name")
		color, err := lifx.NewWhiteString(name)
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

	infraredFlag := args.Flags.String("infrared")
	if infraredFlag != "" {
		infrared := args.Flags.Float64("infrared")
		state.Infrared = infrared
	}

	fast := args.Flags.Bool("fast")
	state.Fast = fast

	if power == "" && kelvinFlag == "" && name == "" && brightnessFlag == "" && infraredFlag == "" {
		printCmdHelp(args.Name)
		return ExitFailure, nil
	}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}

	if !fast {
		PrintResults(r.Results)
	}

	return ExitSuccess, nil
}
