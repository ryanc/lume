package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

import (
	"git.kill0.net/chill9/go-lifx"
)

type Flags struct {
	*flag.FlagSet
}

func (f Flags) String(name string) string {
	return f.FlagSet.Lookup(name).Value.String()
}

func (f Flags) Float64(name string) float64 {
	val, _ := strconv.ParseFloat(f.String(name), 64)
	return val
}

func (f Flags) Bool(name string) bool {
	val, _ := strconv.ParseBool(f.String(name))
	return val
}

func main() {
	var (
		command  string
		selector *string
		r        *lifx.Response
		err      error
	)

	accessToken := os.Getenv("LIFX_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println("LIFX_ACCESS_TOKEN is undefined")
		os.Exit(1)
	}

	selector = flag.String("selector", "all", "LIFX selector")

	setStateCommand := flag.NewFlagSet("set-state", flag.ExitOnError)
	setStateCommand.String("power", "", "Set the power state (on/off)")
	setStateCommand.String("color", "", "Set the color (HSBK)")
	setStateCommand.String("brightness", "", "Set the brightness")
	setStateCommand.String("duration", "", "Set the duration")
	setStateCommand.String("infrared", "", "Set the infrared brightness")
	setStateCommand.Bool("fast", false, "Execute fast (no response)")

	flag.Parse()

	command = flag.Arg(0)

	fmt.Println(command)
	fmt.Println(*selector)

	c := lifx.NewClient(accessToken)

	switch command {
	case "toggle":
		r, err = c.Toggle(*selector, 1)
	case "set-state":
		setStateCommand.Parse(os.Args[4:])

		fs := Flags{setStateCommand}

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

		r, err = c.SetState(*selector, state)
	}

	fmt.Println(r)
	fmt.Println(err)
}
