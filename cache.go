package asset

import (
	"html/template"
	"text/template/parse"
)

var (
	developmentCssAssets = make(map[string]template.CSS)
	productionCssAssets  = make(map[string]template.CSS)

	developmentJsAssets = make(map[string]template.JS)
	productionJsAssets  = make(map[string]template.JS)
	cachedAssetList     = make(map[string]*parse.TextNode)
)
