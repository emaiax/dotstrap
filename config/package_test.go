package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadValidGlobPackages(t *testing.T) {
	env, _ := Load("testdata/packages.validglob.yml")

	assert.Contains(
		t,
		env.Packages["mypackage"].Files,
		PackageFile{
			Name:   "packages.validglob.yml",
			Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.validglob.yml",
			Target: "/root/.packages.validglob.yml",
		},
	)
}

func TestLoadInvalidGlobPackages(t *testing.T) {
	env, _ := Load("testdata/packages.invalidglob.yml")

	assert.Empty(t, env.Packages["mypackage"].Files)
}

func TestLoadLinksPackages(t *testing.T) {
	env, _ := Load("testdata/packages.links.yml")

	assert.ElementsMatch(
		t,
		env.Packages["mypackage"].Files,
		[]PackageFile{
			PackageFile{
				Name:   "packages.glob.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.glob.yml",
				Target: "/root/.globpack",
			},
			PackageFile{
				Name:   "packages.links.yml",
				Source: "/go/src/github.com/emaiax/dotstrap/config/testdata/packages.links.yml",
				Target: "/root/.linkpack",
			},
		},
	)
}

func TestResolveEnvVarToValue(t *testing.T) {
	path := resolveEnvVar("${HOME}", os.UserHomeDir)

	assert.Equal(t, path, "/root")
}

func TestResolveEnvVarToDefault(t *testing.T) {
	path := resolveEnvVar("${NOPE}", os.Getwd)
	pwd, _ := os.Getwd()

	assert.Equal(t, path, pwd)
}

func TestPublicPath(t *testing.T) {
	assert.Equal(t, getPublicPath("/go/", "/xpto/file"), "/go/xpto/file")
}

func TestPrivatePath(t *testing.T) {
	assert.Equal(t, getHiddenPath("/go/", "/xpto/.DS_Store"), "/go/xpto/.DS_Store")
}

func TestInstallStatusEmptyPackage(t *testing.T) {
	pack := Package{}

	assert.Equal(t, pack.InstallStatus(), FullyInstalled)
}

func TestInstallStatusNotInstalledFile(t *testing.T) {
	pack := Package{
		Files: []PackageFile{
			PackageFile{Installed: false},
		},
	}

	assert.Equal(t, pack.InstallStatus(), NotInstalled)
}

func TestInstallStatusPartiallyInstalledFiles(t *testing.T) {
	pack := Package{
		Files: []PackageFile{
			PackageFile{Installed: true},
			PackageFile{Installed: false},
		},
	}

	assert.Equal(t, pack.InstallStatus(), PartiallyInstalled)
}

func TestInstallStatusFullyInstalledFiles(t *testing.T) {
	pack := Package{
		Files: []PackageFile{
			PackageFile{Installed: true},
		},
	}

	assert.Equal(t, pack.InstallStatus(), FullyInstalled)
}
