package terminal

import (
	"strings"
	"testing"

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
