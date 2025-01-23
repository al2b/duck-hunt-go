package intro

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	image = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/layout.8.png",
		"images/layout.24.png",
	))
)
