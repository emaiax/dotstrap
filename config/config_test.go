package config

import (
  "log"
  "os"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestLoadEmptyConfig(t *testing.T) {
  env, _ := Load("./testdata/empty.config.yml")

  defaultConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap/config", Target: "/root"}

  assert.Equal(t, env.Config, defaultConfig)
}

func TestLoadCustomConfig(t *testing.T) {
  env, _ := Load("./testdata/custom.config.yml")

  customConfig := Config{Source: "/root/dotstrap", Target: "/root/dotfiles"}

  assert.Equal(t, env.Config, customConfig)
}


func TestLoadVariablesConfig(t *testing.T) {
  err := os.Setenv("DOTFILES", "..")

  if err != nil {
    log.Fatal(err)
  }

  env, _ := Load("./testdata/variables.config.yml")

  customConfig := Config{Source: "/go/src/github.com/emaiax/dotstrap", Target: "/root/mydotfiles"}

  assert.Equal(t, env.Config, customConfig)
}
