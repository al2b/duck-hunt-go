package engine

import (
	"image"
	"strings"
)

const (
	charWidth  = 5
	charHeight = 5
)

var fontImage = Must(LoadImageFiles(imagesFS,
	"images/font.8.png",
	"images/font.24.png",
))

func NewTextImage(text string) TextImage {
	return TextImage{
		text: text,
	}
}

type TextImage struct {
	text string
}

func (t TextImage) Image8() Image8 {
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

	img := NewImage8(image.Rect(0, 0, width*charWidth, height*charHeight))
	fontImg := fontImage.Image8()

	for l, line := range lines {
		for c, char := range line {
			charIndex := int(char)
			charX := (charIndex % 16) * charWidth
			charY := (charIndex / 16) * charHeight
			charImg := Image8Crop(fontImg, image.Rect(
				charX,
				charY,
				charX+charWidth,
				charY+charHeight,
			))
			Image8Draw(img, charImg, image.Point{
				X: c * charWidth,
				Y: l * charHeight,
			})
		}
	}

	return img
}

func (t TextImage) Image24() Image24 {
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

	img := NewImage24(image.Rect(0, 0, width*charWidth, height*charHeight))
	fontImg := fontImage.Image24()

	for l, line := range lines {
		for c, char := range line {
			charIndex := int(char)
			charX := (charIndex % 16) * charWidth
			charY := (charIndex / 16) * charHeight
			charImg := Image24Crop(fontImg, image.Rect(
				charX,
				charY,
				charX+charWidth,
				charY+charHeight,
			))
			Image24Draw(img, charImg, image.Point{
				X: c * charWidth,
				Y: l * charHeight,
			})
		}
	}

	return img
}
