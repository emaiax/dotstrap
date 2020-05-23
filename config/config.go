package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"

	"gopkg.in/yaml.v2"
)

type Environment struct {
	Config   Config             `yaml:"config"`
	Packages map[string]Package `yaml:"packages"`
}

type Config struct {
  // Source path of the packages, defaults to the current directory (pwd)
	Source string `yaml:"source"`

  // Target path of the packages, defaults to User's home directory (home)
	Target string `yaml:"target"`

	// Prompts user before installing each package
	//
	Confirm bool `yaml:"confirm"`

	// Don't really install the files if it's just a dry run
	//
	DryRun bool `yaml:"dry_run"`
}

func Load(file string) (*Environment, error) {
	config, err := ioutil.ReadFile(file)

	if os.IsNotExist(err) {
		return nil, err
	}

	env := Environment{}

	err = yaml.Unmarshal(config, &env)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	env.Config.resolvePaths()

	for key, pack := range env.Packages {
		pack.Name = key

		if pack.Glob {
      pack.resolveGlobFilePaths(env.Config.Source, env.Config.Target)
		}

    pack.resolveLinkFilePaths(env.Config.Source, env.Config.Target)

		env.Packages[key] = pack
	}

	return &env, nil
}

func (config *Config) resolvePaths() {
	sourcePath, _ := filepath.Abs(resolvePath(config.Source, os.Getwd))
	targetPath, _ := filepath.Abs(resolvePath(config.Target, os.UserHomeDir))

	config.Source = sourcePath
	config.Target = targetPath
}
