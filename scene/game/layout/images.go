package layout

import (
	"duck-hunt-go/engine"
	"embed"
	"image"
	"image/color"
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageLayout = engine.Must(
		engine.LoadImageFile(imagesFS, "images/layout.png"),
	)
	imageSky  = image.NewUniform(color.NRGBA{R: 143, G: 192, B: 255, A: 255})
	imageTree = engine.Must(
		engine.LoadImageFile(imagesFS, "images/tree.png"),
	)
	imageShrub = engine.Must(
		engine.LoadImageFile(imagesFS, "images/shrub.png"),
	)
)
