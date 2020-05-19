package terminal

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/logrusorgru/aurora"
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

func Bold(message string) string {
	return fmt.Sprintf("%s", aurora.Bold(message))
}

func Error(message string) string {
	return fmt.Sprintf("%s", aurora.Red(message))
}

func Info(message string) string {
	return fmt.Sprintf("%s", aurora.Cyan(message))
}

func Success(message string) string {
	return fmt.Sprintf("%s", aurora.Green(message))
}

func Warning(message string) string {
	return fmt.Sprintf("%s", aurora.Yellow(message))
}

func White(message string) string {
	return fmt.Sprintf("%s", aurora.White(message))
}

func Gray(message string) string {
	return fmt.Sprintf("%s", aurora.Gray(10, message))
}
