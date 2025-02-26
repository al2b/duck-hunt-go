package engine

import (
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"io/fs"
)

func NewImage(size Size) *Image {
	return &Image{
		NRGBA: image.NewNRGBA(image.Rect(
			0,
			0,
			size.Width,
			size.Height,
		)),
	}
}

type Image struct {
	*image.NRGBA
}

func (img *Image) Size() Size {
	size := img.Bounds().Size()
	return Size{
		Width:  size.X,
		Height: size.Y,
	}
}

func (img *Image) Set(point image.Point, c color.Color) {
	imgMin := img.Bounds().Min
	img.NRGBA.Set(
		imgMin.X+point.X,
		imgMin.Y+point.Y,
		c,
	)
}

func (img *Image) Draw(drawers ...Drawer) *Image {
	for _, drawer := range drawers {
		drawer.Draw(img)
	}
	return img
}

func (img *Image) SubImage(point image.Point, size Size) *Image {
	return &Image{
		img.NRGBA.SubImage(image.Rect(
			point.X, point.Y,
			point.X+size.Width, point.Y+size.Height,
		)).(*image.NRGBA),
	}
}

func (img *Image) Resize(size Size) *Image {
	if img.Bounds().Dx() == size.Width && img.Bounds().Dy() == size.Height {
		return img
	}

	dst := NewImage(size)

	draw.NearestNeighbor.Scale(dst.NRGBA, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dst
}

func (img *Image) Crop(rectangle image.Rectangle) *Image {
	if img.Bounds() == rectangle {
		return img
	}

	dst := NewImage(Size{
		Width:  rectangle.Dx(),
		Height: rectangle.Dy(),
	})

	draw.Draw(
		dst.NRGBA,
		dst.Bounds(),
		img,
		rectangle.Min,
		draw.Over,
	)

	return dst
}

func (img *Image) FlipHorizontal() *Image {
	size := img.Size()
	dst := NewImage(size)

	for y := 0; y < size.Height; y++ {
		for x := 0; x < size.Width; x++ {
			srcIdx := y*img.Stride + (size.Width-x-1)*4
			dstIdx := y*img.Stride + x*4
			copy(dst.Pix[dstIdx:dstIdx+4], img.Pix[srcIdx:srcIdx+4])
		}
	}

	return dst
}

func (img *Image) FlipVertical() *Image {
	size := img.Size()
	dst := NewImage(size)

	for y := 0; y < size.Height; y++ {
		srcY := size.Height - y - 1
		copy(
			dst.Pix[y*img.Stride:(y+1)*img.Stride],
			img.Pix[srcY*img.Stride:(srcY+1)*img.Stride],
		)
	}

	return dst
}

func (img *Image) Fill(c color.Color) *Image {
	rgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	pixel := [4]byte{rgba.R, rgba.G, rgba.B, rgba.A}

	for i := 0; i < len(img.Pix); i += 4 {
		copy(img.Pix[i:i+4], pixel[:])
	}

	return img
}

/********/
/* Load */
/********/

func LoadImage(fS fs.ReadFileFS, path string) (*Image, error) {
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

	switch img := img.(type) {
	case *image.NRGBA:
		return &Image{NRGBA: img}, nil
	default:
		bounds := img.Bounds()
		nrgba := image.NewNRGBA(bounds)
		draw.Draw(nrgba, bounds, img, bounds.Min, draw.Src)
		return &Image{NRGBA: nrgba}, nil
	}
}

/**********/
/* Drawer */
/**********/

type Drawer interface {
	Draw(*Image)
}

func DrawImage(point image.Point, img *Image) ImageDrawer {
	return ImageDrawer{
		point: point,
		image: img,
	}
}

func DrawCenteredImage(point image.Point, img *Image) ImageDrawer {
	size := img.Size()
	return DrawImage(
		point.Sub(image.Pt(
			(size.Width-1)/2,
			(size.Height-1)/2,
		)),
		img,
	)
}

type ImageDrawer struct {
	point image.Point
	image *Image
}

func (drawer ImageDrawer) Draw(img *Image) {
	bounds := drawer.image.Bounds()
	imgMin := img.Bounds().Min.Add(
		image.Pt(drawer.point.X, drawer.point.Y),
	)
	draw.Draw(
		img.NRGBA,
		image.Rectangle{
			Min: imgMin,
			Max: imgMin.Add(drawer.image.Bounds().Size()),
		},
		drawer.image.NRGBA,
		bounds.Min,
		draw.Over,
	)
}
