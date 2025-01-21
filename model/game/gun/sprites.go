package gun

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var spritesFS embed.FS

var (
	sprite8Image = engine.Must(engine.LoadImage8File(spritesFS, "sprites/gun.8.png"))
)

var (
	sprite24Image = engine.Must(engine.LoadImage24File(spritesFS, "sprites/gun.24.png"))
)
