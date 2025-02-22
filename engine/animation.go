package engine

import (
	"fmt"
	"github.com/kettek/apng"
	"image"
	"image/draw"
	"image/gif"
	"io/fs"
	"path/filepath"
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

/********/
/* File */
/********/

func AnimationFile(fs fs.ReadFileFS, path string) AnimationFileHandler {
	return AnimationFileHandler{
		fs:   fs,
		path: path,
	}
}

type AnimationFileHandler struct {
	fs   fs.ReadFileFS
	path string
}

func (handler AnimationFileHandler) Load() (animation *Animation, err error) {
	var file fs.File

	if file, err = handler.fs.Open(handler.path); err != nil {
		return nil, err
	}

	switch filepath.Ext(handler.path) {
	case ".png", ".apng":
		var src apng.APNG
		if src, err = apng.DecodeAll(file); err != nil {
			return nil, err
		}
		animation = &Animation{
			size: Size{
				Width:  src.Frames[0].Image.Bounds().Dx(),
				Height: src.Frames[0].Image.Bounds().Dy(),
			},
		}
		for _, srcFrame := range src.Frames {
			delayDenominator := srcFrame.DelayDenominator
			if delayDenominator == 0 {
				delayDenominator = 100
			}
			var nrgba *image.NRGBA
			switch srcFrameImage := srcFrame.Image.(type) {
			case *image.NRGBA:
				nrgba = srcFrameImage
			default:
				bounds := srcFrameImage.Bounds()
				nrgba = image.NewNRGBA(bounds)
				draw.Draw(nrgba, bounds, srcFrameImage, bounds.Min, draw.Src)
			}
			imgFrame := AnimationFrame{
				Image:    &Image{NRGBA: nrgba},
				Duration: (time.Duration(srcFrame.DelayNumerator) * time.Second) / time.Duration(delayDenominator),
			}
			animation.frames = append(animation.frames, imgFrame)
			animation.duration += imgFrame.Duration
		}
		return
	case ".gif":
		var src *gif.GIF
		if src, err = gif.DecodeAll(file); err != nil {
			return nil, err
		}
		animation = &Animation{
			size: Size{
				Width:  src.Config.Width,
				Height: src.Config.Height,
			},
		}
		for i, srcFrameImage := range src.Image {
			bounds := srcFrameImage.Bounds()
			nrgba := image.NewNRGBA(bounds)
			draw.Draw(nrgba, bounds, srcFrameImage, bounds.Min, draw.Src)
			imgFrame := AnimationFrame{
				Image:    &Image{NRGBA: nrgba},
				Duration: time.Duration(src.Delay[i]) * 10 * time.Millisecond,
			}
			animation.frames = append(animation.frames, imgFrame)
			animation.duration += imgFrame.Duration
		}
		return
	default:
		return nil, fmt.Errorf("unsupported animation file type")
	}
}
