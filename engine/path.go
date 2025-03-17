package engine

import (
	"math"
	"time"
)

type PathInterface interface {
	Duration() time.Duration
	At(time.Duration) Vector
}

type FixedPath struct {
	Position Vector
}

func (path FixedPath) Duration() time.Duration {
	return 0
}

func (path FixedPath) At(_ time.Duration) Vector {
	return path.Position
}

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

/**********/
/* Player */
/**********/

type PathPlayer struct {
	Path PathInterface
	time time.Duration
}

func (player *PathPlayer) Step(delta time.Duration) {
	player.time += delta
	duration := player.Path.Duration()
	if player.time >= duration {
		player.time -= duration
	}
}

func (player PathPlayer) Position() Vector {
	return player.Path.At(player.time)
}
