package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/rosylilly/packer-provisioner-itamae/itamae"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterProvisioner(new(itamae.Provisioner))
}
