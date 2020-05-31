package packages

import (
	"fmt"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
	// "github.com/logrusorgru/aurora"
)

func Install(pack *config.Package) {
	for index, _ := range pack.Files {
		file := &pack.Files[index]

		if fileExist(file.Target) && !pack.Force {
			fmt.Println(useForceMessage(file.Target))

			continue
		}

		if pack.Link {
			linkFile(file)
		} else {
			copyFile(file)
		}
		// fmt.Printf("%+v\n", file)
	}

	// fmt.Printf("%+v\n", pack)
}

func fileExist(file string) bool {
	_, err := os.Lstat(file)

	return !os.IsNotExist(err)
}

func useForceMessage(file string) string {
	return fmt.Sprintf(
		terminal.Warning("File %s already exist. If you want to override this behaviour, you should use the option %s."),
		terminal.Bold(file),
		terminal.Bold("force"),
	)
}
