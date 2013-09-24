package asset

import (
	"text/template/parse"
)

var (
	developmentCssAssets = make(map[string]string)
	productionCssAssets  = make(map[string]string)

	developmentJsAssets = make(map[string]string)
	productionJsAssets  = make(map[string]string)

	cachedAssetList = make(map[string]*parse.TextNode)
	finishedAssets  = make(map[string]bool)
)
