package engine

import (
	"fmt"
	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-ciede2000"
	"image/color"
	"iter"
	"maps"
	"math"
)

var (
	ColorBlack = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	ColorWhite = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	ColorRed   = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	ColorGreen = color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}
	ColorBlue  = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
)

func NewColorBinder(binding ColorBinding) *ColorBinder {
	return &ColorBinder{
		Binding: binding,
		cache:   map[color.Color]color.Color{},
	}
}

type ColorBinder struct {
	Binding ColorBinding
	cache   map[color.Color]color.Color
}

func (b *ColorBinder) Bind(in color.Color) (out color.Color) {
	var ok bool
	if out, ok = b.cache[in]; !ok {
		dist := math.MaxFloat64
		for i, o := range b.Binding.All() {
			if d := ciede2000.Diff(in, i); d < dist {
				dist, out = d, o
			}
		}
		b.cache[in] = out
	}

	return
}

type ColorBinding interface {
	fmt.Stringer
	All() iter.Seq2[color.Color, color.Color]
}

type ColorBindingANSI256 struct{}

func (c ColorBindingANSI256) String() string { return "ANSI256" }
func (c ColorBindingANSI256) All() iter.Seq2[color.Color, color.Color] {
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

type ColorBindingANSI256Grayscale struct{}

func (c ColorBindingANSI256Grayscale) String() string { return "ANSI256 - Grayscale" }
func (c ColorBindingANSI256Grayscale) All() iter.Seq2[color.Color, color.Color] {
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

type ColorBindingANSI256BlackAndWhite struct{}

func (c ColorBindingANSI256BlackAndWhite) String() string { return "ANSI256 - Black And White" }
func (c ColorBindingANSI256BlackAndWhite) All() iter.Seq2[color.Color, color.Color] {
	return maps.All(map[color.Color]color.Color{
		color.RGBA{R: 0, G: 0, B: 0}:       ansi.ExtendedColor(16),
		color.RGBA{R: 255, G: 255, B: 255}: ansi.ExtendedColor(231),
	})
}

type ColorBindingANSI struct{}

func (c ColorBindingANSI) String() string { return "ANSI" }
func (c ColorBindingANSI) All() iter.Seq2[color.Color, color.Color] {
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

type ColorBindingANSIGrayscale struct{}

func (c ColorBindingANSIGrayscale) String() string { return "ANSI - Grayscale" }
func (c ColorBindingANSIGrayscale) All() iter.Seq2[color.Color, color.Color] {
	return maps.All(map[color.Color]color.Color{
		color.RGBA{R: 0, G: 0, B: 0}:       ansi.Black,
		color.RGBA{R: 192, G: 192, B: 192}: ansi.White,
		color.RGBA{R: 128, G: 128, B: 128}: ansi.BrightBlack,
		color.RGBA{R: 255, G: 255, B: 255}: ansi.BrightWhite,
	})
}

type ColorBindingANSIBlackAndWhite struct{}

func (c ColorBindingANSIBlackAndWhite) String() string { return "ANSI - Black And White" }
func (c ColorBindingANSIBlackAndWhite) All() iter.Seq2[color.Color, color.Color] {
	return maps.All(map[color.Color]color.Color{
		color.RGBA{R: 0, G: 0, B: 0}:       ansi.Black,
		color.RGBA{R: 255, G: 255, B: 255}: ansi.BrightWhite,
	})
}
