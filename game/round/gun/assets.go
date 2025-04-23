package gun

import (
	"duck-hunt-go/engine"
	"embed"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageGun = engine.Must(engine.LoadImage(assets, "assets/gun.png"))
)
