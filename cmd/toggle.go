package lumecmd

import (
	"flag"
	"fmt"

	"git.kill0.net/chill9/go-lifx"
)

func init() {
	fs := flag.NewFlagSet("toggle", flag.ExitOnError)
	duration := fs.Float64("duration", 1.0, "Set the duration")
	fs.Float64Var(duration, "d", 1.0, "Set the duration")
	selector := fs.String("selector", "all", "Set the selector")
	fs.StringVar(selector, "s", "all", "Set the selector")

	RegisterCommand("toggle", Command{
		Func:  ToggleCmd,
		Flags: fs,
	})
}

func ToggleCmd(args CmdArgs) int {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	r, err := c.Toggle(selector, duration)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	PrintResults(r)
	return 0
}

func PrintResults(resp *lifx.Response) {
	var length, idWidth, labelWidth, statusWidth int

	for _, r := range resp.Results {
		length = len(r.Id)
		if idWidth < length {
			idWidth = length
		}

		length = len(r.Label)
		if labelWidth < length {
			labelWidth = length
		}

		length = len(r.Status)
		if statusWidth < length {
			statusWidth = length
		}
	}

	for _, r := range resp.Results {
		fmt.Printf("%*s %*s %*s\n",
			idWidth, r.Id,
			labelWidth, r.Label,
			statusWidth, statusColor(r.Status))
	}
}

func statusColor(s lifx.Status) string {
	fs := "\033[1;31m%s\033[0m"
	if s == "ok" {
		fs = "\033[1;32m%s\033[0m"
	}

	return fmt.Sprintf(fs, s)
}
