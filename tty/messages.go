package tty

import (
	"bufio"
	"fmt"
	"io"
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

func Confirm(message string, ioReader io.Reader) bool {
	reader := bufio.NewReader(ioReader)

	fmt.Println(confirmationMessage(message))

	text, _ := reader.ReadString('\n')

	if text != "\n" {
		fmt.Println()
	}

	return strings.ToLower(text) != strings.ToLower("n\n")
}

func confirmationMessage(text string) string {
	return fmt.Sprintf("%s %ses %so", text, Bold("[Y]"), Bold("[n]"))
}

func Header(texts ...string) string {
	fullText := strings.TrimSpace(strings.Join(texts, " "))
	fullTextCount := strings.Count(fullText, "")

	if fullText != "" {
		dividers := strings.Repeat(DIVIDER, COLUMNS-fullTextCount)

		return fmt.Sprintf("%s %s", Bold(fullText), dividers)
	} else {
		dividers := strings.Repeat(DIVIDER, COLUMNS)

		return dividers
	}
}

func UseForceMessage(file string) string {
	return Sprintf(
		Warning("File %s already exist. If you want to override this behaviour, you should use the option %s."),
		Bold(file),
		Bold("force"),
	)
}

func SourceFileNotFoundMessage(file string) string {
	return Sprintf(
		Warning("File %s not found, %s."),
		Bold(file),
		Warning("skipping").Bold(),
	)
}

func BackupErrorMessage(err string) string {
	return Sprintf(Error("Error creating backup file: %s"), Error(err).Bold())
}

func BackupCreatedMessage(name string) string {
	return Sprintf(Warning("File already exist, created backup to %s"), Bold(name))
}

func LinkErrorMessage(err string) string {
	return Sprintf(Error("Error linking file: %s"), Error(err).Bold())
}

func LinkCreatedMessage(name string) string {
	return Sprintf(Success("Created symlink for %s"), Bold(name))
}

func CopyErrorMessage(err string) string {
	return Sprintf(Error("Error copying file: %s"), Error(err).Bold())
}

func CopyCreatedMessage(name string) string {
	return Sprintf(Success("Created copy for %s"), Bold(name))
}
