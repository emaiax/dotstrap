package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emaiax/dotstrap/config"
	// "github.com/emaiax/dotstrap/files"
	"github.com/emaiax/dotstrap/terminal"
)

const configFile = "dotstrap.yml"

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	terminal.Init()

	env, err := config.Load(configFile)

	if os.IsNotExist(err) {
		fmt.Println(terminal.Bold(terminal.Error("Config file not found.")))

		terminal.Quit()

		os.Exit(-1)
	}

	if err != nil {
		terminal.Quit()
		os.Exit(-1)
	}

	terminal.Config(env)

	if env != nil {
		// if terminal.Confirm("Proceed?", os.Stdin) {
		//   dotPath, _ := files.DotfilesPath()
		//
		//   for _, path := range files.DotfilesFolders(dotPath) {
		//     println(path)
		//   }
		//
		//   _, err := files.InstallDotfiles()
		//
		//   if err != nil {
		//     log.Fatal(err)
		//     terminal.Quit()
		//   } else {
		//     terminal.Finish()
		//   }
		// } else {
		//   terminal.Quit()
		// }
	}

	terminal.Finish()
}
