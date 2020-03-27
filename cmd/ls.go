package lumecmd

import (
	"flag"
	"fmt"
	"time"

	"git.kill0.net/chill9/go-lifx"
)

var (
	idWidth, locationWidth, groupWidth, labelWidth, lastSeenWidth, powerWidth int
)

func init() {
	fs := flag.NewFlagSet("toggle", flag.ExitOnError)
	fs.String("selector", "all", "Set the selector")

	RegisterCommand("ls", Command{
		Func:  LsCmd,
		Flags: fs,
	})
}

func LsCmd(args CmdArgs) int {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	PrintLights(lights)
	return 0
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
			powerWidth, powerColor(l.Power),
		)
	}
}

func powerColor(s string) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "on" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
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
