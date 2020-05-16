package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type Environment struct {
	Config   Config             `yaml:"config"`
	Packages map[string]Package `yaml:"packages"`
}

type Config struct {
	Source string `yaml:"source"`
	Target string `yaml:"target"`
}

type Package struct {
	Os      string `yaml:"os"`
	Install bool   `yaml:"install"`
	Link    bool   `yaml:"link"`
	Force   bool   `yaml:"force"`
}

func Load(file string) (*Environment, error) {
	config, err := ioutil.ReadFile(file)

	if os.IsNotExist(err) {
		return nil, err
	}

	env := Environment{}

	err = yaml.Unmarshal(config, &env)

	if err != nil {
		return nil, err
	}

	// isLinuxError := exec.Command("/bin/bash", "-c", "[ `uname` = Linux ]").Run()
	//
	// if isLinuxError != nil {
	//   fmt.Println("isNotLinux")
	// } else {
	//   fmt.Println("isLinux")
	// }

	env.Config.resolvePaths()

  env.

	return &env, nil
}

func (config *Config) resolvePaths() {
	sourcePath, _ := filepath.Abs(resolvePath(config.Source, os.Getwd))
	targetPath, _ := filepath.Abs(resolvePath(config.Target, os.UserHomeDir))

	config.Source = sourcePath
	config.Target = targetPath
}

func resolvePath(s string, defaultPathFn func() (string, error)) string {
	if len(strings.TrimSpace(s)) == 0 {
		res, _ := defaultPathFn()

		return res
	} else {
		return regexp.MustCompile(`\$\{(.*?)\}`).ReplaceAllStringFunc(s, func(envvar string) string {
			return resolveEnvVar(envvar, defaultPathFn)
		})
	}
}

func resolveEnvVar(s string, defaultPathFn func() (string, error)) string {
	envName := string(s)[2 : len(s)-1]

	if envVar := os.Getenv(envName); len(os.Getenv(envName)) == 0 {
		res, _ := defaultPathFn()

		return res
	} else {
		return envVar
	}
}
