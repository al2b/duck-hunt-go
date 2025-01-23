package mouse

import (
	"duck-hunt-go/engine"
	"embed"
)

const (
	imageWidth  = 3
	imageHeight = 3
)

//go:embed images/*
var imagesFS embed.FS

var (
	image = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/mouse.8.png",
		"images/mouse.24.png",
	))
)
