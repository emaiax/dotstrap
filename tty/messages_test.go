package tty

import (
	"strings"
	"testing"

	"github.com/logrusorgru/aurora"
	"github.com/stretchr/testify/assert"
)

func TestConfirmYes(t *testing.T) {
	ioReader := strings.NewReader("y\n")

	answer := Confirm("Yes?", ioReader)

	assert.True(t, answer)
}

func TestConfirmNo(t *testing.T) {
	ioReader := strings.NewReader("n\n")

	answer := Confirm("Yes?", ioReader)

	assert.False(t, answer)
}

func TestConfirmEmpty(t *testing.T) {
	ioReader := strings.NewReader("\n")

	answer := Confirm("Yes?", ioReader)

	assert.True(t, answer)
}

func TestHeaderEmpty(t *testing.T) {
	assert.Equal(
		t,
		"================================================================================",
		Header(),
	)
}

func TestHeaderMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf("%s %s", aurora.Bold("hi"), "============================================================================="),
		Header("hi"),
	)
}

func TestUseForceMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Yellow("File %s already exist. If you want to override this behaviour, you should use the option %s."),
			aurora.Bold("file"),
			aurora.Bold("force"),
		),
		UseForceMessage("file"),
	)
}

func TestSourceFileNotFoundMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Yellow("File %s not found, %s."),
			aurora.Bold("file"),
			aurora.Yellow("skipping").Bold(),
		),
		SourceFileNotFoundMessage("file"),
	)
}

func TestBackupCreatedMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Yellow("File already exist, created backup to %s"),
			aurora.Bold("file"),
		),
		BackupCreatedMessage("file"),
	)
}

func TestBackupErrorMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Red("Error creating backup file: %s"),
			aurora.Red("error").Bold(),
		),
		BackupErrorMessage("error"),
	)
}

func TestLinkCreatedMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Green("Created symlink for %s"),
			aurora.Bold("file"),
		),
		LinkCreatedMessage("file"),
	)
}

func TestLinkErrorMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Red("Error linking file: %s"),
			aurora.Red("error").Bold(),
		),
		LinkErrorMessage("error"),
	)
}

func TestCopyCreatedMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Green("Created copy for %s"),
			aurora.Bold("file"),
		),
		CopyCreatedMessage("file"),
	)
}

func TestCopyErrorMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Red("Error copying file: %s"),
			aurora.Red("error").Bold(),
		),
		CopyErrorMessage("error"),
	)
}

func TestDryRunFileInstalledMessage(t *testing.T) {
	assert.Equal(
		t,
		aurora.Sprintf(
			aurora.Cyan("%s %s installed").Bold(),
			aurora.Cyan("[dry run]").Bold(),
			aurora.Bold("file"),
		),
		DryRunFileInstalledMessage("file"),
	)
}
