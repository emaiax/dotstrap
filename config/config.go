package config

import (
	"fmt"
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

type Package struct {
	Name string
	Path string `yaml:"path"`

	// Create a symlink instead of copying the file
	//
	Link bool `yaml:"link"`

	// If the symlink or the file already exists, creates a backup of the original symlink/file and re-do the operation
	//
	Force bool `yaml:"force"`

	// Only runs this package for this specified OS
	//
	Os string `yaml:"os"`

	// Allows creating symlinks in multiple but specificed files at once besides linking the package folder
	// Other options such as `link` and `force` will also be applied for these files.
	//
	Links map[string]string `yaml:"links"`

	// Glob allows applying the same operation in multiple files at once WITHOUT linking the package folder.
	// Other options such as `link` and `force` will only be applied for the files inside the `path` option.
	//
	Glob bool `yaml:"glob"`

	// Array that'll be used to keep track of sources and targets when linking from `glob` or `links`
	//
	Files []PackageFile
}

type PackageFile struct {
	Name   string
	Source string
	Target string
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

	// isLinuxError := exec.Command("/bin/bash", "-c", "[ `uname` = Linux ]").Run()
	//
	// if isLinuxError != nil {
	//   fmt.Println("isNotLinux")
	// } else {
	//   fmt.Println("isLinux")
	// }

	env.Config.resolvePaths()

	// if _, debug := os.LookupEnv("DEBUG"); debug {
	// 	fmt.Printf("Config: %+v\n", env.Config)
	// }

	for key, pack := range env.Packages {
		pack.Name = key

		// resolve globs from packages
		//
		if pack.Glob {
			fullPath := getPublicPath(env.Config.Source, pack.Path)

			files, _ := filepath.Glob(fullPath)

			for _, file := range files {
				baseName := filepath.Base(file)

				fileSource := file
				fileTarget := getPrivatePath(env.Config.Target, baseName)

				globFile := PackageFile{Name: baseName, Source: fileSource, Target: fileTarget}

				pack.Files = append(pack.Files, globFile)
			}
		}

		// resolve links from packages
		//
		for targetName, sourceName := range pack.Links {
			baseName := filepath.Base(sourceName)

			sourcePath := getPublicPath(env.Config.Source, sourceName)
			targetPath := getPrivatePath(env.Config.Target, targetName)

			linkFile := PackageFile{Name: baseName, Source: sourcePath, Target: targetPath}

			pack.Files = append(pack.Files, linkFile)
		}

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

func getPublicPath(base string, file string) string {
	return filepath.Join(base, file)
}

func getPrivatePath(base string, file string) string {
  if strings.HasPrefix(file, ".") {
    return filepath.Join(base, file)
  } else {
    return filepath.Join(base, "." + file)
  }
}
