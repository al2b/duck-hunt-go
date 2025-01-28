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
	imageMouse = engine.Must(
		engine.LoadImageFile(imagesFS, "images/mouse.png"),
	)
)
