package lumecmd

func LsCmd(args CmdArgs) (int, error) {
	c := args.Client
	selector := args.Flags.String("selector")
	lights, err := c.ListLights(selector)
	if err != nil {
		return ExitFailure, err
	}
	PrintLights(lights)
	return ExitSuccess, nil
}
