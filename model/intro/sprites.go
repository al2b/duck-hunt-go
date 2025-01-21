package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var spritesFS embed.FS

var (
	sprite8 = &engine.Sprite8{
		Position: &engine.Position{X: 0, Y: 0, Z: 0},
		Image:    engine.Must(engine.LoadImage8File(spritesFS, "sprites/layout.8.png")),
	}
)

var (
	sprite24 = &engine.Sprite24{
		Position: &engine.Position{X: 0, Y: 0, Z: 0},
		Image:    engine.Must(engine.LoadImage24File(spritesFS, "sprites/layout.24.png")),
	}
)
