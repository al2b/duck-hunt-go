package engine

import (
	"image"
	"image/color"
	"image/draw"
	"io/fs"
	"strings"
)

var (
	Font5x5 = Must(LoadFont(assets, "assets/font.5x5.png"))
	Font6x6 = Must(LoadFont(assets, "assets/font.6x6.png"))
)

type Text struct {
	Content string
	Font    FontInterface
	Color   color.Color
}

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

type TextDrawer struct {
	Pointer
	Text
}

func (d TextDrawer) Draw(dst *Image) {
	fontSize := d.Text.Font.Size()
	for y, line := range strings.Split(d.Text.Content, "\n") {
		for x, char := range line {
			dstMin := dst.Bounds().Min.
				Add(image.Point(d.Pointer.Point())).
				Add(image.Pt(x*fontSize.Width, y*fontSize.Height))
			mask := d.Text.Font.CharMask(int(char))
			draw.DrawMask(
				dst.NRGBA,
				image.Rectangle{
					Min: dstMin,
					Max: dstMin.Add(image.Pt(fontSize.Width, fontSize.Height)),
				},
				&image.Uniform{C: d.Text.Color},
				image.Point{},
				mask,
				mask.Bounds().Min,
				draw.Over,
			)
		}
	}
}
