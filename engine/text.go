package engine

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"image"
	"image/color"
	"image/draw"
	"io/fs"
	"strings"
)

type Font interface {
	Render(r rune, c color.Color) *Image
}

type Text struct {
	Content string
	Font    Font
	Color   color.Color
}

func (text Text) Image() *Image {
	var lines []textLine
	textWidth, textHeight := 0, 0

	for _, lineString := range strings.Split(text.Content, "\n") {
		var line textLine
		lineWidth := 0
		for _, charRune := range []rune(lineString) {
			char := text.Font.Render(charRune, text.Color)
			lineWidth += char.Bounds().Dx()
			line.height = max(line.height, char.Bounds().Dy())
			line.chars = append(line.chars, char)
		}
		lines = append(lines, line)
		textWidth = max(textWidth, lineWidth)
		textHeight += line.height
	}

	img := NewImage(Size{textWidth, textHeight})

	point := image.Point{}
	for _, line := range lines {
		point.X = 0
		for _, char := range line.chars {
			charBounds := char.Bounds()
			draw.Draw(
				img.NRGBA,
				image.Rectangle{
					Min: point,
					Max: point.Add(charBounds.Size()),
				},
				char,
				image.Point{},
				draw.Over,
			)
			point.X += charBounds.Dx()
		}
		point.Y += line.height
	}

	return img
}

type textLine struct {
	chars  []*Image
	height int
}

/**********/
/* Bitmap */
/**********/

type BitmapFont struct {
	Map     *image.Alpha
	Mapper  BitmapFontMapper
	Charmap *charmap.Charmap
}

func (font BitmapFont) Render(r rune, c color.Color) *Image {
	char, _ := font.Charmap.EncodeRune(r)
	mask := font.Mapper.Mask(font.Map, char)
	maskBounds := mask.Bounds()

	imgSize := maskBounds.Size()
	img := NewImage(Size{Width: imgSize.X, Height: imgSize.Y})

	draw.DrawMask(
		img.NRGBA,
		img.Bounds(),
		&image.Uniform{C: c}, image.Point{},
		mask, maskBounds.Min,
		draw.Over,
	)

	return img
}

type BitmapFontMapper interface {
	Mask(src *image.Alpha, char byte) *image.Alpha
}

type SquareBitmapFontMapper struct{}

func (mapper SquareBitmapFontMapper) Mask(src *image.Alpha, char byte) *image.Alpha {
	size := src.Bounds().Size().Div(16)
	point := image.Point{int(char%16) * size.X, int(char/16) * size.Y}
	return src.SubImage(image.Rectangle{
		Min: point,
		Max: point.Add(image.Point{size.X, size.Y}),
	}).(*image.Alpha)
}

type BitmapFontLoader struct {
	FS      fs.ReadFileFS
	Path    string
	Mapper  BitmapFontMapper
	Charmap *charmap.Charmap
}

func (loader BitmapFontLoader) Load() (*BitmapFont, error) {
	var (
		file fs.File
		img  image.Image
		err  error
	)

	if file, err = loader.FS.Open(loader.Path); err != nil {
		return nil, err
	}
	defer file.Close()

	if img, _, err = image.Decode(file); err != nil {
		return nil, err
	}

	font := &BitmapFont{
		Mapper:  loader.Mapper,
		Charmap: loader.Charmap,
	}

	switch img := img.(type) {
	case *image.Gray:
		font.Map = &image.Alpha{
			Pix:    img.Pix,
			Stride: img.Stride,
			Rect:   img.Rect,
		}
	default:
		return nil, fmt.Errorf("unsupported bitmap font image type")
	}

	return font, nil
}
