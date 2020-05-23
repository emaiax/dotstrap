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
	cleanSymlinks(env.Config.Target)
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
	cleanSymlinks(env.Config.Target)
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
	cleanSymlinks(env.Config.Target)
}

func TestBackupFileSuccess(t *testing.T) {
	env, _ := config.Load("testdata/install.forcelinks.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// creates a Source backup to prevent real Source removal
	//
	fakeSource := file.Source + ".bkp"
	backupFileName := backupFileName(fakeSource)

	os.Symlink(file.Source, fakeSource)

	// backup file doesn't exist yet
	//
	assert.FileExists(t, fakeSource)
	assert.NoFileExists(t, backupFileName)

	backupFile(fakeSource, backupFileName)

	// fake source renamed to backup file
	//
	assert.NoFileExists(t, fakeSource)
	assert.FileExists(t, backupFileName)

	// cleaning
	//
	os.Remove(backupFileName)
}

func TestBackupFileAlreadyExists(t *testing.T) {
	env, _ := config.Load("testdata/install.forcelinks.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// creates a Source backup to prevent real Source removal
	//
	fakeSource := file.Source + ".bkp"
	os.Symlink(file.Source, fakeSource)

	backupFileName := backupFileName(fakeSource)

	os.Symlink(fakeSource, backupFileName)

	// backup file doesn't exist yet
	//
	assert.FileExists(t, fakeSource)
	assert.FileExists(t, backupFileName)

	newBackupFile, _ := backupFile(fakeSource, backupFileName)

	// fake source renamed to backup file
	//
	assert.FileExists(t, fakeSource)
	assert.NoFileExists(t, backupFileName)
	assert.FileExists(t, newBackupFile)

	// cleaning
	//
	os.Remove(fakeSource)
	os.Remove(newBackupFile)
}

func TestBackupFileName(t *testing.T) {
	backupFilename := backupFileName("file")

	assert.Equal(t, len(backupFilename), 24)

	otherBackupFileName := backupFileName(backupFilename)

	assert.Equal(t, len(otherBackupFileName), 24)
}

func cleanSymlinks(path string) {
	files, _ := filepath.Glob(filepath.Clean(path + "/.*"))

	for _, file := range files {
		_, err := os.Readlink(file)

		if err == nil {
			os.Remove(file)
		}
	}
}
