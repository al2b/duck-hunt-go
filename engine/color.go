package engine

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-ciede2000"
	"image/color"
	"iter"
	"maps"
	"math"
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

func NewColorBinder(binding iter.Seq2[color.Color, color.Color]) *ColorBinder {
	return &ColorBinder{
		binding: binding,
		cache:   map[color.Color]color.Color{},
	}
}

type ColorBinder struct {
	binding iter.Seq2[color.Color, color.Color]
	cache   map[color.Color]color.Color
}

func (b *ColorBinder) Bind(in color.Color) (out color.Color) {
	var ok bool
	if out, ok = b.cache[in]; !ok {
		dist := math.MaxFloat64
		for i, o := range b.binding {
			if d := ciede2000.Diff(in, i); d < dist {
				dist, out = d, o
			}
		}
		b.cache[in] = out
	}

	return out
}

func ColorBindingSystem() iter.Seq2[color.Color, color.Color] {
	return maps.All(map[color.Color]color.Color{
		color.RGBA{R: 0, G: 0, B: 0}:       ansi.Black,
		color.RGBA{R: 128, G: 0, B: 0}:     ansi.Red,
		color.RGBA{R: 0, G: 128, B: 0}:     ansi.Green,
		color.RGBA{R: 128, G: 128, B: 0}:   ansi.Yellow,
		color.RGBA{R: 0, G: 0, B: 128}:     ansi.Blue,
		color.RGBA{R: 128, G: 0, B: 128}:   ansi.Magenta,
		color.RGBA{R: 0, G: 128, B: 128}:   ansi.Cyan,
		color.RGBA{R: 192, G: 192, B: 192}: ansi.White,
		color.RGBA{R: 128, G: 128, B: 128}: ansi.BrightBlack,
		color.RGBA{R: 255, G: 0, B: 0}:     ansi.BrightRed,
		color.RGBA{R: 0, G: 255, B: 0}:     ansi.BrightGreen,
		color.RGBA{R: 255, G: 255, B: 0}:   ansi.BrightYellow,
		color.RGBA{R: 0, G: 0, B: 255}:     ansi.BrightBlue,
		color.RGBA{R: 255, G: 0, B: 255}:   ansi.BrightMagenta,
		color.RGBA{R: 0, G: 255, B: 255}:   ansi.BrightCyan,
		color.RGBA{R: 255, G: 255, B: 255}: ansi.BrightWhite,
	})
}

func ColorBindingExtended() iter.Seq2[color.Color, color.Color] {
	return func(yield func(color.Color, color.Color) bool) {
		// Extended
		for i := 0; i < 216; i++ {
			if !yield(
				color.RGBA{
					R: uint8(((i / 36) * 40) + (55 * (((i / 36) + 5) / 6))),
					G: uint8((((i / 6) % 6) * 40) + (55 * ((((i / 6) % 6) + 5) / 6))),
					B: uint8(((i % 6) * 40) + (55 * (((i % 6) + 5) / 6))),
				},
				ansi.ExtendedColor(16+i),
			) {
				return
			}
		}
		// Grayscale
		for i := 0; i < 24; i++ {
			g := uint8(8 + i*10)
			if !yield(
				color.RGBA{R: g, G: g, B: g},
				ansi.ExtendedColor(232+i),
			) {
				return
			}
		}
	}
}

func ColorBindingGrayscale() iter.Seq2[color.Color, color.Color] {
	return func(yield func(color.Color, color.Color) bool) {
		if !yield(
			color.RGBA{R: 0, G: 0, B: 0},
			ansi.ExtendedColor(16),
		) {
			return
		}
		for i := 1; i < 6; i++ {
			g := uint8((i * 40) + 55)
			if !yield(
				color.RGBA{R: g, G: g, B: g},
				ansi.ExtendedColor(16+(i*43)),
			) {
				return
			}
		}
		for i := 0; i < 24; i++ {
			g := uint8(8 + i*10)
			if !yield(
				color.RGBA{R: g, G: g, B: g},
				ansi.ExtendedColor(232+i),
			) {
				return
			}
		}
	}
}

func ColorBindingMonochrome() iter.Seq2[color.Color, color.Color] {
	return maps.All(map[color.Color]color.Color{
		color.RGBA{R: 0, G: 0, B: 0}:       ansi.ExtendedColor(16),
		color.RGBA{R: 255, G: 255, B: 255}: ansi.ExtendedColor(231),
	})
}
