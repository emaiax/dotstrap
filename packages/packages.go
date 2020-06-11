package packages

import (
	"fmt"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/tty"
)

func Install(pack *config.Package, dryRun bool) {
	var installFile func(name, source, target string) bool

	for index, _ := range pack.Files {
		file := &pack.Files[index]

		if dryRun {
			fmt.Println(tty.DryRunFileInstalledMessage(file.Name))

			file.Installed = true

			continue
		}

		if !fileExist(file.Source) {
			fmt.Println(tty.SourceFileNotFoundMessage(file.Source))

			continue
		}

		if fileExist(file.Target) && !pack.Force {
			fmt.Println(tty.UseForceMessage(file.Target))

			continue
		}

		if pack.Link {
			installFile = linkFile
		} else {
			installFile = copyFile
		}

		file.Installed = installFile(file.Name, file.Source, file.Target)
	}

	fmt.Println()
}

func fileExist(file string) bool {
	_, err := os.Lstat(file)

	return err == nil
}
