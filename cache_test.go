package asset

import (
	"github.com/acsellers/assert"
	"testing"
)

func TestCache(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		BasePath = "/home/andrew/Projects/go/src/github.com/acsellers/asset"
		AssetLocations = []string{"samples"}
		LoadCache()
		test.AreEqual(
			"// jquery.js\n",
			developmentJsAssets["jquery"],
		)
	})
}
