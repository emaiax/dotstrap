package packages

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/emaiax/dotstrap/tty"
)

func backupFileName(file string) string {
	backupFileName := regexp.MustCompile(`\.\d{15,}$`).ReplaceAllString(file, "")

	return fmt.Sprintf("%s.%d", backupFileName, time.Now().UnixNano())
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
