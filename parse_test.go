package asset

import (
	"strings"
	"testing"

	"github.com/acsellers/assert"
)

func TestCommandParse(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		r, err := getCommand("*= require_self")
		test.IsNil(err)
		test.AreEqual("self", r.Type)

		r, err = getCommand("//= require_self")
		test.IsNil(err)
		test.AreEqual("self", r.Type)

		r, err = getCommand("//= this shouldn't call require_self")
		test.IsNotNil(err)
		test.IsNil(r)

		r, err = getCommand("//= require_directory")
		test.IsNil(r)
		test.IsNotNil(err)

		r, err = getCommand("//= require_directory \"blah\"")
		test.IsNil(err)
		test.AreEqual(r.Type, "directory")
		test.AreEqual(r.Item, "blah")

	})
}

func TestSheetParse1(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet(`
  //= require_self
  //= require 'news'

  .name {
    text-color: "chucknorris";
  }
  `)
		test.AreEqual(2, len(reqs))
		test.AreEqual(3, len(content))
		if len(reqs) == 2 {
			test.AreEqual(reqs[0].Type, "self")
			test.AreEqual(reqs[1].Item, "news")
		}
		test.AreEqual(".name {text-color: \"chucknorris\";}", strings.Join(content, ""))
	})
}

func TestSheetParse2(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet(`/*
  *= require_self
  *= require 'news'

  .name {
    text-color: "chucknorris";
  }
  */
  `)
		test.AreEqual(2, len(reqs))
		test.AreEqual(5, len(content))
		if len(reqs) == 2 {
			test.AreEqual(reqs[0].Type, "self")
			test.AreEqual(reqs[1].Item, "news")
		}
		test.AreEqual("/*.name {text-color: \"chucknorris\";}*/", strings.Join(content, ""))
	})
}
