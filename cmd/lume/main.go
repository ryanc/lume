package main

import (
	"fmt"
	"os"

	lumecmd "git.kill0.net/chill9/lume/cmd"
)

func main() {
	exitCode, err := lumecmd.Main(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	os.Exit(exitCode)
}
