package packages

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/terminal"
)

func Install(pack config.Package) bool {
	var result bool = true

	fmt.Println(packNamePrefix(pack.Name), "Installing files...")

	fmt.Println()

	for _, file := range pack.Files {
		if pack.Link {
      if !linkFile(file, pack.Force) {
        result = false
      }
		} else {
      result = copyFile(file)
		}
	}

	return result
}

func packNamePrefix(packName string) string {
	return terminal.Bold(fmt.Sprintf("[%s]", packName))
}

func linkFile(file config.PackageFile, force bool) bool {
	var result bool = true

  err := os.Symlink(file.Source, file.Target)

  if os.IsExist(err) && force {
    if backupExistingLink(file.Target) {
      return linkFile(file, force)
    } else {
      return false
    }
  }

  if os.IsExist(err) {
    fmt.Println(terminal.Error("Link already exists. If you want to override this symlink, you should use the option"), terminal.Bold("force"))
    fmt.Println(err.Error())
    fmt.Println()

    result = false
  } else {
    _, err := filepath.EvalSymlinks(file.Target)

    if err != nil {
      fmt.Println(terminal.Error("ERROR LINKING FILES #2"))
      fmt.Println(err)

      result = false
    } else {
      fmt.Println("Created symlink for", terminal.Bold(file.Name))
    }
  }

  return result
}

func backupExistingLink(link string) bool {
  backupLink := fmt.Sprintf("%s.%d", link, time.Now().UnixNano())

  _, backupLinkErr := os.Lstat(backupLink)

  if os.IsExist(backupLinkErr) {
    return backupExistingLink(link)
  }

  err := os.Rename(link, backupLink)

  if err != nil {
    fmt.Println(terminal.Error("Error renaming backuplink"))

    return backupExistingLink(link)
  } else {
    fmt.Println(terminal.Warning("Symlink already exist, created backup to " + terminal.Bold(backupLink)))

    return true
  }
}

func copyFile(file config.PackageFile) bool {
  fmt.Println(terminal.Error("COPYING FILES IS NOT SUPPORTED AT THIS TIME."))

  return false
}
