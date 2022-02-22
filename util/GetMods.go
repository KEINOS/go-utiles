package util

import (
	"path/filepath"
	"runtime/debug"
	"strings"
)

// ReadBuildInfo is a copy of debug.ReadBuildInfo to ease mocking during test
// for GetMods.
//
// This package uses this util.ReadBuildInfo insetead of debug.ReadBuildInfo.
var ReadBuildInfo = debug.ReadBuildInfo

// GetMods returns a list of external modules used in the package.
// The list contains: module name, path and the version.
func GetMods() []map[string]string {
	mods := getModuleInfo()
	modsFound := []map[string]string{}

	if len(mods) == 0 {
		dummyMod := &debug.Module{
			Path:    "n/a",
			Version: "n/a",
			Sum:     "n/a",
		}
		mods = []*debug.Module{
			dummyMod,
		}
	}

	for _, modDep := range mods {
		modsFound = append(modsFound, map[string]string{
			"name":    getModName(modDep),
			"path":    modDep.Path,
			"version": modDep.Version,
		})
	}

	return modsFound
}

func getModName(modDep *debug.Module) string {
	result := filepath.Base(modDep.Path)

	if chunk := strings.SplitN(modDep.Path, "/", 3); len(chunk) > 2 {
		result = chunk[2] // Get the 3rd chunk
	}

	return result
}

func getModuleInfo() []*debug.Module {
	debugModules := []*debug.Module{}

	if buildInfo, ok := ReadBuildInfo(); ok {
		debugModules = buildInfo.Deps
	}

	return debugModules
}
