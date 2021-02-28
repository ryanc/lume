package lumecmd

import (
	"fmt"
)

func NewCmdVersion() Command {
	return Command{
		Name:  "version",
		Func:  VersionCmd,
		Flags: nil,
		Use:   "",
		Short: "Show version",
	}
}

func VersionCmd(args CmdArgs) (int, error) {
	fmt.Println(Version)
	return ExitSuccess, nil
}
