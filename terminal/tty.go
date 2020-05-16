package terminal

import (
	"fmt"
  "strings"

	"github.com/emaiax/dotstrap/config"
	"github.com/emaiax/dotstrap/files"
)

const (
  COLUMNS      = 80
  DIVIDER      = "="
	BEER         = "\xf0\x9f\x8d\xba"
	BROKEN_HEART = "\xF0\x9F\x92\x94"
	CHECK_MARK   = "\xE2\x9C\x85"
  COMPUTER     = "\xF0\x9F\x92\xBB"
  WAVE         = "\xF0\x9F\x91\x8B"
)

func Init() {
  fmt.Println(textWithDivider(header(), COMPUTER))
  fmt.Println()
}

func Config(env *config.Environment) {
  fmt.Println("Source:", Bold(env.Config.Source))
  fmt.Println("Target:", Bold(env.Config.Target))
}

func Quit() {
  fmt.Println()
  fmt.Println(Bold(Warning("Your dotfiles won't be installed at this time.")))
  fmt.Println(Error(textWithDivider("[exiting now]", BROKEN_HEART)))
}

func Finish() {
  fmt.Println()
  fmt.Println(Bold(textWithDivider("dotfiles installed, please restart your terminal and vim", BEER)))
  fmt.Println(Success(textWithDivider("[exiting now]", WAVE)))
}

func textWithDivider(texts ...string) string {
  fullText := strings.Join(texts, " ")
  fullTextCount := strings.Count(fullText, "")

  dividers := strings.Repeat(DIVIDER, COLUMNS - fullTextCount - 1)

  return fmt.Sprintf("%s %s", Bold(fullText), dividers)
}

func header() string {
  return fmt.Sprintf("[%s - %s]", "dotfiles installation", files.Os())
}
