package lumecmd

import (
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
