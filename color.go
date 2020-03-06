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
	RGBColor struct {
		R, G, B uint8
	}

	HSBKColor struct {
		H float32 `json:"hue"`
		S float32 `json:"saturation"`
		B float32 `json:"brightness"`
		K int16   `json:"kelvin"`
	}

	NamedColor string
)

var (
	Candlelight    = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 1500} }
	Sunset         = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 2000} }
	UltraWarm      = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 2500} }
	Incandescent   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 2700} }
	Warm           = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 3000} }
	Cool           = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 4000} }
	CoolDaylight   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 4500} }
	SoftDaylight   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 5000} }
	Daylight       = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 5600} }
	NoonDaylight   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 6000} }
	BrightDaylight = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 6500} }
	CloudDaylight  = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 7000} }
	BlueDaylight   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 7500} }
	BlueOvercast   = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 8000} }
	BlueIce        = func() *HSBKColor { return &HSBKColor{H: 0, S: 0, K: 9000} }
)

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
		s = append(s, fmt.Sprintf("hue:%g", c.H))
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

func (c NamedColor) ColorString() string {
	return string(c)
}

func (c *Client) ValidateColor(color Color) (Color, error) {
	resp, err := c.Request("GET", EndpointColor(color.ColorString()), nil)
	if err != nil {
		return nil, err
	}

	s := &HSBKColor{}
	err = c.UnmarshalResponse(resp, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
