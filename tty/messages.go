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
