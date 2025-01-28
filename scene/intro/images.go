package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageLayout = engine.Must(
		engine.LoadImageFile(imagesFS, "images/layout.png"),
	)
)
