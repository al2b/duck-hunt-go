package font

import (
	"duck-hunt-go/engine"
	"embed"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Font
	Font = engine.Must(engine.LoadFont(assets, "assets/font.png"))
)
