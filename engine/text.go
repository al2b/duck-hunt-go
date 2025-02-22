package engine

import (
	"image"
	"image/color"
	"strings"
	"sync"
)

var (
	imageFont5x5 = MustLoadImage(ImageFile(assets, "assets/font.5x5.png"))
	imageFont8x8 = MustLoadImage(ImageFile(assets, "assets/font.8x8.png"))
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

func (t Text5x5) Image() *Image {
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

func (t Text8x8) Image() *Image {
	return textImage(t.text, imageFont8x8, 8, 8, t.color)
}

func textImage(text string, imageFont *Image, charWidth, charHeight int, c color.Color) *Image {
	lines := strings.Split(text, "\n")

	// Compute size
	var size Size
	for _, line := range lines {
		width := len(line)
		if width > size.Width {
			size.Width = width
		}
	}
	size.Width *= charWidth
	size.Height = len(lines) * charHeight

	img := NewImage(size)

	for l, line := range lines {
		for c, char := range line {
			charIndex := int(char)
			charX := (charIndex % 16) * charWidth
			charY := (charIndex / 16) * charHeight
			charImg := imageFont.Crop(image.Rect(
				charX,
				charY,
				charX+charWidth,
				charY+charHeight,
			))
			img.Draw(
				DrawImage(
					image.Pt(
						c*charWidth,
						l*charHeight,
					),
					charImg),
			)
		}
	}

	// Color
	if c != nil {
		r, g, b, _ := c.RGBA()
		r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
		var wg sync.WaitGroup
		for y := 0; y < size.Height; y++ {
			wg.Add(1)
			go func(y int) {
				defer wg.Done()
				for x := 0; x < size.Width; x++ {
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
