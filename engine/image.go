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

func (img *Image) Draw(drawers ...Drawer) *Image {
	for _, drawer := range drawers {
		drawer.Draw(img)
	}
	return img
}

func (img *Image) Resize(size Size) *Image {
	if img.Bounds().Dx() == size.Width && img.Bounds().Dy() == size.Height {
		return img
	}

	dst := NewImage(size)

	draw.NearestNeighbor.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

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
		dst,
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

/**********/
/* Loader */
/**********/

type ImageLoader interface {
	Load() (*Image, error)
}

func LoadImage(loader ImageLoader) (*Image, error) {
	return loader.Load()
}

func MustLoadImage(loader ImageLoader) (image *Image) {
	var err error
	if image, err = LoadImage(loader); err != nil {
		panic(err)
	}
	return
}

/********/
/* File */
/********/

func ImageFile(fs fs.ReadFileFS, path string) ImageFileHandler {
	return ImageFileHandler{
		fs:   fs,
		path: path,
	}
}

type ImageFileHandler struct {
	fs   fs.ReadFileFS
	path string
}

func (handler ImageFileHandler) Load() (*Image, error) {
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

type Imager interface {
	Image() *Image
}

func NewStaticImage(image *Image) StaticImage {
	if image == nil {
		return StaticImage{}
	}
	return StaticImage(*image)
}

type StaticImage Image

func (img *StaticImage) Image() *Image {
	return (*Image)(img)
}

func (img *StaticImage) SetImage(image *Image) {
	*img = StaticImage(*image)
}
