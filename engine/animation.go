package engine

import (
	"image"
)

func NewAnimation(image Image, frames []AnimationFrame) *Animation {
	return &Animation{
		image:    image,
		frames:   frames,
		frame:    0,
		velocity: 6,
	}
}

type Animation struct {
	image    Image
	frames   AnimationFrames
	frame    int
	velocity int
}

func (a *Animation) Update() {
	a.frame = (a.frame + 1) % (len(a.frames) * a.velocity)
}

func (a *Animation) Image() Image {
	return &AnimationFrameImage{
		animation: a,
		frame:     a.frames[a.frame/a.velocity],
	}
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

func (s *AnimationFrameImage) Image8() Image8 {
	img := Image8Crop(s.animation.image.Image8(), image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+(s.frame.Width-1),
		s.frame.Y+(s.frame.Height-1),
	))

	if s.frame.FlipH {
		img = Image8FlipH(img)
	}

	if s.frame.FlipV {
		img = Image8FlipV(img)
	}

	return img
}

func (s *AnimationFrameImage) Image24() Image24 {
	img := Image24Crop(s.animation.image.Image24(), image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+(s.frame.Width-1),
		s.frame.Y+(s.frame.Height-1),
	))

	if s.frame.FlipH {
		img = Image24FlipH(img)
	}

	if s.frame.FlipV {
		img = Image24FlipV(img)
	}

	return img
}
