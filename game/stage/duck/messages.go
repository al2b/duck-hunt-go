package duck

import "duck-hunt-go/engine"

type ShotMsg engine.Vector2D

type DiscriminatedShotMsg struct {
	ShotMsg
	Discriminator any
}
