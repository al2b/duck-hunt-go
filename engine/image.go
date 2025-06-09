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

func (img *Image) Image() *Image {
	return img
}

func (img *Image) Size() Size {
	size := img.Bounds().Size()
	return Size{size.X, size.Y}
}

func (img *Image) Set(point Point, c color.Color) {
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

func (img *Image) SubImage(point Point, size Size) *Image {
	imgMin := img.Bounds().Min
	return &Image{
		NRGBA: img.NRGBA.SubImage(image.Rect(
			imgMin.X+point.X, imgMin.Y+point.Y,
			imgMin.X+point.X+size.Width, imgMin.Y+point.Y+size.Height,
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

	dst := NewImage(Size{rectangle.Dx(), rectangle.Dy()})

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
	dst := &image.NRGBA{
		Pix:    make([]uint8, len(img.NRGBA.Pix)),
		Stride: img.NRGBA.Stride,
		Rect:   img.NRGBA.Bounds(),
	}

	size := img.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			srcIdx := y*img.Stride + (size.X-x-1)*4
			dstIdx := y*img.Stride + x*4
			copy(dst.Pix[dstIdx:dstIdx+4], img.Pix[srcIdx:srcIdx+4])
		}
	}

	return &Image{
		NRGBA: dst,
	}
}

func (img *Image) FlipVertical() *Image {
	dst := &image.NRGBA{
		Pix:    make([]uint8, len(img.NRGBA.Pix)),
		Stride: img.NRGBA.Stride,
		Rect:   img.NRGBA.Bounds(),
	}

	size := img.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		srcY := size.Y - y - 1
		copy(
			dst.Pix[y*img.Stride:(y+1)*img.Stride],
			img.Pix[srcY*img.Stride:(srcY+1)*img.Stride],
		)
	}

	return &Image{
		NRGBA: dst,
	}
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

type ImageLoader struct {
	FS   fs.ReadFileFS
	Path string
}

func (loader ImageLoader) Load() (*Image, error) {
	var (
		file fs.File
		img  image.Image
		err  error
	)

	if file, err = loader.FS.Open(loader.Path); err != nil {
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
/* Imager */
/**********/

type Imager interface {
	Image() *Image
}

/**********/
/* Drawer */
/**********/

type Drawer interface {
	Draw(*Image)
}

type ImageDrawer struct {
	Pointer
	Imager
}

func (d ImageDrawer) Draw(dst *Image) {
	if d.Imager == nil {
		return
	}

	src := d.Imager.Image()
	if src == nil {
		return
	}

	bounds := src.Bounds()
	dstMin := dst.Bounds().Min.
		Add(image.Point(d.Pointer.Point()))
	draw.Draw(
		dst.NRGBA,
		image.Rectangle{
			Min: dstMin,
			Max: dstMin.Add(bounds.Size()),
		},
		src.NRGBA,
		bounds.Min,
		draw.Over,
	)
}

/**********/
/* Slicer */
/**********/

type ImageSlicer struct {
	Imager
	Pointer
	Sizer
}

func (slicer ImageSlicer) Image() *Image {
	img := slicer.Imager.Image()
	point := slicer.Pointer.Point()
	size := slicer.Sizer.Size()

	imgMin := img.Bounds().Min.Add(image.Pt(point.X, point.Y))
	return &Image{
		NRGBA: img.NRGBA.SubImage(image.Rect(
			imgMin.X, imgMin.Y,
			imgMin.X+size.Width, imgMin.Y+size.Height,
		)).(*image.NRGBA),
	}
}
