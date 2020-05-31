package packages

import (
	"fmt"
  "io"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func copyFile(file *config.PackageFile) {
  if fileExist(file.Target) {
    backupFileName := backupFileName(file.Target)

    if _, createdBackup := backupFile(file.Target, backupFileName); createdBackup {
      copyFile(file)

      return
    }
  } else {
    sourceFile, err := os.Open(file.Source)

    if err != nil {
      fmt.Println(terminal.Error("Error copying file from source"))
      fmt.Println(err)

      return
    }

    defer sourceFile.Close()

    targetFile, err := os.Create(file.Target)

    if err != nil {
      fmt.Println(terminal.Error("Error copying file to target"))
      fmt.Println(err)

      return
    }

    defer targetFile.Close()

    copiedBytes, err := io.Copy(targetFile, sourceFile)

    if err != nil {
      fmt.Println(terminal.Error("Error copying file from source to target"))
      fmt.Println(err)
    } else {
      fmt.Println(fmt.Printf("Copied %d bytes", copiedBytes))

      file.Installed = true
    }
  }
}
