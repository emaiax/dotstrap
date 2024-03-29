package packages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallDryRun(t *testing.T) {
	env, _ := config.Load("testdata/install.nofile.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	assert.NoFileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	Install(&pack, true)

	assert.NoFileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)
}

func TestInstallSourceDoesntExist(t *testing.T) {
	env, _ := config.Load("testdata/install.nofile.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	assert.NoFileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	Install(&pack, false)

	assert.NoFileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.Equal(t, pack.InstallStatus(), config.NotInstalled)
}

func TestInstallLinkSuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	Install(&pack, false)

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

	Install(&pack, false)

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	assert.Equal(t, pack.InstallStatus(), config.FullyInstalled)

	// cleaning
	//
	os.Remove(file.Target)
}

func TestInstallPartiallySuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.multiple.yml")

	pack := env.Packages["mypackage"]

	copypack := pack.Files[0]
	linkpack := pack.Files[1]

	// this file exist but it won't be installed without FORCE
	//
	os.Symlink(copypack.Source, copypack.Target)

	assert.FileExists(t, copypack.Source)
	assert.FileExists(t, linkpack.Source)
	assert.FileExists(t, copypack.Target)
	assert.NoFileExists(t, linkpack.Target)

	Install(&pack, false)

	assert.Equal(t, pack.InstallStatus(), config.PartiallyInstalled)

	assert.FileExists(t, copypack.Source)
	assert.FileExists(t, linkpack.Source)
	assert.FileExists(t, copypack.Target)
	assert.FileExists(t, linkpack.Target)

	// cleaning
	//
	os.Remove(copypack.Target)
	os.Remove(linkpack.Target)
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

	Install(&pack, false)

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

	Install(&pack, false)

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
