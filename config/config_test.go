package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadNoConfig(t *testing.T) {
	env, err := Load("testdata/nope.yml")

  assert.Nil(t, env)

  assert.EqualError(t, err, "open testdata/nope.yml: no such file or directory")
}

func TestLoadInvaliConfig(t *testing.T) {
	env, err := Load("testdata/config.invalid.yml")

  assert.Nil(t, env)

  assert.EqualError(t, err, "yaml: did not find expected alphabetic or numeric character")
}

func TestLoadEmptyConfig(t *testing.T) {
	env, _ := Load("testdata/config.empty.yml")

	defaultConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap/config", Target: "/root"}

	assert.Equal(t, env.Config, defaultConfig)
}

func TestLoadCustomConfig(t *testing.T) {
	env, _ := Load("testdata/config.custom.yml")

	customConfig := Config{Source: "/root/dotstrap", Target: "/root/dotfiles"}

	assert.Equal(t, env.Config, customConfig)
}

func TestLoadVariablesConfig(t *testing.T) {
	err := os.Setenv("DOTFILES", "..")

	if err != nil {
		log.Fatal(err)
	}

	env, _ := Load("testdata/config.variables.yml")

	customConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap", Target: "/root/mydotfiles"}

	assert.Equal(t, env.Config, customConfig)
}

func TestLoadValidGlobPackages(t *testing.T) {
	env, _ := Load("testdata/packages.validglob.yml")

	assert.Contains(
		t,
		env.Packages["mypackage"].Files,
		PackageFile{
			Name:   "packages.validglob.yml",
			Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.validglob.yml",
			Target: "/root/.packages.validglob.yml",
		},
	)
}

func TestLoadInvalidGlobPackages(t *testing.T) {
	env, _ := Load("testdata/packages.invalidglob.yml")

	assert.Empty(t, env.Packages["mypackage"].Files)
}

func TestLoadLinksPackages(t *testing.T) {
	env, _ := Load("testdata/packages.links.yml")

	assert.ElementsMatch(
		t,
		env.Packages["mypackage"].Files,
		[]PackageFile{
			PackageFile{
				Name:   "packages.glob.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.glob.yml",
				Target: "/root/.globpack",
			},
			PackageFile{
				Name:   "packages.links.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.links.yml",
				Target: "/root/.linkpack",
			},
		},
	)
}

func TestResolveEnvVarToValue(t *testing.T) {
  path := resolveEnvVar("${HOME}", os.UserHomeDir)

  assert.Equal(t, path, "/root")
}

func TestResolveEnvVarToDefault(t *testing.T) {
  path := resolveEnvVar("${NOPE}", os.Getwd)

  assert.Equal(t, path, "/go/src/github.com/emaiax/dotstrap/config")
}

func TestPublicPath(t *testing.T) {
  assert.Equal(t, getPublicPath("/go/", "/xpto/file"), "/go/xpto/file")
}


func TestPrivatePath(t *testing.T) {
  assert.Equal(t, getPrivatePath("/go/", "/xpto/.DS_Store"), "/go/xpto/.DS_Store")
}
