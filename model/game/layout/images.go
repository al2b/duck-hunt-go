package layout

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	image = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/layout.8.png",
		"images/layout.24.png",
	))
	imageSky = engine.NewUniformImage(
		engine.Color8(117),
		engine.Color24{R: 143, G: 192, B: 255, A: 255},
	)
	imageTree = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/tree.8.png",
		"images/tree.24.png",
	))
	imageShrub = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/shrub.8.png",
		"images/shrub.24.png",
	))
)
