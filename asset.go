// Asset is an asset pipeline for revel (maybe possibly in the future)
// Assets have the extension .asset, for instance a javascript
// file would have the extension application.js.asset.
package asset

/*
  When DevelopmentMode is set to true, minified versions of assets will
  not be preferred. When it is set to false, assets with a .min extension
  will be preferred and if there are minifiers registered, they will be
  used on the relevant assets.
*/
var (
	DevelopmentMode bool
	AssetLocations  = []string{
		"public/",
		"app/assets",
	}
)
