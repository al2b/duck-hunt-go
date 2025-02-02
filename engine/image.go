package engine

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/png"
	"io/fs"
)

func NewStaticImage(img image.Image) StaticImage {
	return StaticImage{image: img}
}

type StaticImage struct{ image image.Image }

func (i StaticImage) Image() image.Image { return i.image }

func LoadImageFile(fs fs.ReadFileFS, path string) (img image.Image, err error) {
	var data []byte

	if data, err = fs.ReadFile(path); err != nil {
		return nil, err
	}

	if img, err = png.Decode(bytes.NewReader(data)); err != nil {
		return nil, err
	}

	return img, nil
}

func MustLoadImageFile(fs fs.ReadFileFS, path string) image.Image {
	img, err := LoadImageFile(fs, path)
	if err != nil {
		panic(err)
	}
	return img
}

func ImageResize(src *image.NRGBA, width, height int) *image.NRGBA {
	if src.Bounds().Dx() == width || src.Bounds().Dy() == height {
		return src
	}
	return imaging.Resize(src, width, height, imaging.NearestNeighbor)
}

func ImageLine(img *image.NRGBA, x1, y1, x2, y2 int, c color.Color) {
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
		img.Set(x1, y1, c)
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

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
