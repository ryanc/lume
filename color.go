package lifx

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
		H *float32 `json:"hue"`
		S *float32 `json:"saturation"`
		B *float32 `json:"brightness"`
		K *int16   `json:"kelvin"`
	}

	NamedColor string
)

var (
	Candlelight    = func() HSBKColor { c, _ := NewWhite(1500); return c }
	Sunset         = func() HSBKColor { c, _ := NewWhite(2000); return c }
	UltraWarm      = func() HSBKColor { c, _ := NewWhite(2500); return c }
	Incandescent   = func() HSBKColor { c, _ := NewWhite(2700); return c }
	Warm           = func() HSBKColor { c, _ := NewWhite(3000); return c }
	Cool           = func() HSBKColor { c, _ := NewWhite(4000); return c }
	CoolDaylight   = func() HSBKColor { c, _ := NewWhite(4500); return c }
	SoftDaylight   = func() HSBKColor { c, _ := NewWhite(5000); return c }
	Daylight       = func() HSBKColor { c, _ := NewWhite(5600); return c }
	NoonDaylight   = func() HSBKColor { c, _ := NewWhite(6000); return c }
	BrightDaylight = func() HSBKColor { c, _ := NewWhite(6500); return c }
	CloudDaylight  = func() HSBKColor { c, _ := NewWhite(7000); return c }
	BlueDaylight   = func() HSBKColor { c, _ := NewWhite(7500); return c }
	BlueOvercast   = func() HSBKColor { c, _ := NewWhite(8000); return c }
	BlueIce        = func() HSBKColor { c, _ := NewWhite(9000); return c }
)

func NewRGBColor(r, g, b uint8) (*RGBColor, error) {
	return &RGBColor{R: r, G: g, B: b}, nil
}

func NewHSBColor(h, s, b float32) (HSBKColor, error) {
	var c HSBKColor

	if h < 0 || h > 360 {
		return c, errors.New("hue must be between 0.0-360.0")
	}
	if s < 0 || s > 1 {
		return c, errors.New("saturation must be between 0.0-1.0")
	}
	if b < 0 || b > 1 {
		return c, errors.New("brightness must be between 0.0-1.0")
	}

	c = HSBKColor{
		H: Float32Ptr(h),
		S: Float32Ptr(s),
		B: Float32Ptr(b),
	}

	return c, nil
}

func NewWhite(k int16) (HSBKColor, error) {
	var c HSBKColor

	if k < 1500 || k > 8000 {
		return c, errors.New("kelvin must be between 1500-9000")
	}

	c = HSBKColor{
		H: Float32Ptr(0.0),
		S: Float32Ptr(0.0),
		K: Int16Ptr(k),
	}

	return c, nil
}

func (c RGBColor) ColorString() string {
	return fmt.Sprintf("rgb:%d,%d,%d", c.R, c.G, c.B)
}

func (c RGBColor) Hex() string {
	return fmt.Sprintf("#%x%x%x", c.R, c.G, c.B)
}

func (c HSBKColor) ColorString() string {
	var s []string
	if c.H != nil {
		s = append(s, fmt.Sprintf("hue:%g", *c.H))
	}
	if c.S != nil {
		s = append(s, fmt.Sprintf("saturation:%g", *c.S))
	}
	if c.B != nil {
		s = append(s, fmt.Sprintf("brightness:%g", *c.B))
	}
	if c.K != nil {
		s = append(s, fmt.Sprintf("kelvin:%d", *c.K))
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
	var (
		err  error
		s    *HSBKColor
		resp *http.Response
	)

	if resp, err = c.validateColor(color); err != nil {
		return nil, err
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}
