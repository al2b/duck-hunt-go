package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
	"time"
)

func newCinematicShot(position engine.Vector2D, velocity engine.Vector2D) engine.Cinematic2D {
	directioner := engine.VerticalSemicircleDirectioner{velocity}

	return map[engine.Direction]engine.Cinematic2D{
		engine.DirectionRight: engine.Cinematic2D{
			{position, assets.DuckShotRight, 23 * config.TickInterval},
		},
		engine.DirectionLeft: engine.Cinematic2D{
			{position, assets.DuckShotLeft, 23 * config.TickInterval},
		},
	}[directioner.Direction()]
}

func newCinematicFall(position engine.Vector2D, ground int, velocity engine.Vector2D) cinematicFall {
	directioner := engine.VerticalSemicircleDirectioner{velocity}

	return cinematicFall{
		position:  position,
		ground:    ground,
		direction: directioner.Direction(),
	}
}

type cinematicFall struct {
	position  engine.Vector2D
	ground    int
	direction engine.Direction
}

func (c cinematicFall) Duration() time.Duration {
	return time.Duration((float64(c.ground)-c.position.Y)/2) * config.TickInterval
}

func (c cinematicFall) At(duration time.Duration) (engine.Vector2D, *engine.Image) {
	ticks := duration / config.TickInterval

	position := engine.Vector2D{
		c.position.X + float64(ticks%2),
		c.position.Y + float64(2*ticks),
	}

	if (c.direction != engine.DirectionRight) != ((ticks % 10) < 5) {
		return position, assets.DuckFallRight
	} else {
		return position, assets.DuckFallLeft
	}
}
