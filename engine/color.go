package engine

import (
	"image/color"
)

type Color8 uint8
type Color24 color.NRGBA

var (
	// Red
	Color8Red  = Color8(196)
	Color24Red = Color24{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	// Green
	Color8Green  = Color8(46)
	Color24Green = Color24{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
	// Blue
	Color8Blue  = Color8(21)
	Color24Blue = Color24{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
)
