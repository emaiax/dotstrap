package config

import (
  "os"
	"path/filepath"
	"regexp"
	"strings"
)

type InstallStatus int

const (
  NotInstalled InstallStatus = iota // 0
  PartiallyInstalled                // 1
  FullyInstalled                    // 2
)

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
	Name      string
	Source    string
	Target    string
  Installed bool
}

func (pack *Package) resolveGlobFilePaths(sourcePath, targetPath string) {
  fullPath := getPublicPath(sourcePath, pack.Path)
  files, _ := filepath.Glob(fullPath)

  for _, file := range files {
    baseName := filepath.Base(file)

    fileSource := file
    fileTarget := getPrivatePath(targetPath, baseName)

    globFile := PackageFile{Name: baseName, Source: fileSource, Target: fileTarget}

    pack.Files = append(pack.Files, globFile)
  }
}

func (pack *Package) resolveLinkFilePaths(sourcePath, targetPath string) {
  for targetName, sourceName := range pack.Links {
    baseName := filepath.Base(sourceName)

    sourcePath := getPublicPath(sourcePath, sourceName)
    targetPath := getPrivatePath(targetPath, targetName)

    linkFile := PackageFile{Name: baseName, Source: sourcePath, Target: targetPath}

    pack.Files = append(pack.Files, linkFile)
  }
}

func (pack Package) InstallStatus() InstallStatus {
  var installedFiles = make(map[bool]int)

  for _, file := range pack.Files {
    installedFiles[file.Installed] += 1
  }

  if installedFiles[false] == 0 {
    return FullyInstalled
  }

  if installedFiles[true] > 0 {
    return PartiallyInstalled
  }

  return NotInstalled
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
