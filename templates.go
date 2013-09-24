package asset

import (
	"path"
	"strings"
	"text/template/parse"
)

func Parse(name, content string) (map[string]*parse.Tree, error) {
	reqs, self := parseSheet(content)
	return map[string]*parse.Tree{
		cleanName(name): compileSheet(name, reqs, self),
	}, nil
}

func cleanName(n string) string {
	if path.Ext(n) == ".asset" {
		return n[:len(n)-len(".asset")]
	} else {
		i := strings.Index(n, ".asset")
		if i > 0 {
			return n[:i] + n[i+len(".asset"):]
		}
	}
	return n
}
