package gun

import (
	"duck-hunt-go/engine"
	"embed"
)

const (
	imageWidth  = 37
	imageHeight = 37
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageGun = engine.Must(
		engine.LoadImageFile(imagesFS, "images/gun.png"),
	)
)
