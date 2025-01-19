package engine

import (
	"bytes"
	"github.com/charmbracelet/x/ansi"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
)

func LoadImage8(data []byte) (image.PalettedImage, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	img8, ok := img.(image.PalettedImage)
	if !ok {
		panic("not an indexed image")
	}

	return img8, nil
}

func LoadImage24(data []byte) (image.Image, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	img24, ok := img.(*image.NRGBA)
	if !ok {
		panic("not a rgba image")
	}

	return img24, nil
}

type imageSet func(x, y int)

func imageSet8(img *image.Paletted, c Color8) imageSet {
	return imageSet(func(x, y int) {
		img.SetColorIndex(x, y, uint8(c))
	})
}

func imageSet24(img *image.NRGBA, c Color24) imageSet {
	return imageSet(func(x, y int) {
		img.Set(x, y, color.NRGBA(c))
	})
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func imageLine(x1, y1, x2, y2 int, set imageSet) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}
	err := dx - dy
	for {
		set(x1, y1)
		if x1 == x2 && y1 == y2 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func ImageLine8(img *image.Paletted, x1, y1, x2, y2 int, c Color8) {
	imageLine(x1, y1, x2, y2, imageSet8(img, c))
}

func ImageLine24(img *image.NRGBA, x1, y1, x2, y2 int, c Color24) {
	imageLine(x1, y1, x2, y2, imageSet24(img, c))
}

func imageRectangle(width, height int, set imageSet) {
	imageLine(0, 0, width-2, 0, set)
	imageLine(width-1, 0, width-1, height-2, set)
	imageLine(width-2, height-2, 0, height-2, set)
	imageLine(0, height-2, 0, 1, set)
}

func ImageRectangle8(width, height int, c Color8) image.PalettedImage {
	img := image.NewPaletted(image.Rect(0, 0, width, height), nil)
	imageRectangle(width, height, imageSet8(img, c))
	return img
}

func ImageRectangle24(width, height int, c Color24) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	imageRectangle(width, height, imageSet24(img, c))
	return img
}

func NewImageUniform8(c Color8) *ImageUniform8 {
	return &ImageUniform8{
		color: c,
	}
}

// ImageUniform8 is an infinite-sized [image.PalettedImage] of uniform color.
// See: [image.Uniform]
type ImageUniform8 struct {
	color Color8
}

func (i *ImageUniform8) ColorIndexAt(_, _ int) uint8 {
	return uint8(i.color)
}

func (i *ImageUniform8) ColorModel() color.Model {
	return i
}

func (i *ImageUniform8) Convert(color.Color) color.Color {
	return ansi.ExtendedColor(i.color)
}

func (i *ImageUniform8) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: math.MinInt, Y: math.MinInt},
		Max: image.Point{X: math.MaxInt, Y: math.MaxInt},
	}
}

func (i *ImageUniform8) At(_, _ int) color.Color {
	return ansi.ExtendedColor(i.color)
}

func NewImageUniform24(c Color24) *image.Uniform {
	return image.NewUniform(color.NRGBA(c))
}

func ImageDraw8(dst *image.Paletted, src image.PalettedImage, point image.Point) {
	rec := dst.Bounds().Intersect(src.Bounds())

	for y := 0; y < rec.Dy(); y++ {
		for x := 0; x < rec.Dx(); x++ {
			index := src.ColorIndexAt(x, y)
			if index != 0 {
				dst.SetColorIndex(x+point.X, y+point.Y, index)
			}
		}
	}
}

func ImageDraw24(dst draw.Image, src image.Image, point image.Point) {
	draw.Draw(dst, dst.Bounds(), src, image.Point{X: -point.X, Y: -point.Y}, draw.Over)
}

func ImageCrop8(src image.PalettedImage, rect image.Rectangle) *image.Paletted {
	r := rect.Intersect(src.Bounds()).Sub(src.Bounds().Min)
	dst := image.NewPaletted(image.Rect(0, 0, r.Dx(), r.Dy()), nil)

	for y := 0; y < r.Dy(); y++ {
		for x := 0; x < r.Dx(); x++ {
			index := src.ColorIndexAt(x+r.Min.X, y+r.Min.Y)
			if index != 0 {
				dst.SetColorIndex(x, y, index)
			}
		}
	}

	return dst
}

func ImageCrop24(src image.Image, rect image.Rectangle) *image.NRGBA {
	return imaging.Crop(src, rect)
}

func ImageFlipH8(img image.PalettedImage) *image.Paletted {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	dst := image.NewPaletted(image.Rect(0, 0, width, height), nil)

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			index := img.ColorIndexAt(width-x, y)
			if index != 0 {
				dst.SetColorIndex(x, y, index)
			}
		}
	}

	return dst
}

func ImageFlipH24(img image.Image) *image.NRGBA {
	return imaging.FlipH(img)
}

func ImageFlipV8(img image.PalettedImage) *image.Paletted {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	dst := image.NewPaletted(image.Rect(0, 0, width, height), nil)

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			index := img.ColorIndexAt(x, height-y)
			if index != 0 {
				dst.SetColorIndex(x, y, index)
			}
		}
	}

	return dst
}

func ImageFlipV24(img image.Image) *image.NRGBA {
	return imaging.FlipV(img)
}

func ImageResize8(img image.PalettedImage, width, height int) *image.Paletted {
	dst := image.NewPaletted(image.Rect(0, 0, width, height), nil)

	dx := float64(img.Bounds().Dx()) / float64(width)
	dy := float64(img.Bounds().Dy()) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcX := int(float64(x) * dx)
			srcY := int(float64(y) * dy)

			if srcX >= img.Bounds().Dx() {
				srcX = img.Bounds().Dx() - 1
			}

			if srcY >= img.Bounds().Dy() {
				srcY = img.Bounds().Dy() - 1
			}

			dst.SetColorIndex(
				x,
				y,
				img.ColorIndexAt(srcX, srcY),
			)
		}
	}

	return dst
}

func ImageResize24(img image.Image, width, height int) *image.NRGBA {
	return imaging.Resize(img, width, height, imaging.NearestNeighbor)
}
