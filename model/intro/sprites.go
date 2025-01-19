package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var sprites embed.FS

var sprites8 = map[string]*engine.Sprite8{
	"background": {
		Position: &engine.Position{X: 0, Y: 0, Z: 0},
		Image:    engine.Must(engine.LoadImage8File(sprites, "sprites/background.8.png")),
	},
}

var sprites24 = map[string]*engine.Sprite24{
	"background": {
		Position: &engine.Position{X: 0, Y: 0, Z: 0},
		Image:    engine.Must(engine.LoadImage24File(sprites, "sprites/background.24.png")),
	},
}
