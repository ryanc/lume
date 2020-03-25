package lumecmd

import (
	"flag"
	"strconv"

	"git.kill0.net/chill9/go-lifx"
)

type CmdArgs struct {
	Flags    *Flags
	Client   *lifx.Client
	Selector string
}

type Flags struct {
	*flag.FlagSet
}

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
