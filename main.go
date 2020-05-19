package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/packages"
	"github.com/emaiax/dotstrap/terminal"
)

const configFile = "dotstrap.yml"

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	terminal.Start()

	env, err := config.Load(configFile)

	if os.IsNotExist(err) {
		fmt.Println(terminal.Bold(terminal.Error("Config file not found.")))

		terminal.Quit()

		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(terminal.Bold(terminal.Error(fmt.Sprint(err))))

		terminal.Quit()

		os.Exit(-1)
	}

	terminal.PrintConfigs(env.Config.Source, env.Config.Target, env.Config.DryRun)

	if terminal.Confirm("Proceed to install?", os.Stdin) {
		packagesInstall := make(map[string]bool)

		for _, pack := range env.Packages {
			packagesInstall[pack.Name] = packages.Install(pack, env.Config)
		}

		terminal.PrintRevision(packagesInstall)
		terminal.Finish()
	} else {
		terminal.Quit()
	}
}
