package layout

import (
	"duck-hunt-go/engine"
	"embed"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageLayout = engine.Must(engine.LoadImage(assets, "assets/layout.png"))
	imageShrub  = engine.Must(engine.LoadImage(assets, "assets/shrub.png"))
	imageTree   = engine.Must(engine.LoadImage(assets, "assets/tree.png"))
)
