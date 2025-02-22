package engine

import (
	"bytes"
	"github.com/kettek/apng"
	"image"
	"image/draw"
	"image/gif"
	"io/fs"
	"time"
)

type Animation struct {
	size     Size
	duration time.Duration
	frames   []AnimationFrame
	time     time.Duration
	current  int
}

func (a *Animation) Size() Size {
	return a.size
}

func (a *Animation) Step(delta time.Duration) {
	a.time += delta
	if a.time >= a.duration {
		a.time -= a.duration
	}
}

func (a *Animation) Image() *Image {
	var duration time.Duration
	for i, frame := range a.frames {
		duration += frame.Duration
		if a.time <= duration {
			return a.frames[i].Image
		}
	}
	return a.frames[len(a.frames)-1].Image
}

type AnimationFrame struct {
	Image    *Image
	Duration time.Duration
}

type Animator interface {
	Animation() *Animation
}

/**********/
/* Loader */
/**********/

type AnimationLoader interface {
	Load() (*Animation, error)
}

func LoadAnimation(loader AnimationLoader) (*Animation, error) {
	return loader.Load()
}

func MustLoadAnimation(loader AnimationLoader) (animation *Animation) {
	var err error
	if animation, err = LoadAnimation(loader); err != nil {
		panic(err)
	}
	return
}

/************/
/* Gif File */
/************/

func AnimationGifFile(fs fs.ReadFileFS, path string) AnimationGifFileHandler {
	return AnimationGifFileHandler{
		fs:   fs,
		path: path,
	}
}

type AnimationGifFileHandler struct {
	fs   fs.ReadFileFS
	path string
}

func (handler AnimationGifFileHandler) Load() (animation *Animation, err error) {
	var (
		file   []byte
		imgGif *gif.GIF
	)

	if file, err = handler.fs.ReadFile(handler.path); err != nil {
		return nil, err
	}

	if imgGif, err = gif.DecodeAll(bytes.NewReader(file)); err != nil {
		return nil, err
	}

	size := Size{
		Width:  imgGif.Config.Width,
		Height: imgGif.Config.Height,
	}

	animation = &Animation{
		size: size,
	}

	for i, img := range imgGif.Image {
		imgFrame := AnimationFrame{
			Image:    NewImage(size),
			Duration: time.Duration(imgGif.Delay[i]) * 10 * time.Millisecond,
		}

		bounds := img.Bounds()
		draw.Draw(imgFrame.Image.NRGBA, bounds, img, bounds.Min, draw.Src)

		animation.frames = append(animation.frames, imgFrame)
		animation.duration += imgFrame.Duration
	}

	return
}

/************/
/* Png File */
/************/

func AnimationPngFile(fs fs.ReadFileFS, path string) AnimationPngFileHandler {
	return AnimationPngFileHandler{
		fs:   fs,
		path: path,
	}
}

type AnimationPngFileHandler struct {
	fs   fs.ReadFileFS
	path string
}

func (handler AnimationPngFileHandler) Load() (animation *Animation, err error) {
	var (
		file   []byte
		imgPng apng.APNG
	)

	if file, err = handler.fs.ReadFile(handler.path); err != nil {
		return nil, err
	}

	if imgPng, err = apng.DecodeAll(bytes.NewReader(file)); err != nil {
		return nil, err
	}

	animation = &Animation{
		size: Size{
			Width:  imgPng.Frames[0].Image.Bounds().Dx(),
			Height: imgPng.Frames[0].Image.Bounds().Dy(),
		},
	}

	for _, imgPngFrame := range imgPng.Frames {
		delayDenominator := imgPngFrame.DelayDenominator
		if delayDenominator == 0 {
			delayDenominator = 100 // Par défaut, APNG considère 100 si den == 0
		}

		imgFrame := AnimationFrame{
			Image:    &Image{NRGBA: imgPngFrame.Image.(*image.NRGBA)},
			Duration: (time.Duration(imgPngFrame.DelayNumerator) * time.Second) / time.Duration(delayDenominator),
		}

		animation.frames = append(animation.frames, imgFrame)
		animation.duration += imgFrame.Duration
	}

	return
}
