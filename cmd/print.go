package lumecmd

import (
	"fmt"
	"time"

	"git.kill0.net/chill9/lifx-go"
	"github.com/fatih/color"
)

func ColorizePower(s string) string {
	c := color.New(color.FgRed)
	if s == "on" {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(s)
}

func ColorizeStatus(s lifx.Status) string {
	c := color.New(color.FgRed)
	if s == "ok" {
		c = color.New(color.FgGreen)
	}

	return c.Sprint(s)
}

func PrintWithIndent(indent int, s string) {
	fmt.Printf("%*s%s", indent, "", s)
}

func PrintfWithIndent(indent int, format string, a ...interface{}) (n int, err error) {
	format = fmt.Sprintf("%*s%s", indent, "", format)
	return fmt.Printf(format, a...)
}

func PrintResults(res []lifx.Result) {
	var length int
	var widths map[string]int

	widths = make(map[string]int)

	for _, r := range res {
		length = len(r.Id)
		if widths["id"] < length {
			widths["id"] = length
		}

		length = len(r.Label)
		if widths["label"] < length {
			widths["label"] = length
		}

		length = len(r.Status)
		if widths["status"] < length {
			widths["status"] = length
		}
	}

	sortResults(res)

	for _, r := range res {
		fmt.Printf("%*s %*s %*s\n",
			widths["id"], r.Id,
			widths["label"], r.Label,
			widths["status"], ColorizeStatus(r.Status))
	}
}

func PrintLights(lights []lifx.Light) {
	var length int
	var widths map[string]int

	widths = make(map[string]int)

	for _, l := range lights {
		length = len(l.Id)
		if widths["id"] < length {
			widths["id"] = length
		}

		length = len(l.Location.Name)
		if widths["location"] < length {
			widths["location"] = length
		}

		length = len(l.Group.Name)
		if widths["group"] < length {
			widths["group"] = length
		}

		length = len(l.Label)
		if widths["label"] < length {
			widths["label"] = length
		}

		length = len(l.LastSeen.Local().Format(time.RFC3339))
		if widths["last_seen"] < length {
			widths["last_seen"] = length
		}

		length = len(l.Power)
		if widths["power"] < length {
			widths["power"] = length
		}
	}

	sortLights(lights)

	fmt.Printf("total %d\n", len(lights))
	for _, l := range lights {
		fmt.Printf(
			"%*s %*s %*s %*s %*s %-*s\n",
			widths["id"], l.Id,
			widths["loction"], l.Location.Name,
			widths["group"], l.Group.Name,
			widths["label"], l.Label,
			widths["last_seen"], l.LastSeen.Local().Format(time.RFC3339),
			widths["power"], ColorizePower(l.Power),
		)
	}
}
