package lumecmd

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const lumercFile string = ".lumerc"

type Config struct {
	AccessToken  string               `toml:"access_token"`
	OutputFormat string               `toml:"output_format"`
	Colors       map[string][]float32 `toml:"colors"`
}

// Validate configuration struct
func (c *Config) Validate() error {
	var err error
	if c.AccessToken == "" {
		err = errors.New("access_token is not set")
	}

	if err = c.validateColors(); err != nil {
		return err
	}

	return err
}

func (c *Config) validateColors() (err error) {
	if len(c.Colors) > 0 {
		for name, hsb := range c.Colors {
			if len(hsb) != 3 {
				return fmt.Errorf("color '%s' needs three values", name)
			}
			h, s, b := hsb[0], hsb[1], hsb[2]
			if h < 0 || h > 360 {
				return fmt.Errorf("color '%s' hue value must be between 0.0-360.0", name)
			}
			if s < 0 || b > 1 {
				return fmt.Errorf("color '%s' saturation value must be between 0.0-1.0", name)
			}
			if b < 0 || b > 1 {
				return fmt.Errorf("color '%s' brightness value must be between 0.0-1.0", name)
			}
		}
	}
	return err
}

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
