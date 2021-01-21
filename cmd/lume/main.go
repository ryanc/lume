package main

import (
	"os"

	lumecmd "git.kill0.net/chill9/lume/cmd"
)

func main() {
	lumecmd.ExitWithCode(lumecmd.Main(os.Args))
}