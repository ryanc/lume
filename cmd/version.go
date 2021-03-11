package lumecmd

import (
	"fmt"
	"strings"
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
	var b strings.Builder

	fmt.Fprintf(&b, "lume %s", Version)
	b.WriteString(" ")
	if GitCommit != "" {
		fmt.Fprintf(&b, "(git: %s)", GitCommit)
		b.WriteString(" ")
	}
	if BuildDate != "" {
		fmt.Fprintf(&b, "build_date: %s", BuildDate)
	}

	fmt.Println(b.String())
	return ExitSuccess, nil
}
