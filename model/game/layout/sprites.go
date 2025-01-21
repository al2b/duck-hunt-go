package layout

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
	// Sky
	engine.NewUniformSprite(skyCoordinates,
		engine.Color8(117),
		engine.Color24{R: 143, G: 192, B: 255, A: 255},
	),
	// Tree
	engine.NewImageSprite(treeCoordinates,
		engine.Must(engine.LoadImage8File(spritesFS, "sprites/tree.8.png")),
		engine.Must(engine.LoadImage24File(spritesFS, "sprites/tree.24.png")),
	),
	// Shrub
	engine.NewImageSprite(shrubCoordinates,
		engine.Must(engine.LoadImage8File(spritesFS, "sprites/shrub.8.png")),
		engine.Must(engine.LoadImage24File(spritesFS, "sprites/shrub.24.png")),
	),
}
