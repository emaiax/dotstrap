package packages

import (
	"fmt"
	"os"

	"github.com/emaiax/dotstrap/tty"
)

func linkFile(name, source, target string) bool {
	if fileExist(target) {
		backupFileName := backupFileName(target)

		if _, createdBackup := backupFile(target, backupFileName); createdBackup {
			return linkFile(name, source, target)
		}
	} else {
		err := os.Symlink(source, target)

		if err != nil {
			fmt.Println(tty.LinkErrorMessage(fmt.Sprint(err)))

			return false
		} else {
			fmt.Println(tty.LinkCreatedMessage(name))
		}
	}

	return true
}
