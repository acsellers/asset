package asset

import (
	"html/template"
	"strings"
)

var RequiredFuncs = template.FuncMap{
	"assetCheckCache": func(n string) bool {
		v, ok := finishedAssets[n]
		if DevelopmentMode {
			return false
		} else {
			if ok && v {
				return true
			}
			finishedAssets[n] = false
			//go startCache(n)
			return false
		}
	},
	"assetLoad": func(t, n string) template.HTML {
		if DevelopmentMode {
			if asset, ok := developmentCssAssets[cleanExts(n)]; ok {
				return template.HTML(asset)
			}
			if asset, ok := developmentJsAssets[cleanExts(n)]; ok {
				return template.HTML(asset)
			}
		} else {
			if asset, ok := productionCssAssets[cleanExts(n)]; ok {
				return template.HTML(asset)
			}
			if asset, ok := productionJsAssets[cleanExts(n)]; ok {
				return template.HTML(asset)
			}
		}
		return template.HTML("")
	},
	"assetCssLoad": func(n string) template.CSS {
		if DevelopmentMode {
			if asset, ok := developmentCssAssets[cleanExts(n)]; ok {
				return template.CSS(asset)
			}
		} else {
			if asset, ok := productionCssAssets[cleanExts(n)]; ok {
				return template.CSS(asset)
			}
		}
		return template.CSS("")
	},
	"assetJsLoad": func(n string) template.JS {
		if DevelopmentMode {
			if asset, ok := developmentJsAssets[cleanExts(n)]; ok {
				return template.JS(asset)
			}
		} else {
			if asset, ok := productionJsAssets[cleanExts(n)]; ok {
				return template.JS(asset)
			}
		}
		return template.JS("")
	},
	"assetTree": func(t, n string) template.HTML {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentJsAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		} else {
			for name, asset := range productionJsAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		}

		return template.HTML(output)
	},
	"assetJsTree": func(n string) template.CSS {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentJsAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
			for name, asset := range developmentCssAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		} else {
			for name, asset := range productionCssAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}

			for name, asset := range productionJsAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		}

		return template.CSS(output)
	},
	"assetCssTree": func(n string) template.JS {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentCssAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		} else {
			for name, asset := range productionCssAssets {
				if strings.HasPrefix(name, n+"/") {
					output += asset
				}
			}
		}
		return template.JS(output)
	},
	"assetDir": func(t, n string) template.HTML {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentJsAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {

					output += asset
				}
			}
			for name, asset := range developmentCssAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}
		} else {
			for name, asset := range productionCssAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}

			for name, asset := range productionJsAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}
		}

		return template.HTML(output)
	},
	"assetJsDir": func(n string) template.CSS {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentJsAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {

					output += asset
				}
			}
		} else {

			for name, asset := range productionJsAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}
		}

		return template.CSS(output)
	},
	"assetCssDir": func(n string) template.JS {
		var output string
		if DevelopmentMode {
			for name, asset := range developmentCssAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}
		} else {
			for name, asset := range productionCssAssets {
				if strings.HasPrefix(name, n+"/") && !strings.Contains(name[len(n):], "/") {
					output += asset
				}
			}
		}

		return template.JS(output)
	},
}

func cleanExts(s string) string {
	return s
}
