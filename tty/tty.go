package tty

import "github.com/logrusorgru/aurora"

func Success(message string) aurora.Value {
	return aurora.Green(message)
}

func Info(message string) aurora.Value {
	return aurora.Cyan(message)
}

func Error(message string) aurora.Value {
	return aurora.Red(message)
}

func Warning(message string) aurora.Value {
	return aurora.Yellow(message)
}

func Bold(message string) aurora.Value {
	return aurora.Bold(message)
}

func White(message string) aurora.Value {
	return aurora.White(message)
}

func Sprintf(format interface{}, args ...interface{}) string {
	return aurora.Sprintf(format, args...)
}
