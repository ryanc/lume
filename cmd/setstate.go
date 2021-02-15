package lumecmd

import (
	"git.kill0.net/chill9/lifx-go"
)

func SetStateCmd(args CmdArgs) (int, error) {
	c := args.Client
	state := lifx.State{}
	selector := args.Flags.String("selector")

	power := args.Flags.String("power")
	if power != "" {
		state.Power = power
	}

	color := args.Flags.String("color")
	if color != "" {
		state.Color = lifx.NamedColor(color)
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

	if power == "" && color == "" && brightnessFlag == "" && infraredFlag == "" {
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
