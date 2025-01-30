package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	imageLayout = engine.MustLoadImageFile(imagesFS, "images/layout.png")
)
