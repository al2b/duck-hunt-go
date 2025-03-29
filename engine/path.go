package engine

import (
	"math"
	"time"
)

type PathInterface interface {
	Duration() time.Duration
	At(time.Duration) Vector
}

/*********/
/* Fixed */
/*********/

type FixedPath struct {
	Position Vector
	Span     time.Duration
}

func (path FixedPath) Duration() time.Duration {
	return path.Span
}

func (path FixedPath) At(_ time.Duration) Vector {
	return path.Position
}

/**********/
/* Linear */
/**********/

type LinearPath struct {
	Start, End Vector
	Span       time.Duration
}

func (path LinearPath) Duration() time.Duration {
	return path.Span
}

func (path LinearPath) At(at time.Duration) Vector {
	product := float64(at) / float64(path.Span)
	return Vector{
		X: path.Start.X + ((path.End.X - path.Start.X) * product),
		Y: path.Start.Y + ((path.End.Y - path.Start.Y) * product),
	}
}

/********/
/* Step */
/********/

type StepPath struct {
	Start Vector
	Delta Vector
	Span  time.Duration
	Count int
}

func (path StepPath) Duration() time.Duration {
	return path.Span * time.Duration(path.Count)
}

func (path StepPath) At(at time.Duration) Vector {
	n := min(path.Count, int(at/path.Span))
	return path.Start.Add(
		path.Delta.Scale(float64(n)),
	)
}

/***********/
/* Elastic */
/***********/

type ElasticPath struct {
	Start, End        Vector
	Span              time.Duration
	Amplitude, Period float64
}

func (path ElasticPath) Duration() time.Duration {
	return path.Span
}

func (path ElasticPath) At(at time.Duration) Vector {
	product := float64(at) / float64(path.Span)
	factor := math.Pow(2, -10*product)*math.Sin((product-path.Period/4)*(2*math.Pi)/path.Period)*path.Amplitude + 1
	return Vector{
		X: path.Start.X + ((path.End.X - path.Start.X) * factor),
		Y: path.Start.Y + ((path.End.Y - path.Start.Y) * factor),
	}
}

/*********/
/* Chain */
/*********/

type ChainPath []PathInterface

func (chain ChainPath) Duration() (duration time.Duration) {
	for _, path := range chain {
		duration += path.Duration()
	}
	return
}

func (chain ChainPath) At(at time.Duration) Vector {
	var (
		duration time.Duration
		position Vector
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
	return Vector{}
}

/**********/
/* Player */
/**********/

type PathPlayer struct {
	Path PathInterface
	Loop bool
	time time.Duration
}

func (player *PathPlayer) Step(delta time.Duration) {
	player.time += delta
	duration := player.Path.Duration()
	if player.time > duration {
		switch player.Loop {
		case true:
			player.time -= duration
		case false:
			player.time = duration
		}
	}
}

func (player PathPlayer) Position() Vector {
	return player.Path.At(player.time)
}

func (player *PathPlayer) Reset() {
	player.time = 0
}
