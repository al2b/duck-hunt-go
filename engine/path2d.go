package engine

import (
	"math"
	"time"
)

type Path2DInterface interface {
	Duration() time.Duration
	At(time.Duration) Vector2D
}

/**********/
/* Static */
/**********/

type StaticPath2d struct {
	Position Vector2D
	Span     time.Duration
}

func (p StaticPath2d) Duration() time.Duration {
	return p.Span
}

func (p StaticPath2d) At(_ time.Duration) Vector2D {
	return p.Position
}

/**********/
/* Linear */
/**********/

type LinearPath2D struct {
	Start, End Vector2D
	Span       time.Duration
}

func (p LinearPath2D) Duration() time.Duration {
	return p.Span
}

func (p LinearPath2D) At(at time.Duration) Vector2D {
	product := float64(at) / float64(p.Span)
	return Vector2D{
		X: p.Start.X + ((p.End.X - p.Start.X) * product),
		Y: p.Start.Y + ((p.End.Y - p.Start.Y) * product),
	}
}

/********/
/* Step */
/********/

type StepPath2D struct {
	Start Vector2D
	Delta Vector2D
	Span  time.Duration
	Count int
}

func (p StepPath2D) Duration() time.Duration {
	return p.Span * time.Duration(p.Count)
}

func (p StepPath2D) At(at time.Duration) Vector2D {
	n := min(p.Count, int(at/p.Span))
	return p.Start.Add(
		p.Delta.Scale(float64(n)),
	)
}

/***********/
/* Elastic */
/***********/

type ElasticPath2D struct {
	Start, End        Vector2D
	Span              time.Duration
	Amplitude, Period float64
}

func (p ElasticPath2D) Duration() time.Duration {
	return p.Span
}

func (p ElasticPath2D) At(at time.Duration) Vector2D {
	product := float64(at) / float64(p.Span)
	factor := math.Pow(2, -10*product)*math.Sin((product-p.Period/4)*(2*math.Pi)/p.Period)*p.Amplitude + 1
	return Vector2D{
		X: p.Start.X + ((p.End.X - p.Start.X) * factor),
		Y: p.Start.Y + ((p.End.Y - p.Start.Y) * factor),
	}
}

/*********/
/* Chain */
/*********/

type ChainPath2D []Path2DInterface

func (c ChainPath2D) Duration() (duration time.Duration) {
	for _, path := range c {
		duration += path.Duration()
	}
	return
}

func (c ChainPath2D) At(at time.Duration) Vector2D {
	var (
		duration time.Duration
		position Vector2D
	)
	for _, path := range c {
		pathDuration := path.Duration()
		duration += pathDuration
		if duration < at {
			position = position.Add(
				path.At(pathDuration),
			)
			continue
		}
		return position.Add(
			path.At(pathDuration - (duration - at)),
		)
	}
	return Vector2D{}
}

/**********/
/* Player */
/**********/

type Path2DPlayer struct {
	Path  Path2DInterface
	OnEnd PlayerOnEnd
	Player
}

func (p *Path2DPlayer) Step(delta time.Duration) {
	p.Player.Step(delta, p.Path.Duration(), p.OnEnd)
}

func (p *Path2DPlayer) Position() Vector2D {
	if p.Player.Stopped() {
		return Vector2D{}
	}

	return p.Path.At(p.time)
}
