package packages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallLinkFilesSuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

	file := env.Packages["mypackage"].Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.True(t, Install(env.Packages["mypackage"]))

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// cleaning
	//
	cleanSymlinks(env.Config.Target)
}

func TestInstallLinkFilesError(t *testing.T) {
	env, _ := config.Load("testdata/install.links.yml")

	file := env.Packages["mypackage"].Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.True(t, Install(env.Packages["mypackage"]))

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// error when trying to link again without FORCE option
	//
	assert.False(t, Install(env.Packages["mypackage"]))

	// cleaning
	//
	cleanSymlinks(env.Config.Target)
}

func TestInstallForceLinkFiles(t *testing.T) {
	env, _ := config.Load("testdata/install.forcelinks.yml")

	file := env.Packages["mypackage"].Files[0]

	// target file doesn't exist yet
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.True(t, Install(env.Packages["mypackage"]))

	// now both files exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// success when trying to link again with FORCE option
	//
	assert.True(t, Install(env.Packages["mypackage"]))

	// now both files still exist
	//
	assert.FileExists(t, file.Source)
	assert.FileExists(t, file.Target)

	// cleaning
	//
	cleanSymlinks(env.Config.Target)
}

// copy files

func cleanSymlinks(path string) {
	files, _ := filepath.Glob(filepath.Clean(path + "/.*"))

	for _, file := range files {
		_, err := os.Readlink(file)

		if err == nil {
			os.Remove(file)
		}
	}
}
