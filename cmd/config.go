package lumecmd

import (
	"os"
	"path"
)

const lumercFile string = ".lumerc"

func getConfigPath() string {
	var tryPath, configPath string

	// ~/.lumerc
	homeDir, err := os.UserHomeDir()
	if err == nil {
		tryPath = path.Join(homeDir, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	// ./.lumerc
	cwd, err := os.Getwd()
	if err == nil {
		tryPath = path.Join(cwd, lumercFile)
		if _, err := os.Stat(tryPath); !os.IsNotExist(err) {
			configPath = tryPath
		}
	}

	return configPath
}
