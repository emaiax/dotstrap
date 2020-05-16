package files

import (
  "path/filepath"
	"fmt"
  "log"
  "io/ioutil"
  "os"
  "runtime"
)

func DotfilesPath() (string, bool) {
  var path string

  dotfilesPath, exists := os.LookupEnv("DOTFILES_PATH")

  if exists {
    path = dotfilesPath
  } else {
    currentPath, _ := os.Getwd()

    path = currentPath
  }

  path, _ = filepath.Abs(path)

  return path, exists
}

func InstallationPath() string {
  path, _ := os.UserHomeDir()

  return fmt.Sprintf("%s/.dotfiles", path)
}

func Os() string {
  return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

func InstallDotfiles() (bool, error) {
  return false, nil
}

func DotfilesFolders(path string) []string {
	var directories []string

	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		publicFolder, err := isPublicFolder(f.Name())

		if err != nil {
			log.Fatal(err)
		}

		if publicFolder {
			directories = append(directories, f.Name())
		}
	}

	return directories
}

// Checks if a given path is:
//   - a folder
//   - not private

func isPublicFolder(path string) (bool, error) {
	file, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	isPublic := string(file.Name()[0]) != "."
	isPublicFolder := file.IsDir() && isPublic

	if err == nil {
		return isPublicFolder, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return isPublicFolder, err
}
