package engine

import (
	"image"
	"image/color"
	"image/draw"
	"io/fs"
	"strings"
)

var Font5x5 = Must(LoadFont(assets, "assets/font.5x5.png"))

type FontInterface interface {
	Size() Size
	CharMask(int) *image.Alpha
}

type Font struct {
	size Size
	mask *image.Alpha
}

func (font Font) Size() Size {
	return font.size
}

func (font Font) CharMask(char int) *image.Alpha {
	imgMin := image.Pt((char%16)*font.size.Width, (char/16)*font.size.Height)
	return font.mask.SubImage(image.Rectangle{
		Min: imgMin,
		Max: imgMin.Add(image.Pt(font.size.Width, font.size.Height)),
	}).(*image.Alpha)
}

/********/
/* Load */
/********/

func LoadFont(fS fs.ReadFileFS, path string) (*Font, error) {
	var (
		file fs.File
		img  image.Image
		err  error
	)

	if file, err = fS.Open(path); err != nil {
		return nil, err
	}
	defer file.Close()

	if img, _, err = image.Decode(file); err != nil {
		return nil, err
	}

	var mask *image.Alpha
	switch img := img.(type) {
	case *image.Gray:
		mask = &image.Alpha{
			Pix:    img.Pix,
			Stride: img.Stride,
			Rect:   img.Rect,
		}
	default:
		panic("Unsupported font image type")
	}

	size := img.Bounds().Size()

	return &Font{
		size: Size{
			size.X / 16,
			size.Y / 16,
		},
		mask: mask,
	}, nil
}

/**********/
/* Drawer */
/**********/

func DrawText(point image.Point, text string, font FontInterface, color color.Color) TextDrawer {
	return TextDrawer{
		point: point,
		text:  text,
		font:  font,
		color: color,
	}
}

type TextDrawer struct {
	point image.Point
	text  string
	font  FontInterface
	color color.Color
}

func (drawer TextDrawer) Draw(img *Image) {
	fontSize := drawer.font.Size()
	for y, line := range strings.Split(drawer.text, "\n") {
		for x, char := range line {
			imgMin := img.Bounds().Min.
				Add(drawer.point).
				Add(image.Pt(x*fontSize.Width, y*fontSize.Height))
			mask := drawer.font.CharMask(int(char))
			draw.DrawMask(
				img.NRGBA,
				image.Rectangle{
					Min: imgMin,
					Max: imgMin.Add(image.Pt(fontSize.Width, fontSize.Height)),
				},
				&image.Uniform{C: drawer.color},
				image.Point{},
				mask,
				mask.Bounds().Min,
				draw.Over,
			)
		}
	}
}
