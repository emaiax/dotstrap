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
			fmt.Println(backupError(err))
		} else {
			fmt.Println(backupCreatedWarning(newFileName))

			return newFileName, true
		}
	} else {
		fmt.Println(backupError(err))
	}

	return "", false
}

func backupError(err error) string {
	return tty.Sprintf(
		tty.Error("Error creating backup file: %s"),
		tty.Error(fmt.Sprint(err)).Bold(),
	)
}

func backupCreatedWarning(name string) string {
	return tty.Sprintf(
		tty.Warning("File already exist, created backup to %s"),
		tty.Warning(name).Bold(),
	)
}
