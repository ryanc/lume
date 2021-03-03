package lumecmd

import (
	"fmt"
	"runtime"
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
	fmt.Printf("lume %s\n", Version)
	fmt.Printf("  os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("  go version: %s\n", runtime.Version())
	if BuildDate != "" {
		fmt.Printf("  build date: %s\n", BuildDate)
	}
	return ExitSuccess, nil
}
