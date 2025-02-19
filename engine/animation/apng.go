package animation

import (
	"bytes"
	"duck-hunt-go/engine"
	"github.com/kettek/apng"
	"image"
	"io/fs"
	"time"
)

func ApngFile(fs fs.ReadFileFS, path string) ApngFileLoader {
	return ApngFileLoader{
		fs:   fs,
		path: path,
	}
}

type ApngFileLoader struct {
	fs   fs.ReadFileFS
	path string
}

func (loader ApngFileLoader) Load() (animation *Animation, err error) {
	var (
		file   []byte
		imgPng apng.APNG
	)

	if file, err = loader.fs.ReadFile(loader.path); err != nil {
		return nil, err
	}

	if imgPng, err = apng.DecodeAll(bytes.NewReader(file)); err != nil {
		return nil, err
	}

	animation = &Animation{
		Size: engine.Size{
			Width:  imgPng.Frames[0].Image.Bounds().Dx(),
			Height: imgPng.Frames[0].Image.Bounds().Dy(),
		},
	}

	for _, imgPngFrame := range imgPng.Frames {
		delayDenominator := imgPngFrame.DelayDenominator
		if delayDenominator == 0 {
			delayDenominator = 100 // Par défaut, APNG considère 100 si den == 0
		}

		imgFrame := Frame{
			Image:    &engine.Image{NRGBA: imgPngFrame.Image.(*image.NRGBA)},
			Duration: (time.Duration(imgPngFrame.DelayNumerator) * time.Second) / time.Duration(delayDenominator),
		}

		animation.Frames = append(animation.Frames, imgFrame)
		animation.Duration += imgFrame.Duration
	}

	return
}
