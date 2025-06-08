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

var (
	Font5x5 = Must(BitmapFontLoader{
		assets, "assets/font.5x5.png",
		SquareBitmapFontMaskMapper{}, charmap.CodePage437,
	}.Load())
	Font6x6 = Must(BitmapFontLoader{
		assets, "assets/font.6x6.png",
		SquareBitmapFontMaskMapper{}, charmap.CodePage437,
	}.Load())
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
	Mask       *image.Alpha
	MaskMapper BitmapFontMaskMapper
	Charmap    *charmap.Charmap
}

func (font BitmapFont) Render(r rune, c color.Color) *Image {
	char, _ := font.Charmap.EncodeRune(r)
	mask := font.MaskMapper.Map(font.Mask, char)
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

type BitmapFontMaskMapper interface {
	Map(mask *image.Alpha, char byte) *image.Alpha
}

type SquareBitmapFontMaskMapper struct{}

func (mapper SquareBitmapFontMaskMapper) Map(mask *image.Alpha, char byte) *image.Alpha {
	size := mask.Bounds().Size()
	width, height := size.X/16, size.Y/16
	point := image.Pt(int(char%16)*width, int(char/16)*height)
	return mask.SubImage(image.Rectangle{
		Min: point,
		Max: point.Add(image.Pt(width, height)),
	}).(*image.Alpha)
}

type BitmapFontLoader struct {
	FS      fs.ReadFileFS
	Path    string
	Mapper  BitmapFontMaskMapper
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
		MaskMapper: loader.Mapper,
		Charmap:    loader.Charmap,
	}

	switch img := img.(type) {
	case *image.Gray:
		font.Mask = &image.Alpha{
			Pix:    img.Pix,
			Stride: img.Stride,
			Rect:   img.Rect,
		}
	default:
		return nil, fmt.Errorf("unsupported bitmap font image type")
	}

	return font, nil
}
