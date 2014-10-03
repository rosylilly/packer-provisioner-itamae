package itamae

import (
	"github.com/mitchellh/packer/packer"
	"io/ioutil"
	"os"
	"testing"
)

func testConfig() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func TestProvisioner_Impl(t *testing.T) {
	var raw interface{}
	raw = &Provisioner{}
	if _, ok := raw.(packer.Provisioner); !ok {
		t.Fatalf("must be a Provisioner")
	}
}

func TestProvisionerPrepare_Recipes(t *testing.T) {
	var p Provisioner
	config := testConfig()

	err := p.Prepare(config)
	if err == nil {
		t.Fatal("should have error")
	}

	config["recipes"] = []string{""}
	err = p.Prepare(config)
	if err == nil {
		t.Fatal("should have error")
	}

	recipeFile, err := ioutil.TempFile("", "recipe.rb")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	defer os.Remove(recipeFile.Name())

	config["recipes"] = []string{recipeFile.Name()}
	err = p.Prepare(config)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
