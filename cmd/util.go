package lumecmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"git.kill0.net/chill9/lifx-go"
)

func powerColor(s string) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "on" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
}

func statusColor(s lifx.Status) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "ok" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
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
			widths["status"], statusColor(r.Status))
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
			widths["power"], powerColor(l.Power),
		)
	}
}

func parseRGB(s string) (lifx.RGBColor, error) {
	var c lifx.RGBColor
	rgb := strings.SplitN(s, ",", 3)
	r, err := strconv.ParseUint(rgb[0], 10, 8)
	if err != nil {
		return c, err
	}
	g, err := strconv.ParseUint(rgb[1], 10, 8)
	if err != nil {
		return c, err
	}
	b, err := strconv.ParseUint(rgb[2], 10, 8)
	if err != nil {
		return c, err
	}
	return lifx.NewRGBColor(uint8(r), uint8(g), uint8(b))
}

func sortLights(lights []lifx.Light) {
	sort.Slice(lights, func(i, j int) bool {
		if lights[i].Group.Name < lights[j].Group.Name {
			return true
		}
		if lights[i].Group.Name > lights[j].Group.Name {
			return false
		}
		return lights[i].Label < lights[j].Label

	})
}

func sortResults(res []lifx.Result) {
	sort.Slice(res, func(i, j int) bool {
		return res[i].Label < res[j].Label
	})
}

func ExitWithCode(code int, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	os.Exit(code)
}

func YesNo(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

func PrintWithIndent(indent int, s string) {
	fmt.Printf("%*s%s", indent, "", s)
}

func PrintfWithIndent(indent int, format string, a ...interface{}) (n int, err error) {
	format = fmt.Sprintf("%*s%s", indent, "", format)
	return fmt.Printf(format, a...)
}
