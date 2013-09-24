package asset

import (
	"github.com/acsellers/assert"
	"testing"
)

func TestCleanExts(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		test.AreEqual("one", cleanExts("one.ts.js"))
		test.AreEqual("one", cleanExts("one.coffee.js"))
		test.AreEqual("one", cleanExts("one.coffee"))
		test.AreEqual("one", cleanExts("one.cs"))
		test.AreEqual("one", cleanExts("one.css"))
	})
}

func TestCleanName(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		test.AreEqual("one.css", cleanName("one.asset.css"))
		test.AreEqual("one.js", cleanName("one.asset.js"))
	})
}
