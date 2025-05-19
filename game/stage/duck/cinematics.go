package duck

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/game/assets"
	"duck-hunt-go/game/config"
)

func cinematicFall(start engine.Vector2D) engine.SequenceCinematic2D {
	return engine.SequenceCinematic2D{
		engine.Cinematic2D{
			{start, assets.DuckShot, 23 * config.TickInterval},
		},
	}
}
