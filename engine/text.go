package engine

import (
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	"strings"
)

const (
	charWidth  = 5
	charHeight = 5
)

var (
	imageFont5x5 = Must(
		LoadImageFile(imagesFS, "images/font.5x5.png"),
	)
)

func NewText(text string) Text {
	return Text{
		text: text,
	}
}

type Text struct {
	text string
}

func (t Text) Image() image.Image {
	lines := strings.Split(t.text, "\n")

	// Compute dimensions
	var width, height int
	for _, line := range lines {
		w := len(line)
		if w > width {
			width = w
		}
	}
	height = len(lines)

	img := image.NewNRGBA(image.Rect(0, 0, width*charWidth, height*charHeight))
	fontImg := imageFont5x5

	for l, line := range lines {
		for c, char := range line {
			charIndex := int(char)
			charX := (charIndex % 16) * charWidth
			charY := (charIndex / 16) * charHeight
			charImg := imaging.Crop(fontImg, image.Rect(
				charX,
				charY,
				charX+charWidth,
				charY+charHeight,
			))
			draw.Draw(img, img.Bounds(), charImg, image.Point{
				X: -c * charWidth,
				Y: -l * charHeight,
			}, draw.Over)
		}
	}

	return img
}
