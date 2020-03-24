package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	fc "github.com/fatih/color"
)

import (
	"git.kill0.net/chill9/go-lifx"
)

var (
	idWidth       int = 0
	locationWidth int = 0
	groupWidth    int = 0
	labelWidth    int = 0
	lastSeenWidth int = 0
	powerWidth    int = 0
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
		//r        *lifx.Response
		err   error
		color lifx.HSBKColor
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

	setWhiteCommand := flag.NewFlagSet("set-white", flag.ExitOnError)
	setWhiteCommand.String("name", "", "Set the kelvin by name")
	setWhiteCommand.String("kelvin", "", "Set the kelvin by value")
	setWhiteCommand.String("brightness", "", "Set the brightness")
	setWhiteCommand.String("duration", "", "Set the duration")
	setWhiteCommand.Bool("fast", false, "Execute fast (no response)")

	flag.Parse()

	command = flag.Arg(0)

	c := lifx.NewClient(accessToken)

	switch command {
	case "toggle":
		_, err = c.Toggle(*selector, 1)
	case "ls":
		lights, err := c.ListLights(*selector)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		PrintLights(lights)
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

		_, err = c.SetState(*selector, state)
	case "set-white":
		setWhiteCommand.Parse(os.Args[4:])

		fs := Flags{setWhiteCommand}

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

		_, err = c.SetState(*selector, state)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func calculateWidths(lights []lifx.Light) {
	var length int

	for _, l := range lights {
		length = len(l.Id)
		if idWidth < length {
			idWidth = length
		}

		length = len(l.Location.Name)
		if locationWidth < length {
			locationWidth = length
		}

		length = len(l.Group.Name)
		if groupWidth < length {
			groupWidth = length
		}

		length = len(l.Label)
		if labelWidth < length {
			labelWidth = length
		}

		length = len(l.LastSeen.Local().Format(time.RFC3339))
		if lastSeenWidth < length {
			lastSeenWidth = length
		}

		length = len(l.Power)
		if powerWidth < length {
			powerWidth = length
		}
	}
}

func PrintLights(lights []lifx.Light) {
	calculateWidths(lights)

	fmt.Printf("total %d\n", len(lights))
	for _, l := range lights {
		fmt.Printf(
			"%*s %*s %*s %*s %*s %-*s\n",
			idWidth, l.Id,
			locationWidth, l.Location.Name,
			groupWidth, l.Group.Name,
			labelWidth, l.Label,
			lastSeenWidth, l.LastSeen.Local().Format(time.RFC3339),
			powerWidth, PowerColor(l.Power),
		)
	}
}

func PowerColor(s string) string {
	color := fc.New(fc.FgRed).SprintFunc()
	if s == "on" {
		color = fc.New(fc.FgGreen).SprintFunc()
	}

	return color(s)
}
