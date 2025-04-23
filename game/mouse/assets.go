package mouse

import (
	"duck-hunt-go/engine"
	"embed"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageMouse = engine.Must(engine.LoadImage(assets, "assets/mouse.png"))
)
