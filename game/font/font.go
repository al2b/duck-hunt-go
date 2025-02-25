package font

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed assets/*.png
var assets embed.FS

var Font = engine.Must(engine.LoadFont(assets, "assets/font.png"))
