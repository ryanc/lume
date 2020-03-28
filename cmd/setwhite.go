package lumecmd

import (
	"flag"
	"fmt"

	"git.kill0.net/chill9/go-lifx"
)

func init() {
	var cmdName string = "set-white"

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	selector := fs.String("selector", "all", "the selector")
	fs.StringVar(selector, "s", "all", "the selector")

	power := fs.String("power", "", "power state")
	fs.StringVar(power, "p", "", "power state")

	kelvin := fs.String("kelvin", "", "kelvin level")
	fs.StringVar(kelvin, "k", "", "kelvin level")

	name := fs.String("name", "", "named white level")
	fs.StringVar(name, "n", "", "named white level")

	brightness := fs.String("brightness", "", "brightness state")
	fs.StringVar(brightness, "b", "", "brightness state")

	duration := fs.Float64("duration", 1.0, "duration state")
	fs.Float64Var(duration, "d", 1.0, "duration state")

	infrared := fs.String("infrared", "", "infrared state")
	fs.StringVar(infrared, "i", "", "infrared state")

	fast := fs.Bool("fast", false, "fast state")
	fs.BoolVar(fast, "f", false, "fast state")

	RegisterCommand(cmdName, Command{
		Func:  SetWhiteCmd,
		Flags: fs,
	})
}

func SetWhiteCmd(args CmdArgs) int {
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
			fmt.Printf("fatal: %s\n", err)
			return 1
		}
		state.Color = color
	}

	name := args.Flags.String("name")
	if name != "" {
		name := args.Flags.String("name")
		color, err := lifx.NewWhiteString(name)
		if err != nil {
			fmt.Printf("fatal: %s\n", err)
			return 1
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
