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
	image = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/gun.8.png",
		"images/gun.24.png",
	))
)
