package lumecmd

import (
	"flag"
	"fmt"
	"strconv"

	lifx "git.kill0.net/chill9/lume"
)

const (
	ExitSuccess = iota
	ExitError
)

type Config struct {
	AccessToken string               `toml:"access_token"`
	Colors      map[string][]float32 `toml:"colors"`
}

type CmdArgs struct {
	Flags  Flags
	Client *lifx.Client
	Config Config
}

type Flags struct {
	*flag.FlagSet
}

type Command struct {
	Name  string
	Func  func(CmdArgs) (int, error)
	Flags *flag.FlagSet
	Use   string
	Short string
	Long  string
}

var commandRegistry = make(map[string]Command)

var (
	defaultSelector    string  = "all"
	defaultDuration    float64 = 1.0
	defaultPower       string  = ""
	defaultColor       string  = ""
	defaultBrightness  string  = ""
	defaultInfrared    string  = ""
	defaultFast        bool    = false
	defaultWhiteKelvin string  = ""
	defaultWhiteName   string  = ""
	defaultHue         string  = ""
	defaultSaturation  string  = ""
	defaultRGB         string  = ""
	defaultName        string  = ""
)

func (f Flags) String(name string) string {
	return f.FlagSet.Lookup(name).Value.String()
}

func (f Flags) Float32(name string) float32 {
	val, _ := strconv.ParseFloat(f.String(name), 32)
	return float32(val)
}

func (f Flags) Float64(name string) float64 {
	val, _ := strconv.ParseFloat(f.String(name), 64)
	return val
}

func (f Flags) Int16(name string) int16 {
	val, _ := strconv.ParseInt(f.String(name), 10, 16)
	return int16(val)
}

func (f Flags) Bool(name string) bool {
	val, _ := strconv.ParseBool(f.String(name))
	return val
}

func RegisterCommand(name string, cmd Command) error {
	if _, ok := commandRegistry[name]; ok {
		return fmt.Errorf("%s command is already registered")
	}
	cmd.Name = name
	commandRegistry[name] = cmd
	return nil
}

func GetCommand(name string) (Command, bool) {
	cmd, ok := commandRegistry[name]
	return cmd, ok
}
