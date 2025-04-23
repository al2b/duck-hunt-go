package title

import (
	"duck-hunt-go/engine"
	"embed"
	"image/color"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Colors
	colorText = color.RGBA{R: 0xff, G: 0xa0, B: 0x00}

	// Images
	imageLayout = engine.Must(engine.LoadImage(assets, "assets/layout.png"))
	imageCursor = engine.Must(engine.LoadImage(assets, "assets/cursor.png"))
)
