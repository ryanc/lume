package lumecmd

import (
	"flag"
	"fmt"
	"strconv"

	"git.kill0.net/chill9/lifx-go"
)

const (
	ExitSuccess = iota
	ExitFailure
)

type Context struct {
	Flags  Flags
	Args   []string
	Client *lifx.Client
	Config Config
	Name   string
}

type Flags struct {
	*flag.FlagSet
}

type Command struct {
	Name  string
	Func  func(Context) (int, error)
	Flags *flag.FlagSet
	Use   string
	Short string
	Long  string
}

var commandRegistry = make(map[string]Command)

var (
	defaultSelector     string  = "all"
	defaultDuration     float64 = 1.0
	defaultPower        string  = ""
	defaultColor        string  = ""
	defaultBrightness   string  = ""
	defaultInfrared     string  = ""
	defaultFast         bool    = false
	defaultWhiteKelvin  string  = ""
	defaultWhiteName    string  = ""
	defaultHue          string  = ""
	defaultSaturation   string  = ""
	defaultRGB          string  = ""
	defaultName         string  = ""
	defaultOutputFormat string  = ""
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

func RegisterCommand(cmd Command) error {
	if _, ok := commandRegistry[cmd.Name]; ok {
		return fmt.Errorf("%s command is already registered")
	}

	if cmd.Flags == nil {
		cmd.Flags = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	}

	mergeGlobalFlags(cmd.Flags)

	commandRegistry[cmd.Name] = cmd
	return nil
}

func GetCommand(name string) (Command, bool) {
	cmd, ok := commandRegistry[name]
	return cmd, ok
}

func mergeGlobalFlags(fs *flag.FlagSet) {
	fs.Bool("debug", false, "Enable debug mode")

	outputFormat := fs.String("output-format", defaultOutputFormat, "Set the output format")
	fs.StringVar(outputFormat, "o", defaultOutputFormat, "Set the output format")
}
