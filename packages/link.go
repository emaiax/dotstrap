package packages

import (
	"fmt"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func linkFile(file *config.PackageFile) {
  if fileExist(file.Target) {
    backupFileName := backupFileName(file.Target)

    if _, createdBackup := backupFile(file.Target, backupFileName); createdBackup {
      linkFile(file)

      return
    }
  } else {
    err := os.Symlink(file.Source, file.Target)

    if err != nil {
      fmt.Println(terminal.Error("Error linking file"))
      fmt.Println(err)
    } else {
      file.Installed = true

      fmt.Println("Created symlink for", terminal.Bold(file.Name))
    }
  }
}
