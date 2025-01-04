package tree

import (
	"duck-hunt-go/engine"
	_ "embed"
)

var (
	//go:embed sprite.8.png
	sprite8File []byte
	sprite8     = &engine.Sprite8{
		Position: position,
		Image:    engine.Must(engine.LoadImage8(sprite8File)),
	}
)

var (
	//go:embed sprite.24.png
	sprite24File []byte
	sprite24     = &engine.Sprite24{
		Position: position,
		Image:    engine.Must(engine.LoadImage24(sprite24File)),
	}
)
