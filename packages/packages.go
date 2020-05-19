package packages

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/terminal"
)

func Install(pack config.Package, config config.Config) bool {
	var result bool = true

	fmt.Println(packNamePrefix(pack.Name), "Installing files...")

	fmt.Println()

	for _, file := range pack.Files {
		if pack.Link {
			err := os.Symlink(file.Source, file.Target)

			if os.IsExist(err) && pack.Force {
				renameFileName := fmt.Sprintf("%s.%d", file.Target, time.Now().Unix())

				fmt.Println(terminal.Warning("Symlink already exists, renamed to " + renameFileName))

				os.Rename(file.Target, renameFileName)
				err = os.Symlink(file.Source, file.Target)
			}

			if err != nil {
				fmt.Println(terminal.Error("ERROR LINKING FILES #1"))
				fmt.Println(err)

				result = false
			} else {
				_, err := filepath.EvalSymlinks(file.Target)

				if err != nil {
					fmt.Println(terminal.Error("ERROR LINKING FILES #2"))
					fmt.Println(err)

					result = false
				} else {
					fmt.Println("Created symlink for", terminal.Bold(file.Name))
				}
			}
		} else {
			fmt.Println(terminal.Error("COPYING FILES IS NOT SUPPORTED AT THIS TIME."))

			result = false
		}
	}

	fmt.Println()

	return result
}

func packNamePrefix(packName string) string {
	return terminal.Bold(fmt.Sprintf("[%s]", packName))
}
