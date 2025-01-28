package engine

import (
	"github.com/disintegration/imaging"
	"image"
)

func NewAnimation(image image.Image, frames []AnimationFrame) *Animation {
	return &Animation{
		image:    image,
		frames:   frames,
		frame:    0,
		velocity: 6,
	}
}

type Animation struct {
	image    image.Image
	frames   AnimationFrames
	frame    int
	velocity int
}

func (a *Animation) Update() {
	a.frame = (a.frame + 1) % (len(a.frames) * a.velocity)
}

func (a *Animation) Image() image.Image {
	return (&AnimationFrameImage{
		animation: a,
		frame:     a.frames[a.frame/a.velocity],
	}).Image()
}

type AnimationFrame struct {
	X, Y          int
	Width, Height int
	FlipH, FlipV  bool
}

type AnimationFrames []AnimationFrame

type AnimationFrameImage struct {
	animation *Animation
	frame     AnimationFrame
}

func (s *AnimationFrameImage) Image() image.Image {
	img := imaging.Crop(s.animation.image, image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+(s.frame.Width-1),
		s.frame.Y+(s.frame.Height-1),
	))

	if s.frame.FlipH {
		img = imaging.FlipH(img)
	}

	if s.frame.FlipV {
		img = imaging.FlipV(img)
	}

	return img
}
