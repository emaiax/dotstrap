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
	assert.Equal(t, "================================================================================", Header())
}

func TestHeaderMessage(t *testing.T) {
	assert.Equal(t, aurora.Sprintf("%s %s", aurora.Bold("hi"), "============================================================================="), Header("hi"))
}
