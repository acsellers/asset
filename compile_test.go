package asset

import (
	"bytes"
	"github.com/acsellers/assert"
	"html/template"
	"strings"
	"testing"
)

func newTemplate() (*template.Template, map[string]string) {
	assets := map[string]string{}

	return template.New("compile").Funcs(map[string]interface{}{
		"assetCheckCache": func(n string) bool { return false },
		"assetLoad":       func(t, n string) template.HTML { return template.HTML(assets[t+":"+n]) },
		"assetCssLoad":    func(n string) template.CSS { return template.CSS(assets["css:"+n]) },
		"assetJsLoad":     func(n string) template.JS { return template.JS(assets["js:"+n]) },
		"assetTree": func(t, n string) template.HTML {
			var output string
			for name, asset := range assets {
				if strings.HasPrefix(name, t+":"+n+"/") {
					output += asset
				}
			}
			return template.HTML(output)
		},
		"assetJsTree": func(n string) template.CSS {
			var output string
			for name, asset := range assets {
				if strings.HasPrefix(name, "js:"+n+"/") {
					output += asset
				}
			}
			return template.CSS(output)
		},
		"assetCssTree": func(n string) template.JS {
			var output string
			for name, asset := range assets {
				if strings.HasPrefix(name, "css:"+n+"/") {
					output += asset
				}
			}
			return template.JS(output)
		},
	}), assets
}

func TestCompile1(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet(".name {\ncolor: #ffffff;\n}")
		test.AreEqual(0, len(reqs))
		test.AreEqual(3, len(content))

		tmpl, _ := newTemplate()
		b := new(bytes.Buffer)
		tree := compileSheet("compile", reqs, content)
		tmpl.AddParseTree("compile", tree)
		test.IsNil(tmpl.ExecuteTemplate(b, "compile", nil))
		test.AreEqual(".name {\ncolor: #ffffff;\n}", b.String())
	})
}

func TestCompile2(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet("//= require_self\n.name {\ncolor: #ffffff;\n}")
		test.AreEqual(1, len(reqs))
		test.AreEqual(3, len(content))

		tmpl, _ := newTemplate()
		b := new(bytes.Buffer)
		tree := compileSheet("compile", reqs, content)
		tmpl.AddParseTree("compile", tree)
		test.IsNil(tmpl.ExecuteTemplate(b, "compile", nil))
		test.AreEqual(".name {\ncolor: #ffffff;\n}", b.String())
	})
}

func TestCompile3(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet("//= require 'test'\n.name {\ncolor: #ffffff;\n}")
		test.AreEqual(1, len(reqs))
		test.AreEqual(3, len(content))

		tmpl, assets := newTemplate()
		assets[":test"] = "h1 { color: #afafaf }"
		b := new(bytes.Buffer)
		tree := compileSheet("compile", reqs, content)
		tmpl.AddParseTree("compile", tree)
		test.IsNil(tmpl.ExecuteTemplate(b, "compile", nil))
		test.AreEqual("h1 { color: #afafaf }.name {\ncolor: #ffffff;\n}", b.String())
	})
}

func TestCompile4(t *testing.T) {
	assert.Within(t, func(test *assert.Test) {
		reqs, content := parseSheet("//= require_tree 'test'\n.name {\ncolor: #ffffff;\n}")
		test.AreEqual(1, len(reqs))
		test.AreEqual(3, len(content))

		tmpl, assets := newTemplate()
		assets[":test/one"] = "h1 { color: #afafaf }"
		assets[":test/two"] = "h2 { color: #afafaf }"
		b := new(bytes.Buffer)
		tree := compileSheet("compile", reqs, content)
		tmpl.AddParseTree("compile", tree)
		test.IsNil(tmpl.ExecuteTemplate(b, "compile", nil))
		test.AreEqual(
			"h1 { color: #afafaf }h2 { color: #afafaf }.name {\ncolor: #ffffff;\n}",
			b.String(),
		)
	})
}
