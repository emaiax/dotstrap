package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/packages"
	"github.com/emaiax/dotstrap/tty"
)

const configFile = "dotstrap.yml"

var env *config.Environment

func init() {
  log.SetOutput(os.Stdout)

  environment, err := config.Load(configFile)

  // TODO: check if user has read access to env.Config.Source
  // TODO: check if user has write access to env.Config.Target
  //
  if os.IsNotExist(err) {
    fmt.Println(terminal.Bold(terminal.Error("Installer config file not found, exiting now.")))

    os.Exit(-1)
  }

  if err != nil {
    fmt.Println(terminal.Bold(terminal.Error(fmt.Sprint(err))))

    os.Exit(-1)
  }

  env = environment
}

func main() {
	terminal.Start()

	terminal.PrintConfigs(env.Config.Source, env.Config.Target, env.Config.DryRun)

	if terminal.Confirm("Proceed to install?", os.Stdin) {
		packagesInstall := make(map[string]bool)

		for _, pack := range env.Packages {
			packagesInstall[pack.Name] = packages.Install(pack)
		}

		terminal.PrintRevision(packagesInstall)
		terminal.Finish()
	} else {
		terminal.Quit()
	}
}
