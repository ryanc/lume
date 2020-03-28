package lumecmd

import (
	"flag"
	"fmt"

	"git.kill0.net/chill9/lume"
)

func init() {
	var cmdName string = "set-state"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	selector := fs.String("selector", "all", "Set the selector")
	fs.StringVar(selector, "s", "all", "Set the selector")

	power := fs.String("power", "", "power state")
	fs.StringVar(power, "p", "", "power state")

	color := fs.String("color", "", "color state")
	fs.StringVar(color, "c", "", "color state")

	brightness := fs.String("brightness", "", "brightness state")
	fs.StringVar(brightness, "b", "", "brightness state")

	duration := fs.Float64("duration", 1.0, "duration state")
	fs.Float64Var(duration, "d", 1.0, "duration state")

	infrared := fs.String("infrared", "", "infrared state")
	fs.StringVar(infrared, "i", "", "infrared state")

	fast := fs.Bool("fast", false, "fast state")
	fs.BoolVar(fast, "f", false, "fast state")

	RegisterCommand(cmdName, Command{
		Func:  SetStateCmd,
		Flags: fs,
	})
}

func SetStateCmd(args CmdArgs) int {
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

	r, err := c.SetState(selector, state)
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		return 1
	}

	if !fast {
		PrintResults(r.Results)
	}

	return 0
}
