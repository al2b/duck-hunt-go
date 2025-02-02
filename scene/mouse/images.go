package mouse

import (
	"duck-hunt-go/engine"
	"embed"
)

const (
	width  = 3
	height = 3
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageMouse = engine.MustLoadImageFile(imagesFS, "images/mouse.png")
)
