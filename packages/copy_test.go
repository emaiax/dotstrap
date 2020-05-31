package packages

import (
	"os"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallCopyFilesSuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.copy.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	Install(&pack)

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)

	// cleaning
	//
	os.Remove(file.Target)
}

func TestInstallCopyFilesError(t *testing.T) {
	env, _ := config.Load("testdata/install.copy.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

  os.Symlink(file.Source, file.Target)

	// both files exists
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	Install(&pack)

	assert.Equal(t, pack.InstallStatus(), config.NotInstalled)

	// cleaning
	//
	os.Remove(file.Target)
}
