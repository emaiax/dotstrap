package files

import (
  "os"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestDefaultDotfilesPath(t *testing.T) {
  path, _ := DotfilesPath()

  assert.Equal(t, path, "/go/src/github.com/emaiax/dotstrap/files")
}

func TestCustomDotfilesPath(t *testing.T) {
  os.Setenv("DOTFILES_PATH", "/customfiles")
  path, _ := DotfilesPath()

  assert.Equal(t, path, "/customfiles")
}
