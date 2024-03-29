package packages

import (
	"fmt"
	"io"
	"os"

	"github.com/emaiax/dotstrap/tty"
)

func copyFile(name, source, target string) bool {
	if fileExist(target) {
		backupFileName := backupFileName(target)

		if backupFile, createdBackup := backupFile(target, backupFileName); createdBackup {
			fmt.Println(backupFile)
			return copyFile(name, source, target)
		}
	} else {
		sourceFile, err := os.Open(source)

		if err != nil {
			fmt.Println(tty.CopyErrorMessage(fmt.Sprint(err)))

			return false
		}

		defer sourceFile.Close()

		targetFile, err := os.Create(target)

		if err != nil {
			fmt.Println(tty.CopyErrorMessage(fmt.Sprint(err)))

			return false
		}

		defer targetFile.Close()

		io.Copy(targetFile, sourceFile)

		fmt.Println(tty.CopyCreatedMessage(name))
	}

	return true
}
