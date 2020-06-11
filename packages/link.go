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
			fmt.Println(linkError(err))

			return false
		} else {
			fmt.Println(linkCreatedWarning(name))
		}
	}

	return true
}

func linkError(err error) string {
	return tty.Sprintf(
		tty.Error("Error linking file: %s"),
		tty.Error(fmt.Sprint(err)).Bold(),
	)
}

func linkCreatedWarning(name string) string {
	return tty.Sprintf(
		tty.Warning("Created symlink for %s"),
		tty.Warning(name).Bold(),
	)
}
