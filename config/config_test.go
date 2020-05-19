package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEmptyConfig(t *testing.T) {
	env, _ := Load("../testdata/config.empty.yml")

	defaultConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap/config", Target: "/root"}

	assert.Equal(t, env.Config, defaultConfig)
}

func TestLoadCustomConfig(t *testing.T) {
	env, _ := Load("../testdata/config.custom.yml")

	customConfig := Config{Source: "/root/dotstrap", Target: "/root/dotfiles"}

	assert.Equal(t, env.Config, customConfig)
}

func TestLoadVariablesConfig(t *testing.T) {
	err := os.Setenv("DOTFILES", "..")

	if err != nil {
		log.Fatal(err)
	}

	env, _ := Load("../testdata/config.variables.yml")

	customConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap", Target: "/root/mydotfiles"}

	assert.Equal(t, env.Config, customConfig)
}

func TestLoadGlobPackages(t *testing.T) {
	env, _ := Load("../testdata/packages.glob.yml")

	assert.Contains(
		t,
		env.Packages["mypackage"].Files,
		PackageFile{
			Name:   "packages.glob.yml",
			Source: "/go/src/github.com/emaiax/dotstrap/testdata/packages.glob.yml",
			Target: "/root/.packages.glob.yml",
		},
	)
}

func TestLoadLinksPackages(t *testing.T) {
	env, _ := Load("../testdata/packages.links.yml")

	assert.ElementsMatch(
		t,
		env.Packages["mypackage"].Files,
		[]PackageFile{
			PackageFile{
				Name:   "packages.glob.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/testdata/packages.glob.yml",
				Target: "/root/.globpack",
			},
			PackageFile{
				Name:   "packages.links.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/testdata/packages.links.yml",
				Target: "/root/.linkpack",
			},
		},
	)
}
