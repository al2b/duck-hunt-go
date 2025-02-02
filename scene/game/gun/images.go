package gun

import (
	"duck-hunt-go/engine"
	"embed"
)

const (
	width  = 37
	height = 37
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageGun = engine.MustLoadImageFile(imagesFS, "images/gun.png")
)
