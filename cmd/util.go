package lumecmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"git.kill0.net/chill9/lifx-go"
)

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
