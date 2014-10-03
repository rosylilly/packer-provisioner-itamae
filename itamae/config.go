package itamae

import (
	"fmt"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
	"os"
)

const (
	DefaultCommand = "itamae"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	tpl                 *packer.ConfigTemplate

	// The command to run itamae
	Command string

	// An array of local paths of recipe files to executed by itamae
	Recipes []string `mapstructure:"recipes"`
}

func (c Config) validateRecipes() []error {
	if len(c.Recipes) == 0 {
		return []error{fmt.Errorf("recipes must be specified.")}
	}

	errors := []error{}
	for _, filename := range c.Recipes {
		info, err := os.Stat(filename)
		if err != nil {
			errors = append(errors, fmt.Errorf("recipes: %s is invalid: %s", filename, err))
		} else if info.IsDir() {
			errors = append(errors, fmt.Errorf("recipes: %s must be point to a file", filename))
		}
	}

	return errors
}
