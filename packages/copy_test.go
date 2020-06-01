package packages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFileSuccess(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "testfile.copy"

	os.Create(source)

	// target file doesn't exist yet
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	assert.True(t, copyFile("testfile", source, target))

	// now both files exist
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	// teardown
	//
	os.Remove(source)
	os.Remove(target)
}

func TestCopyFileBackupSuccess(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "testfile.copy"

	os.Create(source)
	os.Create(target)

	// target file exists
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	assert.True(t, copyFile("testfile", source, target))

	// now both files exist
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	// teardown
	//
	os.Remove(source)
	os.Remove(target)

	cleanCopyBackups()
}

func TestCopyFileInvalidSourceError(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "testfile.copy"

	// no file exists
	//
	assert.NoFileExists(t, source)
	assert.NoFileExists(t, target)

	assert.False(t, copyFile("testfile", source, target))

	// now both files exist
	//
	assert.NoFileExists(t, source)
	assert.NoFileExists(t, target)
}

func TestCopyFileInvalidTargetError(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "invalid/testfile.copy"

	os.Create(source)

	// no file exists
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	assert.False(t, copyFile("testfile", source, target))

	// now both files exist
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	// teardown
	//
	os.Remove(source)
}

func TestCopyFilePermissionError(t *testing.T) {
	// setup
	//
	os.Mkdir("noperm", fullPerms)

	source := "testfile"
	target := "noperm/testfile.copy"

	os.Create(source)

	// target file doesn't exist yet
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	os.Chmod("noperm", noPerms)

	assert.False(t, copyFile("testfile", source, target))

	// target file still doesn't exist
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	// teardown
	//
	os.Chmod("noperm", fullPerms)
	os.RemoveAll("noperm")

	os.Remove(source)
}

func cleanCopyBackups() {
	files, _ := filepath.Glob("*.copy.*")

	for _, file := range files {
		os.Remove(filepath.Clean(file))
	}
}
