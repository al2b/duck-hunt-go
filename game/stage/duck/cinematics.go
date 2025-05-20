package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
)

func cinematicFall(start engine.Vector2D, velocity engine.Vector2D) engine.SequenceCinematic2D {
	img := assets.DuckShotRight

	angle := velocity.Angle()
	if 90 <= angle && angle < 270 {
		img = assets.DuckShotLeft
	}

	return engine.SequenceCinematic2D{
		engine.Cinematic2D{
			{start, img, 23 * config.TickInterval},
		},
	}
}
