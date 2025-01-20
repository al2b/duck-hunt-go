package layout

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed sprites/*
var sprites embed.FS

var (
	sprite8 = &engine.Sprite8{
		Position: position,
		Image:    engine.Must(engine.LoadImage8File(sprites, "sprites/layout.8.png")),
	}
	skySprite8 = &engine.Sprite8{
		Position: skyPosition,
		Image:    engine.NewImageUniform8(engine.Color8(117)),
	}
	treeSprite8 = &engine.Sprite8{
		Position: treePosition,
		Image:    engine.Must(engine.LoadImage8File(sprites, "sprites/tree.8.png")),
	}
	shrubSprite8 = &engine.Sprite8{
		Position: shrubPosition,
		Image:    engine.Must(engine.LoadImage8File(sprites, "sprites/shrub.8.png")),
	}
)

var (
	sprite24 = &engine.Sprite24{
		Position: position,
		Image:    engine.Must(engine.LoadImage24File(sprites, "sprites/layout.24.png")),
	}
	skySprite24 = &engine.Sprite24{
		Position: skyPosition,
		Image:    engine.NewImageUniform24(engine.Color24{R: 143, G: 192, B: 255, A: 255}),
	}
	treeSprite24 = &engine.Sprite24{
		Position: treePosition,
		Image:    engine.Must(engine.LoadImage24File(sprites, "sprites/tree.24.png")),
	}
	shrubSprite24 = &engine.Sprite24{
		Position: shrubPosition,
		Image:    engine.Must(engine.LoadImage24File(sprites, "sprites/shrub.24.png")),
	}
)
