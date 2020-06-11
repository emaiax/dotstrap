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
			fmt.Println(copyError(err))

			return false
		}

		defer sourceFile.Close()

		targetFile, err := os.Create(target)

		if err != nil {
			fmt.Println(copyError(err))

			return false
		}

		defer targetFile.Close()

		io.Copy(targetFile, sourceFile)

		fmt.Println(copyCreatedWarning(name))
	}

	return true
}

func copyError(err error) string {
	return tty.Sprintf(
		tty.Error("Error copying file: %s"),
		tty.Error(fmt.Sprint(err)).Bold(),
	)
}

func copyCreatedWarning(name string) string {
	return tty.Sprintf(
		tty.Warning("Created copy for %s"),
		tty.Warning(name).Bold(),
	)
}
