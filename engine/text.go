package engine

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"strings"
	"sync"
)

var (
	imageFont5x5 = Must(
		LoadImageFile(imagesFS, "images/font.5x5.png"),
	)
	imageFont8x8 = Must(
		LoadImageFile(imagesFS, "images/font.8x8.png"),
	)
)

/* *** */
/* 5x5 */
/* *** */

func NewText5x5(text string, color color.Color) Text5x5 {
	return Text5x5{
		text:  text,
		color: color,
	}
}

type Text5x5 struct {
	text  string
	color color.Color
}

func (t Text5x5) Image() image.Image {
	return textImage(t.text, imageFont5x5, 5, 5, t.color)
}

/* *** */
/* 8x8 */
/* *** */

func NewText8x8(text string, color color.Color) Text8x8 {
	return Text8x8{
		text:  text,
		color: color,
	}
}

type Text8x8 struct {
	text  string
	color color.Color
}

func (t Text8x8) Image() image.Image {
	return textImage(t.text, imageFont8x8, 8, 8, t.color)
}

func textImage(text string, imageFont image.Image, charWidth, charHeight int, c color.Color) image.Image {
	lines := strings.Split(text, "\n")

	// Compute dimensions
	var width, height int
	for _, line := range lines {
		w := len(line)
		if w > width {
			width = w
		}
	}
	width *= charWidth
	height = len(lines) * charHeight

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for l, line := range lines {
		for c, char := range line {
			charIndex := int(char)
			charX := (charIndex % 16) * charWidth
			charY := (charIndex / 16) * charHeight
			charImg := imaging.Crop(imageFont, image.Rect(
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

	// Color
	if c != nil {
		r, g, b, _ := c.RGBA()
		r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
		var wg sync.WaitGroup
		for y := 0; y < height; y++ {
			wg.Add(1)
			go func(y int) {
				defer wg.Done()
				for x := 0; x < width; x++ {
					offset := img.PixOffset(x, y)
					if img.Pix[offset] == 255 && img.Pix[offset] == img.Pix[offset+1] && img.Pix[offset] == img.Pix[offset+2] {
						img.Pix[offset] = r8
						img.Pix[offset+1] = g8
						img.Pix[offset+2] = b8
					}
				}
			}(y)
		}
		wg.Wait()
	}

	return img
}
