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
		AssetDigest("template/" + a)
		info.IsDir()
		info.ModTime()
		info.Mode()
		info.Name()
		info.Size()
		info.Sys()
	}

	_, err = AssetDir("x")
	assert.Error(t, err, "Should have an error for a missing dir")

	_, err = AssetDigest("x")
	assert.Error(t, err, "Should have an error for a missing file")

	digests, derr := Digests()
	assert.NoError(t, derr, "Should not have an error getting digests")
	assert.Len(t, digests, 20, "Did you add or remove an template?")

	_, err = AssetInfo("x")
	assert.Error(t, err, "Should have an error for a missing dir")
	assert.Panics(t, func() { MustAsset("x") })
	assert.Panics(t, func() { MustAssetString("x") })

	RestoreAssets(os.TempDir(), "x")
	tmpDir := os.TempDir()
	err = RestoreAssets(tmpDir, "template")
	if err != nil {
		t.Logf("Error restoring assets: %s", err)
	}
}
