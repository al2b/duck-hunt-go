package engine

import (
	"image"
)

func NewAnimation(image8 image.PalettedImage, image24 image.Image, frames []AnimationFrame) *Animation {
	return &Animation{
		image8:   image8,
		image24:  image24,
		frames:   frames,
		frame:    0,
		velocity: 6,
	}
}

type Animation struct {
	image8   image.PalettedImage
	image24  image.Image
	frames   AnimationFrames
	frame    int
	velocity int
}

func (a *Animation) Update() {
	a.frame = (a.frame + 1) % (len(a.frames) * a.velocity)
}

func (a *Animation) Sprite(coordinates Point) Sprite {
	return &AnimationFrameSprite{
		animation: a,
		frame:     a.frames[a.frame/a.velocity],
		Point:     coordinates,
	}
}

type AnimationFrame struct {
	X, Y          int
	Width, Height int
	FlipH, FlipV  bool
}

type AnimationFrames []AnimationFrame

type AnimationFrameSprite struct {
	animation *Animation
	frame     AnimationFrame
	Point
}

func (s *AnimationFrameSprite) Image8() image.PalettedImage {
	img := ImageCrop8(s.animation.image8, image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+(s.frame.Width-1),
		s.frame.Y+(s.frame.Height-1),
	))

	if s.frame.FlipH {
		img = ImageFlipH8(img)
	}

	if s.frame.FlipV {
		img = ImageFlipV8(img)
	}

	return img
}

func (s *AnimationFrameSprite) Image24() image.Image {
	img := ImageCrop24(s.animation.image24, image.Rect(
		s.frame.X,
		s.frame.Y,
		s.frame.X+(s.frame.Width-1),
		s.frame.Y+(s.frame.Height-1),
	))

	if s.frame.FlipH {
		img = ImageFlipH24(img)
	}

	if s.frame.FlipV {
		img = ImageFlipV24(img)
	}

	return img
}
