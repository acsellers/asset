package asset

import (
	"path"
	"strings"
	"text/template/parse"
)

func compileSheet(name string, reqs []*require, content []string) *parse.Tree {
	contentType := contentTypeFor(name)
	var elseList []parse.Node
	var selfAdded bool

	for _, req := range reqs {
		if req.Type == "self" {
			elseList = append(elseList,
				newTextNode(strings.Join(content, "\n")),
			)
			selfAdded = true
		} else {
			elseList = append(elseList, actionNodeFor(contentType, req))
		}
	}

	if !selfAdded {
		elseList = append(elseList,
			newTextNode(strings.Join(content, "\n")),
		)
	}

	return newTree(name, newListNode(elseList))
}

func contentTypeFor(s string) string {
	switch path.Ext(s) {
	case "js", "coffee", "dart", "ts":
		return "js"
	case "css", "sass", "scss", "less":
		return "css"
	}
	return ""
}
