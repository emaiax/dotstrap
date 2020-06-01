package packages

import (
	"fmt"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func Install(pack *config.Package) {
	var installFile func(name, source, target string) bool

	for index, _ := range pack.Files {
		file := &pack.Files[index]

		if fileExist(file.Target) && !pack.Force {
			fmt.Println(useForceMessage(file.Target))

			continue
		}

		if pack.Link {
			installFile = linkFile
		} else {
			installFile = copyFile
		}

		file.Installed = installFile(file.Name, file.Source, file.Target)
	}
}

func fileExist(file string) bool {
	_, err := os.Lstat(file)

	return err == nil
}

func useForceMessage(file string) string {
	return fmt.Sprintf(
		terminal.Warning("File %s already exist. If you want to override this behaviour, you should use the option %s."),
		terminal.Bold(file),
		terminal.Bold("force"),
	)
}
