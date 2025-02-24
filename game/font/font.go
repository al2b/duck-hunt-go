package font

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed assets/*.png
var assets embed.FS

var Font = engine.MustLoadFont(engine.FontImageFile(assets, "assets/font.png"))
