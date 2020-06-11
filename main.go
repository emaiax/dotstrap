package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

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
		fmt.Println(tty.Error("Installer config file not found, exiting now.").Bold())

		os.Exit(-1)
	}

	if err != nil {
		fmt.Println(tty.Error(fmt.Sprint(err)).Bold())

		os.Exit(-1)
	}

	env = environment
}

func main() {
	header := fmt.Sprintf("[%s/%s - %s]", runtime.GOOS, runtime.GOARCH, "dotfiles installation")

	fmt.Println(tty.Header(header, tty.COMPUTER))
	fmt.Println()

	if env.Config.DryRun {
		fmt.Println(tty.Sprintf("This installer is running in %s mode, nothing will really be installed.", tty.Info("dry run").Bold()))
		fmt.Println()
	}

	fmt.Println("Source:", tty.Bold(env.Config.Source))
	fmt.Println("Target:", tty.Bold(env.Config.Target))
	fmt.Println()

	if tty.Confirm("Proceed to install?", os.Stdin) {
		for _, pack := range env.Packages {
			fmt.Println(tty.Sprintf("[%s] installing files...", tty.Bold(pack.Name)))
			fmt.Println()

			packages.Install(&pack, env.Config.DryRun)
		}

		fmt.Println(tty.Header("Install revision"))
		fmt.Println()

		for _, pack := range env.Packages {
			switch pack.InstallStatus() {
			case config.NotInstalled:
				fmt.Println(tty.PackageNotInstalledMessage(pack.Name))
			case config.PartiallyInstalled:
				fmt.Println(tty.PackagePartiallyInstalledMessage(pack.Name))
			case config.FullyInstalled:
				fmt.Println(tty.PackageFullyInstalledMessage(pack.Name))
			}
		}

		fmt.Println()
		fmt.Println(tty.Bold(tty.Header("dotfiles installed, please restart.", tty.BEER)))
	} else {
		fmt.Println(tty.Bold(tty.Header("your dotfiles won't be installed at this time.", tty.BROKEN_HEART)))
	}

	fmt.Println(tty.Header("[exiting now]", tty.WAVE))
}
