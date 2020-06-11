package tty

import (
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	assert.Equal(t, Success("green"), aurora.Green("green"))
}

func TestInfo(t *testing.T) {
	assert.Equal(t, Info("info"), aurora.Cyan("info"))
}

func TestWarning(t *testing.T) {
	assert.Equal(t, Warning("warning"), aurora.Yellow("warning"))
}

func TestError(t *testing.T) {
	assert.Equal(t, Error("error"), aurora.Red("error"))
}

func TestWhite(t *testing.T) {
	assert.Equal(t, White("default"), aurora.White("default"))
}

func TestBold(t *testing.T) {
	assert.Equal(t, Bold("bold"), aurora.Bold("bold"))

	assert.Equal(t, Success("bold").Bold(), aurora.Green("bold").Bold())
	assert.Equal(t, Info("bold").Bold(), aurora.Cyan("bold").Bold())
	assert.Equal(t, Warning("bold").Bold(), aurora.Yellow("bold").Bold())
	assert.Equal(t, Error("bold").Bold(), aurora.Red("bold").Bold())
}

func TestSprintf(t *testing.T) {
	assert.Equal(
		t,
		Sprintf(Error("test %s"), Bold("passed")),
		aurora.Sprintf(aurora.Red("test %s"), aurora.Bold("passed")),
	)

	assert.Equal(
		t,
		Sprintf(
			Warning("File %s already exist. If you want to override this behaviour, you should use the option %s."),
			Bold("file"),
			Bold("force"),
		),
		aurora.Sprintf(
			aurora.Yellow("File %s already exist. If you want to override this behaviour, you should use the option %s."),
			aurora.Bold("file"),
			aurora.Bold("force"),
		),
	)
}
