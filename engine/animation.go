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

func (animation Animation) Duration() (duration time.Duration) {
	for _, frame := range animation {
		duration += frame.Duration
	}
	return
}

func (animation Animation) At(at time.Duration) *Image {
	var duration time.Duration
	for _, frame := range animation {
		duration += frame.Duration
		if duration < at {
			continue
		}
		return frame.Image
	}
	return nil
}

/********/
/* Load */
/********/

func LoadAnimation(fS fs.ReadFileFS, path string) (animation Animation, err error) {
	var (
		file fs.File
	)

	if file, err = fS.Open(path); err != nil {
		return nil, err
	}

	switch filepath.Ext(path) {
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

/************/
/* Sequence */
/************/

type AnimationSequence []AnimationInterface

func (sequence AnimationSequence) Duration() (duration time.Duration) {
	for _, animation := range sequence {
		duration += animation.Duration()
	}
	return
}

func (sequence AnimationSequence) At(at time.Duration) *Image {
	var duration time.Duration
	for _, animation := range sequence {
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
	time      time.Duration
}

func (player *AnimationPlayer) Step(delta time.Duration) {
	player.time += delta
	duration := player.Animation.Duration()
	if player.time >= duration {
		player.time -= duration
	}
}

func (player AnimationPlayer) Image() *Image {
	return player.Animation.At(player.time)
}
