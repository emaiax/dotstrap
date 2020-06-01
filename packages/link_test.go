package packages

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	noPerms   = 000
	fullPerms = 777
)

func TestLinkFileSuccess(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "testfile.link"

	os.Create(source)

	// target file doesn't exist yet
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	assert.True(t, linkFile("testfile", source, target))

	// now both files exist
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	// teardown
	//
	os.Remove(source)
	os.Remove(target)
}

func TestLinkFileBackupSuccess(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "testfile.link"

	os.Create(source)
	os.Create(target)

	// target file doesn't exist yet
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	assert.True(t, linkFile("testfile", source, target))

	// now both files exist
	//
	assert.FileExists(t, source)
	assert.FileExists(t, target)

	// teardown
	//
	os.Remove(source)
	os.Remove(target)

	cleanLinkBackups()
}

func TestLinkFileInvalidTargetError(t *testing.T) {
	// setup
	//
	source := "testfile"
	target := "invalid/testfile.link"

	os.Create(source)

	// target file exists
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	assert.False(t, linkFile("testfile", source, target))

	// target file still doesn't exist
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	// teardown
	//
	os.Remove(source)
}

func TestLinkFilePermissionError(t *testing.T) {
	// setup
	//
	os.Mkdir("noperm", fullPerms)

	source := "testfile"
	target := "noperm/testfile.link"

	os.Create(source)

	// target file doesn't exist yet
	//
	assert.FileExists(t, source)
	assert.NoFileExists(t, target)

	os.Chmod("noperm", noPerms)

	assert.False(t, linkFile("testfile", source, target))

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

func cleanLinkBackups() {
	files, _ := filepath.Glob("*.link.*")

	for _, file := range files {
		os.Remove(filepath.Clean(file))
	}
}
