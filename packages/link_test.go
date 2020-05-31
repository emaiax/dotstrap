package packages

import (
	"os"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallLinkFilesSuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

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

func TestInstallLinkFilesError(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	os.Symlink(file.Source, file.Target)

	// both files exists
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// error when trying to link again without FORCE option
	//
	Install(&pack)

	assert.Equal(t, pack.InstallStatus(), config.NotInstalled)

	// cleaning
	//
	os.Remove(file.Target)
}

func TestInstallLinkFilesForce(t *testing.T) {
	env, _ := config.Load("testdata/install.forcelinks.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	Install(&pack)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// success when trying to link again with FORCE option
	//
	Install(&pack)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)

	// both files still exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// cleaning
	//
	os.Remove(file.Target)
}
