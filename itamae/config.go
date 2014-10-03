package itamae

import (
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
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
