package engine

import (
	"image"
)

func NewAnimation(image *Image, frames []AnimationFrame) *Animation {
	return &Animation{
		image:    image,
		frames:   frames,
		frame:    0,
		velocity: 6,
	}
}

type Animation struct {
	image    *Image
	frames   AnimationFrames
	frame    int
	velocity int
}

func (a *Animation) Update() {
	a.frame = (a.frame + 1) % (len(a.frames) * a.velocity)
}

func (a *Animation) Image() *Image {
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

func (s *AnimationFrameImage) Image() *Image {
	img := s.animation.image.Crop(image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+s.frame.Width,
		s.frame.Y+s.frame.Height,
	))

	if s.frame.FlipH {
		img = img.FlipHorizontal()
	}

	if s.frame.FlipV {
		img = img.FlipVertical()
	}

	return img
}
