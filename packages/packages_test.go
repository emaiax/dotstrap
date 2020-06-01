package packages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallLinkSuccess(t *testing.T) {
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

func TestInstallCopySuccess(t *testing.T) {
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

func TestInstallSkipWithoutForceWhenTargetExist(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

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

func TestInstallSuccessWithForceWhenTargetExist(t *testing.T) {
	env, _ := config.Load("testdata/install.forcelinks.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	os.Symlink(file.Source, file.Target)

	// both files exists
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	Install(&pack)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)

	// cleaning
	//
	os.Remove(file.Target)
	cleanSymlinks()
}

func cleanSymlinks() {
	files, _ := filepath.Glob("*.linkpack.*")

	for _, file := range files {
		os.Remove(file)
	}
}
