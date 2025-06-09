package engine

import (
	"embed"
	"golang.org/x/text/encoding/charmap"
)

var (
	//go:embed assets/*.png
	assets embed.FS

	// Fonts
	Font5x5 = MustLoad[*BitmapFont](BitmapFontLoader{
		assets, "assets/font.5x5.png",
		SquareBitmapFontMapper{}, charmap.CodePage437,
	})
	Font6x6 = MustLoad[*BitmapFont](BitmapFontLoader{
		assets, "assets/font.6x6.png",
		SquareBitmapFontMapper{}, charmap.CodePage437,
	})
)
