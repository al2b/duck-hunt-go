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

type AnimationInterface interface {
	Duration() time.Duration
	At(time.Duration) *Image
}

type AnimationFrame struct {
	Image    *Image
	Duration time.Duration
}

type Animation []AnimationFrame

func (a Animation) Duration() (duration time.Duration) {
	for _, frame := range a {
		duration += frame.Duration
	}
	return
}

func (a Animation) At(at time.Duration) *Image {
	var duration time.Duration
	for _, frame := range a {
		duration += frame.Duration
		if duration < at {
			continue
		}
		return frame.Image
	}
	return nil
}

/**********/
/* Loader */
/**********/

type AnimationLoader struct {
	FS   fs.ReadFileFS
	Path string
}

func (loader AnimationLoader) Load() (animation Animation, err error) {
	var (
		file fs.File
	)

	if file, err = loader.FS.Open(loader.Path); err != nil {
		return nil, err
	}

	switch filepath.Ext(loader.Path) {
	case ".png", ".apng":
		var src apng.APNG
		if src, err = apng.DecodeAll(file); err != nil {
			return nil, err
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
			animation = append(animation, AnimationFrame{
				Image:    &Image{NRGBA: nrgba},
				Duration: (time.Duration(srcFrame.DelayNumerator) * time.Second) / time.Duration(delayDenominator),
			})
		}
		return
	case ".gif":
		var src *gif.GIF
		if src, err = gif.DecodeAll(file); err != nil {
			return nil, err
		}
		for i, srcFrameImage := range src.Image {
			bounds := srcFrameImage.Bounds()
			nrgba := image.NewNRGBA(bounds)
			draw.Draw(nrgba, bounds, srcFrameImage, bounds.Min, draw.Src)
			animation = append(animation, AnimationFrame{
				Image:    &Image{NRGBA: nrgba},
				Duration: time.Duration(src.Delay[i]) * 10 * time.Millisecond,
			})
		}
		return
	default:
		return nil, fmt.Errorf("unsupported animation file type")
	}
}

/**********/
/* Repeat */
/**********/

type RepeatAnimation struct {
	Animation AnimationInterface
	Count     int
}

func (a RepeatAnimation) Duration() time.Duration {
	return a.Animation.Duration() * time.Duration(a.Count)
}

func (a RepeatAnimation) At(at time.Duration) *Image {
	return a.Animation.At(
		at % a.Animation.Duration(),
	)
}

/************/
/* Sequence */
/************/

type SequenceAnimation []AnimationInterface

func (s SequenceAnimation) Duration() (duration time.Duration) {
	for _, animation := range s {
		duration += animation.Duration()
	}
	return
}

func (s SequenceAnimation) At(at time.Duration) *Image {
	var duration time.Duration
	for _, animation := range s {
		animationDuration := animation.Duration()
		duration += animationDuration
		if duration < at {
			continue
		}
		return animation.At(animationDuration - (duration - at))
	}
	return nil
}

/**********/
/* Player */
/**********/

type AnimationPlayer struct {
	Animation AnimationInterface
	OnEnd     PlayerOnEnd
	Player
}

func (p *AnimationPlayer) Step(delta time.Duration) {
	p.Player.Step(delta, p.Animation.Duration(), p.OnEnd)
}

func (p *AnimationPlayer) Image() *Image {
	if p.Player.Stopped() {
		return nil
	}

	return p.Animation.At(p.Player.Time())
}
