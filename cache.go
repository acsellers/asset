package asset

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template/parse"

	//"github.com/robfig/revel"
)

var (
	developmentCssAssets = make(map[string]string)
	productionCssAssets  = make(map[string]string)

	developmentJsAssets = make(map[string]string)
	productionJsAssets  = make(map[string]string)

	cachedAssetList = make(map[string]*parse.TextNode)
	finishedAssets  = make(map[string]bool)

	BasePath string
)

func LoadCache() {
	for _, location := range AssetLocations {
		currentLocation := filepath.Join(BasePath, location)
		filepath.Walk(currentLocation, func(p string, info os.FileInfo, err error) error {
			if info != nil && !info.IsDir() {
				content, err := ioutil.ReadFile(p)
				if err == nil {
					name := p[len(currentLocation)+1:]
					ct := contentTypeFor(name)
					name = cleanExts(name)
					switch ct {
					case "js":
						developmentJsAssets[name] = string(content)
						productionJsAssets[name] = string(content)
					case "css":
						developmentCssAssets[name] = string(content)
						productionCssAssets[name] = string(content)
					}
				}
			}
			return nil
		})
	}
}
