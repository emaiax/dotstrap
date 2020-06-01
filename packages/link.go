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
			fmt.Println(fmt.Sprintf(terminal.Error("Error linking file %s"), terminal.Bold(name)))
			fmt.Println(err)

			return false
		} else {
			fmt.Println(fmt.Sprintf(terminal.Warning("Created symlink for %s"), terminal.Bold(name)))
		}
	}

	return true
}
