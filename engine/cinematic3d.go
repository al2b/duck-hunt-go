package engine

import "time"

type Cinematic3DInterface interface {
	Duration() time.Duration
	At(time.Duration) (Vector3D, *Image)
}

type Cinematic3DFrame struct {
	Position Vector3D
	Image    *Image
	Duration time.Duration
}

type Cinematic3D []Cinematic3DFrame

func (c Cinematic3D) Duration() (duration time.Duration) {
	for _, frame := range c {
		duration += frame.Duration
	}
	return
}

func (c Cinematic3D) At(at time.Duration) (Vector3D, *Image) {
	var duration time.Duration
	for _, frame := range c {
		duration += frame.Duration
		if duration < at {
			continue
		}
		return frame.Position, frame.Image
	}
	return Vector3D{}, nil
}

/**********/
/* Player */
/**********/

type Cinematic3DPlayer struct {
	Cinematic Cinematic3DInterface
	OnEnd     PlayerOnEnd
	Player
}

func (p *Cinematic3DPlayer) Step(delta time.Duration) {
	p.Player.Step(delta, p.Cinematic.Duration(), p.OnEnd)
}

func (p *Cinematic3DPlayer) Position() Vector3D {
	if p.Player.Stopped() {
		return Vector3D{}
	}

	position, _ := p.Cinematic.At(p.Player.Time())
	return position
}

func (p *Cinematic3DPlayer) Image() *Image {
	if p.Player.Stopped() {
		return nil
	}

	_, image := p.Cinematic.At(p.Player.Time())
	return image
}
