// Silly Test to get coverage up

package generator

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	assets, err := AssetDir("template")
	assert.NoError(t, err, "Error getting template dir")

	for _, a := range assets {
		fmt.Printf("%s\n", a)
		info, _ := AssetInfo("template/" + a)
		info.IsDir()
		info.ModTime()
		info.Mode()
		info.Name()
		info.Size()
		info.Sys()
	}

	_, err = AssetDir("x")
	assert.Error(t, err, "Should have an error for a missing dir")

	_, err = AssetInfo("x")
	assert.Error(t, err, "Should have an error for a missing dir")
	assert.Panics(t, func() { MustAsset("x") })

	RestoreAssets(os.TempDir(), "x")
	RestoreAssets(os.TempDir(), "template")
}
