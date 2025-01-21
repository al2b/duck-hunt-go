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
var spritesFS embed.FS

var (
	sprites8Image = engine.Must(engine.LoadImage8File(spritesFS, "sprites/duck.8.png"))
)

var (
	sprites24Image = engine.Must(engine.LoadImage24File(spritesFS, "sprites/duck.24.png"))
)
