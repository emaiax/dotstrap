package packages

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func Install(pack config.Package) bool {
  var installedFiles = make(map[bool]int)

	fmt.Println(packNamePrefix(pack.Name), "installing files...")

	fmt.Println()

  for _, file := range pack.Files {
    var result bool

    if pack.Link {
      result = linkFile(file, pack.Force)
    } else {
      result = copyFile(file)
    }

    installedFiles[result] += 1
  }

	return getInstallationResult(installedFiles) > 0
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

// returns between -1 and 1 that responds to:
//
// 1: all files installed
// 0: some files installed
// -1: no files installed
//
func getInstallationResult(installedFiles map[bool]int) int {
  result := 1

  if installedFiles[false] > 0 {
    result -= 1
  }

  if installedFiles[true] == 0 {
    result -= 1
  }

  return result
}
