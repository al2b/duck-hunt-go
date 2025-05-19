package engine

import "time"

type Cinematic2DInterface interface {
	Duration() time.Duration
	At(time.Duration) (Vector2D, *Image)
}

type Cinematic2DFrame struct {
	Position Vector2D
	Image    *Image
	Duration time.Duration
}

type Cinematic2D []Cinematic2DFrame

func (c Cinematic2D) Duration() (duration time.Duration) {
	for _, frame := range c {
		duration += frame.Duration
	}
	return
}

func (c Cinematic2D) At(at time.Duration) (Vector2D, *Image) {
	var duration time.Duration
	for _, frame := range c {
		duration += frame.Duration
		if duration < at {
			continue
		}
		return frame.Position, frame.Image
	}
	return Vector2D{}, nil
}

/************/
/* Sequence */
/************/

type SequenceCinematic2D []Cinematic2DInterface

func (s SequenceCinematic2D) Duration() (duration time.Duration) {
	for _, cinematic := range s {
		duration += cinematic.Duration()
	}
	return
}

func (s SequenceCinematic2D) At(at time.Duration) (Vector2D, *Image) {
	var duration time.Duration
	for _, cinematic := range s {
		cinematicDuration := cinematic.Duration()
		duration += cinematicDuration
		if duration < at {
			continue
		}
		return cinematic.At(cinematicDuration - (duration - at))
	}
	return Vector2D{}, nil
}

/**********/
/* Player */
/**********/

type Cinematic2DPlayer struct {
	Cinematic Cinematic2DInterface
	OnEnd     PlayerOnEnd
	Player
}

func (p *Cinematic2DPlayer) Step(delta time.Duration) {
	p.Player.Step(delta, p.Cinematic.Duration(), p.OnEnd)
}

func (p *Cinematic2DPlayer) Position() Vector2D {
	if p.Player.Stopped() {
		return Vector2D{}
	}

	position, _ := p.Cinematic.At(p.Player.Time())
	return position
}

func (p *Cinematic2DPlayer) Image() *Image {
	if p.Player.Stopped() {
		return nil
	}

	_, image := p.Cinematic.At(p.Player.Time())
	return image
}
