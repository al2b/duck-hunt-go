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

func (path StaticPath2d) Duration() time.Duration {
	return path.Span
}

func (path StaticPath2d) At(_ time.Duration) Vector2D {
	return path.Position
}

/**********/
/* Linear */
/**********/

type LinearPath2D struct {
	Start, End Vector2D
	Span       time.Duration
}

func (path LinearPath2D) Duration() time.Duration {
	return path.Span
}

func (path LinearPath2D) At(at time.Duration) Vector2D {
	product := float64(at) / float64(path.Span)
	return Vector2D{
		X: path.Start.X + ((path.End.X - path.Start.X) * product),
		Y: path.Start.Y + ((path.End.Y - path.Start.Y) * product),
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

func (path StepPath2D) Duration() time.Duration {
	return path.Span * time.Duration(path.Count)
}

func (path StepPath2D) At(at time.Duration) Vector2D {
	n := min(path.Count, int(at/path.Span))
	return path.Start.Add(
		path.Delta.Scale(float64(n)),
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

func (path ElasticPath2D) Duration() time.Duration {
	return path.Span
}

func (path ElasticPath2D) At(at time.Duration) Vector2D {
	product := float64(at) / float64(path.Span)
	factor := math.Pow(2, -10*product)*math.Sin((product-path.Period/4)*(2*math.Pi)/path.Period)*path.Amplitude + 1
	return Vector2D{
		X: path.Start.X + ((path.End.X - path.Start.X) * factor),
		Y: path.Start.Y + ((path.End.Y - path.Start.Y) * factor),
	}
}

/*********/
/* Chain */
/*********/

type ChainPath2D []Path2DInterface

func (chain ChainPath2D) Duration() (duration time.Duration) {
	for _, path := range chain {
		duration += path.Duration()
	}
	return
}

func (chain ChainPath2D) At(at time.Duration) Vector2D {
	var (
		duration time.Duration
		position Vector2D
	)
	for _, path := range chain {
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
