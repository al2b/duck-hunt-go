package duck

import (
	"duck-hunt-go/engine"
	"embed"
)

//go:embed images/*
var imagesFS embed.FS

var (
	image = engine.Must(engine.LoadImageFiles(imagesFS,
		"images/duck.8.png",
		"images/duck.24.png",
	))
)
