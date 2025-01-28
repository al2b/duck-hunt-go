package duck

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageDuck = engine.Must(
		engine.LoadImageFile(imagesFS, "images/duck.png"),
	)
)
