package main

import (
	"os"

	lumecmd "git.kill0.net/chill9/lume/cmd"
)

func main() {
	os.Exit(lumecmd.Main(os.Args))
}
