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
