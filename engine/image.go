package engine

import (
	"bytes"
	"fmt"
	"github.com/charmbracelet/x/ansi"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/fs"
	"math"
)

type Image interface {
	Image8() Image8
	Image24() Image24
}

type Image8 image.PalettedImage
type Image24 image.Image

func NewImage8(rect image.Rectangle) *image.Paletted {
	return image.NewPaletted(rect, nil)
}

func NewImage24(rect image.Rectangle) *image.NRGBA {
	return image.NewNRGBA(rect)
}

type BufferedImage struct {
	buffer8  Image8
	buffer24 Image24
}

func (i BufferedImage) Image8() Image8 {
	return i.buffer8
}

func (i BufferedImage) Image24() Image24 {
	return i.buffer24
}

func LoadImageFiles(fs fs.ReadFileFS, path8, path24 string) (*BufferedImage, error) {
	var (
		data []byte
		img  image.Image
		err  error
	)

	// Image 8
	if data, err = fs.ReadFile(path8); err != nil {
		return nil, err
	}
	if img, err = png.Decode(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	img8, ok := img.(image.PalettedImage)
	if !ok {
		return nil, fmt.Errorf("not an indexed image: %s", path8)
	}

	// Image 24
	if data, err = fs.ReadFile(path24); err != nil {
		return nil, err
	}
	if img, err = png.Decode(bytes.NewReader(data)); err != nil {
		return nil, err
	}
	img24, ok := img.(*image.NRGBA)
	if !ok {
		return nil, fmt.Errorf("not a rgba image: %s", path24)
	}

	return &BufferedImage{
		buffer8:  img8,
		buffer24: img24,
	}, nil
}

func NewUniformImage(color8 Color8, color24 Color24) *UniformImage {
	return &UniformImage{
		color8:  color8,
		color24: color24,
	}
}

type UniformImage struct {
	color8  Color8
	color24 Color24
}

func (i UniformImage) Image8() Image8 {
	return NewImage8Uniform(i.color8)
}

func (i UniformImage) Image24() Image24 {
	return NewImage24Uniform(i.color24)
}

func NewImage8Uniform(color8 Color8) *Image8Uniform {
	return &Image8Uniform{
		color: color8,
	}
}

type Image8Uniform struct {
	color Color8
}

func (i *Image8Uniform) ColorIndexAt(_, _ int) uint8 {
	return uint8(i.color)
}

func (i *Image8Uniform) ColorModel() color.Model {
	return i
}

func (i *Image8Uniform) Convert(color.Color) color.Color {
	return ansi.ExtendedColor(i.color)
}

func (i *Image8Uniform) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: math.MinInt, Y: math.MinInt},
		Max: image.Point{X: math.MaxInt, Y: math.MaxInt},
	}
}

func (i *Image8Uniform) At(_, _ int) color.Color {
	return ansi.ExtendedColor(i.color)
}

func NewImage24Uniform(color24 Color24) *image.Uniform {
	return image.NewUniform(color.NRGBA(color24))
}

func Image8Crop(src Image8, rect image.Rectangle) *image.Paletted {
	r := rect.Intersect(src.Bounds()).Sub(src.Bounds().Min)
	dst := NewImage8(image.Rect(0, 0, r.Dx(), r.Dy()))

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

func Image24Crop(src Image24, rect image.Rectangle) *image.NRGBA {
	return imaging.Crop(src, rect)
}

func Image8FlipH(src Image8) *image.Paletted {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	dst := NewImage8(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.SetColorIndex(width-x-1, y,
				src.ColorIndexAt(x, y),
			)
		}
	}

	return dst
}

func Image24FlipH(src Image24) *image.NRGBA {
	return imaging.FlipH(src)
}

func Image8FlipV(src Image8) *image.Paletted {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()

	dst := NewImage8(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dst.SetColorIndex(x, height-y-1,
				src.ColorIndexAt(x, y),
			)
		}
	}

	return dst
}

func Image24FlipV(src Image24) *image.NRGBA {
	return imaging.FlipV(src)
}

func Image8Resize(img Image8, width, height int) *image.Paletted {
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

func Image24Resize(src Image24, width, height int) *image.NRGBA {
	return imaging.Resize(src, width, height, imaging.NearestNeighbor)
}

func ImageResize(src *image.NRGBA, width, height int) *image.NRGBA {
	if src.Bounds().Dx() == width || src.Bounds().Dy() == height {
		return src
	}
	return imaging.Resize(src, width, height, imaging.NearestNeighbor)
}

func Image8Draw(dst *image.Paletted, src image.PalettedImage, point image.Point) {
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

func Image24Draw(dst draw.Image, src image.Image, point image.Point) {
	draw.Draw(dst, dst.Bounds(), src, image.Point{X: -point.X, Y: -point.Y}, draw.Over)
}

type imageSet func(x, y int)

func image8Set(img *image.Paletted, c Color8) imageSet {
	return imageSet(func(x, y int) {
		img.SetColorIndex(x, y, uint8(c))
	})
}

func image24Set(img *image.NRGBA, c Color24) imageSet {
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
