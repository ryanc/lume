package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

import (
	"git.kill0.net/chill9/go-lifx"
	lumecmd "git.kill0.net/chill9/go-lifx/cmd"
)

func main() {
	var (
		command  string
		selector string
		//r        *lifx.Response
		err   error
		color lifx.HSBKColor
	)

	accessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("LIFX_ACCESS_TOKEN is undefined")
		os.Exit(1)
	}

	flag.StringVar(&selector, "selector", "all", "LIFX selector")

	setStateCommand := flag.NewFlagSet("set-state", flag.ExitOnError)
	setStateCommand.String("power", "", "Set the power state (on/off)")
	setStateCommand.String("color", "", "Set the color (HSBK)")
	setStateCommand.String("brightness", "", "Set the brightness")
	setStateCommand.String("duration", "", "Set the duration")
	setStateCommand.String("infrared", "", "Set the infrared brightness")
	setStateCommand.Bool("fast", false, "Execute fast (no response)")

	setWhiteCommand := flag.NewFlagSet("set-white", flag.ExitOnError)
	setWhiteCommand.String("name", "", "Set the kelvin by name")
	setWhiteCommand.String("kelvin", "", "Set the kelvin by value")
	setWhiteCommand.String("brightness", "", "Set the brightness")
	setWhiteCommand.String("duration", "", "Set the duration")
	setWhiteCommand.Bool("fast", false, "Execute fast (no response)")

	flag.Parse()

	command = flag.Arg(0)

	c := lifx.NewClient(accessToken)

	cmdArgs := lumecmd.CmdArgs{
		Client: c,
	}

	switch command {
	case "toggle":
		_, err = c.Toggle(selector, 1)
	case "ls":
		lumecmd.LsCmd(cmdArgs)
	case "set-state":
		setStateCommand.Parse(os.Args[4:])

		fs := lumecmd.Flags{setStateCommand}

		power := fs.String("power")
		color := fs.String("color")
		brightness := fs.String("brightness")
		duration := fs.String("duration")
		infrared := fs.String("infrared")
		fast := fs.String("fast")

		state := lifx.State{}

		if power != "" {
			state.Power = power
		}
		if color != "" {
			state.Color = lifx.NamedColor(color)
		}
		if brightness != "" {
			state.Brightness, err = strconv.ParseFloat(brightness, 64)
		}
		if duration != "" {
			state.Duration, err = strconv.ParseFloat(duration, 64)
		}
		if infrared != "" {
			state.Infrared, err = strconv.ParseFloat(infrared, 64)
		}
		if fast != "" {
			state.Fast, err = strconv.ParseBool(fast)
		}

		_, err = c.SetState(selector, state)
	case "set-white":
		setWhiteCommand.Parse(os.Args[4:])

		fs := lumecmd.Flags{setWhiteCommand}

		name := fs.String("name")
		kelvin := fs.String("kelvin")
		brightness := fs.String("brightness")
		duration := fs.String("duration")
		fast := fs.String("fast")

		state := lifx.State{}

		if name != "" {
			color, err := lifx.NewWhiteString(name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			state.Color = color
		}
		if kelvin != "" {
			k, _ := strconv.ParseInt(kelvin, 10, 16)
			color, err = lifx.NewWhite(int16(k))
			state.Color = color
		}
		if brightness != "" {
			state.Brightness, err = strconv.ParseFloat(brightness, 64)
		}
		if duration != "" {
			state.Duration, err = strconv.ParseFloat(duration, 64)
		}
		if fast != "" {
			state.Fast, err = strconv.ParseBool(fast)
		}

		_, err = c.SetState(selector, state)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
