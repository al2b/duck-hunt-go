package duck

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageDuck = engine.MustLoadImageFile(imagesFS, "images/duck.png")
)
