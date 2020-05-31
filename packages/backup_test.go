package packages

import (
	"os"
	"testing"

	"github.com/emaiax/dotstrap/config"
	"github.com/stretchr/testify/assert"
)

func TestBackupFileSuccess(t *testing.T) {
	env, _ := config.Load("testdata/backup.yml")

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

func TestBackupFileInvalidSourcePath(t *testing.T) {
	fakeSource := "/invalidpath/invalid-path.file"
	backupFileName := backupFileName(fakeSource)

	// invalid file doesn't exist
	//
	assert.NoFileExists(t, fakeSource)
	assert.NoFileExists(t, backupFileName)

	newBackupFile, backupCreated := backupFile(fakeSource, backupFileName)

	assert.NoFileExists(t, fakeSource)
	assert.NoFileExists(t, backupFileName)

	assert.Empty(t, newBackupFile)
	assert.False(t, backupCreated)
}

func TestBackupFileInvalidTargetPath(t *testing.T) {
	env, _ := config.Load("testdata/backup.yml")

	pack := env.Packages["mypackage"]
	file := pack.Files[0]

	// invalid file doesn't exist
	//
	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	newBackupFile, backupCreated := backupFile(file.Source, "/invalidpath/invalid.file")

	assert.FileExists(t, file.Source)
	assert.NoFileExists(t, file.Target)

	assert.Empty(t, newBackupFile)
	assert.False(t, backupCreated)
}

func TestBackupFileAlreadyExists(t *testing.T) {
	env, _ := config.Load("testdata/backup.yml")

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

func TestBackupFilePermissionError(t *testing.T) {
	os.Mkdir("testdata/noperm", 0777)
	os.Create("testdata/noperm/file")

	assert.FileExists(t, "testdata/noperm/file")

	os.Chmod("testdata/noperm", 0000)

	assert.FileExists(t, "testdata/noperm/file")

	backupFile, backupCreated := backupFile("testdata/noperm/file", "testdata/noperm/file.bkp")

	assert.Empty(t, backupFile)
	assert.False(t, backupCreated)

	// cleaning
	//
	os.Chmod("testdata/noperm", 0777)
	os.RemoveAll("testdata/noperm")
}

func TestBackupFileName(t *testing.T) {
	backupFilename := backupFileName("file")

	assert.Equal(t, len(backupFilename), 24)

	otherBackupFileName := backupFileName(backupFilename)

	assert.Equal(t, len(otherBackupFileName), 24)
}
