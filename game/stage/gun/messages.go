package gun

import "duck-hunt-go/engine"

type ShotMsg engine.Vector2D

func (msg ShotMsg) Position() engine.Vector2D {
	return engine.Vector2D(msg)
}
