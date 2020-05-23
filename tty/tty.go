package terminal

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	COLUMNS      = 80
	DIVIDER      = "="
	BEER         = "\xf0\x9f\x8d\xba"
	BROKEN_HEART = "\xF0\x9F\x92\x94"
	CHECK_MARK   = "\xE2\x9C\x85"
	COMPUTER     = "\xF0\x9F\x92\xBB"
	WAVE         = "\xF0\x9F\x91\x8B"
	BUG          = "\xF0\x9F\x90\x9B"
)

func Start() {
	systemInfo := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	header := fmt.Sprintf("[%s - %s]", "dotfiles installation", systemInfo)

	fmt.Println(textWithDivider(header, COMPUTER))
	fmt.Println()
}

func PrintConfigs(source, target string, dryRun bool) {
	if dryRun {
		fmt.Println("This installer is running in", Info(Bold("dry run")), "mode, nothing will really be installed.")
		fmt.Println()
	}

	fmt.Println("Source:", Bold(source))
	fmt.Println("Target:", Bold(target))
	fmt.Println()
}

func PrintRevision(packagesInstall map[string]bool) {
	fmt.Println()
	fmt.Println(Bold(textWithDivider("dotfiles revision")))
	fmt.Println(packagesInstall)

	for packName, installed := range packagesInstall {
		var message string
		fmt.Println(packName)

		if installed {
			message = fmt.Sprintf(
				"%s %s %s",
				Bold(Success("[+]")),
				Bold(White(packName)),
				Bold(Success("was successfully installed")),
			)
		} else {
			message = fmt.Sprintf(
				"%s %s %s",
				Bold(Warning("[-]")),
				Bold(White(packName)),
				Bold(Warning("was partially installed (or not installed at all), please review")),
			)
		}

		fmt.Println(Bold(message))
	}

	fmt.Println()
}

func Quit() {
	fmt.Println()
	fmt.Println(Bold(Warning("Your dotfiles won't be installed at this time.")))
	fmt.Println(Error(textWithDivider("[exiting now]", BROKEN_HEART)))
}

func Finish() {
	fmt.Println(Bold(textWithDivider("dotfiles installed, please restart.", BEER)))
	fmt.Println(Success(textWithDivider("[exiting now]", WAVE)))
}

func textWithDivider(texts ...string) string {
	fullText := strings.Join(texts, " ")
	fullTextCount := strings.Count(fullText, "")

	dividers := strings.Repeat(DIVIDER, COLUMNS-fullTextCount-1)

	return fmt.Sprintf("%s %s", Bold(fullText), dividers)
}
