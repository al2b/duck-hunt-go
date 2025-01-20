package gun

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var sprites embed.FS

var (
	sprite8Image = engine.Must(engine.LoadImage8File(sprites, "sprites/gun.8.png"))
)

var (
	sprite24Image = engine.Must(engine.LoadImage24File(sprites, "sprites/gun.24.png"))
)
