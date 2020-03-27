package lumecmd

import (
	"flag"
	"fmt"
	"strconv"

	"git.kill0.net/chill9/go-lifx"
)

type CmdArgs struct {
	Flags    Flags
	Client   *lifx.Client
	Selector string
}

type Flags struct {
	*flag.FlagSet
}

type Command struct {
	Func  func(CmdArgs) int
	Flags *flag.FlagSet
}

var commandRegistry = make(map[string]Command)

func (f Flags) String(name string) string {
	return f.FlagSet.Lookup(name).Value.String()
}

func (f Flags) Float64(name string) float64 {
	val, _ := strconv.ParseFloat(f.String(name), 64)
	return val
}

func (f Flags) Bool(name string) bool {
	val, _ := strconv.ParseBool(f.String(name))
	return val
}

func RegisterCommand(name string, cmd Command) error {
	if _, ok := commandRegistry[name]; ok {
		return fmt.Errorf("%s command is already registered")
	}
	commandRegistry[name] = cmd
	return nil
}

func GetCommand(name string) (Command, bool) {
	cmd, ok := commandRegistry[name]
	return cmd, ok
}
