//nolint:dupl // lines are dulicate to ease reading
package util_test

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Example usage
// ----------------------------------------------------------------------------

func ExampleGetMods() {
	listMods := util.GetMods()

	for _, modInfo := range listMods {
		fmt.Println(modInfo["name"])
		fmt.Println(modInfo["path"])
		fmt.Println(modInfo["version"])
	}
	// Output:
	// a
	// n/a
	// n/a
}

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func TestGetMods(t *testing.T) {
	oldReadBuildInfo := util.ReadBuildInfo
	defer func() {
		util.ReadBuildInfo = oldReadBuildInfo
	}()

	util.ReadBuildInfo = func() (info *debug.BuildInfo, ok bool) {
		i := new(debug.BuildInfo)

		i.Deps = append(i.Deps, &debug.Module{
			Path: "foo/bar", Version: "buz",
		})

		return i, true
	}

	listMods := util.GetMods()

	assert.Equal(t, "foo/bar", listMods[0]["path"])
	assert.Equal(t, "bar", listMods[0]["name"])
	assert.Equal(t, "buz", listMods[0]["version"])
}

func TestGetMods_issue1(t *testing.T) {
	oldReadBuildInfo := util.ReadBuildInfo
	defer func() {
		util.ReadBuildInfo = oldReadBuildInfo
	}()

	util.ReadBuildInfo = func() (info *debug.BuildInfo, ok bool) {
		i := new(debug.BuildInfo)

		i.Deps = append(i.Deps, &debug.Module{
			Path: "github.com/jedib0t/go-pretty/v6", Version: "v6",
		})

		return i, true
	}

	listMods := util.GetMods()

	assert.Equal(t, "github.com/jedib0t/go-pretty/v6", listMods[0]["path"])
	assert.Equal(t, "go-pretty/v6", listMods[0]["name"])
	assert.Equal(t, "v6", listMods[0]["version"])
}
