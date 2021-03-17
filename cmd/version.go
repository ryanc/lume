package lumecmd

import (
	"fmt"
	"runtime"
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

	fmt.Fprintf(&b, "lume, version %s\n", Version)

	if GitCommit != "" {
		fmt.Fprintf(&b, "  revision:   %s\n", GitCommit)
	}

	if BuildDate != "" {
		fmt.Fprintf(&b, "  build date: %s\n", BuildDate)
	}

	fmt.Fprintf(&b, "  go version: %s\n", runtime.Version())
	fmt.Fprintf(&b, "  platform:   %s\n", runtime.GOOS+"/"+runtime.GOARCH)

	fmt.Print(b.String())

	return ExitSuccess, nil
}
