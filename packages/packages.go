package packages

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func Install(pack *config.Package) {
  for index, _ := range pack.Files {
    file := &pack.Files[index]

    if pack.Link {
      linkFile(file, pack.Force)
    } else {
      copyFile(file, pack.Force)
    }
    // fmt.Printf("%+v\n", file)
  }

  // fmt.Printf("%+v\n", pack)
}

func linkFile(file *config.PackageFile, force bool) {
  err := os.Symlink(file.Source, file.Target)

  if os.IsExist(err) {
    if force {
      backupFileName := backupFileName(file.Target)

      if _, createdBackup := backupFile(file.Target, backupFileName); createdBackup {
        linkFile(file, force)

        return
      }
    } else {
      fmt.Println(terminal.Error("Link already exists. If you want to override this symlink, you should use the option"), terminal.Bold("force"))
      fmt.Println(err.Error())
      fmt.Println()

      file.Installed = false
    }
  } else {
    file.Installed = true

    fmt.Println("Created symlink for", terminal.Bold(file.Name))
  }
}

func copyFile(file *config.PackageFile, force bool) {
  fmt.Println(terminal.Error("COPYING FILES IS NOT SUPPORTED"))
}

func backupFile(oldFileName, newFileName string) (string, bool) {
  _, err := os.Lstat(newFileName)

  if err == nil {
    return backupFile(newFileName, backupFileName(newFileName))
  } else if os.IsNotExist(err) {
    err = os.Rename(oldFileName, newFileName)

    if err != nil {
      fmt.Println(terminal.Error("Error creating backup file [1]"))
      fmt.Println(err)
    } else {
      fmt.Println(terminal.Warning("File already exist, created backup to " + terminal.Bold(newFileName)))

      return newFileName, true
    }
  } else {
    fmt.Println(terminal.Error("Error creating backup file [2]"))
    fmt.Println(err)
  }

  return "", false
}

func backupFileName(file string) string {
  backupFileName := regexp.MustCompile(`\.\d{15,}$`).ReplaceAllString(file, "")

  return fmt.Sprintf("%s.%d", backupFileName, time.Now().UnixNano())
}
