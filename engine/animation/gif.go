package animation

import (
	"bytes"
	"duck-hunt-go/engine"
	"image/draw"
	"image/gif"
	"io/fs"
	"time"
)

func GifFile(fs fs.ReadFileFS, path string) GifFileLoader {
	return GifFileLoader{
		fs:   fs,
		path: path,
	}
}

type GifFileLoader struct {
	fs   fs.ReadFileFS
	path string
}

func (loader GifFileLoader) Load() (animation *Animation, err error) {
	var (
		file   []byte
		imgGif *gif.GIF
	)

	if file, err = loader.fs.ReadFile(loader.path); err != nil {
		return nil, err
	}

	if imgGif, err = gif.DecodeAll(bytes.NewReader(file)); err != nil {
		return nil, err
	}

	size := engine.Size{
		Width:  imgGif.Config.Width,
		Height: imgGif.Config.Height,
	}

	animation = &Animation{
		size: size,
	}

	for i, img := range imgGif.Image {
		imgFrame := Frame{
			Image:    engine.NewImage(size),
			Duration: time.Duration(imgGif.Delay[i]) * 10 * time.Millisecond,
		}

		bounds := img.Bounds()
		draw.Draw(imgFrame.Image.NRGBA, bounds, img, bounds.Min, draw.Src)

		animation.frames = append(animation.frames, imgFrame)
		animation.duration += imgFrame.Duration
	}

	return
}
