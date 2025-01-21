package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var spritesFS embed.FS

var sprites = engine.Sprites{
	engine.NewImageSprite(coordinates,
		engine.Must(engine.LoadImage8File(spritesFS, "sprites/layout.8.png")),
		engine.Must(engine.LoadImage24File(spritesFS, "sprites/layout.24.png")),
	),
}
