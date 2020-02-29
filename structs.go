package lifx

import (
	"fmt"
	"strings"
)

type (
	Color interface {
		ColorString() string
	}
)

type (
	Status string

	RGBColor struct {
		R, G, B uint8
	}

	HSBKColor struct {
		H, K int16
		S, B float32
	}
)

const (
	OK       Status = "ok"
	TimedOut Status = "timed_out"
	Offline  Status = "offline"
)

func (s Status) Success() bool {
	return s == OK
}

func NewRGBColor(r, g, b uint8) (*RGBColor, error) {
	return &RGBColor{R: r, G: g, B: b}, nil
}

func NewHSBKColor() HSBKColor {
	var c HSBKColor
	c.H, c.S, c.B, c.K = -1, -1, -1, -1
	return c
}

func (c RGBColor) ColorString() string {
	return fmt.Sprintf("rgb:%d,%d,%d", c.R, c.G, c.B)
}

func (c RGBColor) Hex() string {
	return fmt.Sprintf("#%x%x%x", c.R, c.G, c.B)
}

func (c HSBKColor) ColorString() string {
	var s []string
	if c.H >= 0 {
		s = append(s, fmt.Sprintf("hue:%d", c.H))
	}
	if c.S >= 0 {
		s = append(s, fmt.Sprintf("saturation:%g", c.S))
	}
	if c.B >= 0 {
		s = append(s, fmt.Sprintf("brightness:%g", c.B))
	}
	if c.K >= 0 {
		s = append(s, fmt.Sprintf("kelvin:%d", c.K))
	}
	return strings.Join(s, " ")
}

func (c HSBKColor) MarshalText() ([]byte, error) {
	return []byte(c.ColorString()), nil
}
