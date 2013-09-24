package asset

import (
	"html/template"
	"text/template/parse"
)

func Parse() (map[string]*parse.Tree, error) {
	return map[string]*parse.Tree{}, nil
}

func LoadCssAsset(name, assetName string) template.CSS {
	if DevelopmentMode {
		return developmentCssAssets[name]
	} else {
		return productionCssAssets[name]
	}
}
func LoadJsAsset(name, assetName string) template.JS {
	if DevelopmentMode {
		return developmentJsAssets[name]
	} else {
		return productionJsAssets[name]
	}

}
