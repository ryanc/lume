package lumecmd

import (
	"git.kill0.net/chill9/lifx-go"
)

func PoweroffCmd(args CmdArgs) (int, error) {
	c := args.Client
	duration := args.Flags.Float64("duration")
	selector := args.Flags.String("selector")
	state := lifx.State{Power: "off", Duration: duration}

	r, err := c.SetState(selector, state)
	if err != nil {
		return ExitFailure, err
	}
	PrintResults(r.Results)
	return ExitSuccess, nil
}
