package duck

import (
	"duck-hunt-go/engine"
	"embed"
)

const (
	spriteWidth  = 32
	spriteHeight = 32
)

//go:embed sprites/*
var sprites embed.FS

var (
	sprites8Image = engine.Must(engine.LoadImage8File(sprites, "sprites/duck.8.png"))
)

var (
	sprites24Image = engine.Must(engine.LoadImage24File(sprites, "sprites/duck.24.png"))
)
