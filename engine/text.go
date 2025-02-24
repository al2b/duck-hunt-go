package engine

import (
	"image"
	"image/color"
	"image/draw"
	"io/fs"
	"strings"
)

var Font5x5 = MustLoadFont(FontImageFile(assets, "assets/font.5x5.png"))

type Font struct {
	size Size
	mask *image.Alpha
}

func (font Font) Size() Size {
	return font.size
}

func (font Font) DrawChar(img *Image, point image.Point, char int, c color.Color) {
	imgMin := img.Bounds().Min.Add(point)
	draw.DrawMask(
		img,
		image.Rectangle{
			Min: imgMin,
			Max: imgMin.Add(image.Pt(font.size.Width, font.size.Height)),
		},
		&image.Uniform{c},
		image.Point{},
		font.mask,
		image.Pt(
			(char%16)*font.size.Width,
			(char/16)*font.size.Height,
		),
		draw.Over,
	)
}

/**********/
/* Loader */
/**********/

type FontLoader interface {
	Load() (*Font, error)
}

func LoadFont(loader FontLoader) (*Font, error) {
	return loader.Load()
}

func MustLoadFont(loader FontLoader) (font *Font) {
	var err error
	if font, err = LoadFont(loader); err != nil {
		panic(err)
	}
	return
}

/**************/
/* Image File */
/**************/

func FontImageFile(fs fs.ReadFileFS, path string) FontImageFileHandler {
	return FontImageFileHandler{
		fs:   fs,
		path: path,
	}
}

type FontImageFileHandler struct {
	fs   fs.ReadFileFS
	path string
}

func (handler FontImageFileHandler) Load() (*Font, error) {
	var (
		file fs.File
		img  image.Image
		err  error
	)

	if file, err = handler.fs.Open(handler.path); err != nil {
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

func DrawText(point image.Point, text string, font *Font, color color.Color) TextDrawer {
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
	font  *Font
	color color.Color
}

func (drawer TextDrawer) Draw(img *Image) {
	size := drawer.font.Size()
	for y, line := range strings.Split(drawer.text, "\n") {
		for x, char := range line {
			drawer.font.DrawChar(
				img,
				drawer.point.Add(
					image.Pt(x*size.Width, y*size.Height),
				),
				int(char),
				drawer.color,
			)
		}
	}
}
