package gun

import (
	"duck-hunt-go/engine"
	_ "embed"
)

var (
	//go:embed sprite.8.png
	sprite8File  []byte
	sprite8Image = engine.Must(engine.LoadImage8(sprite8File))
)

var (
	//go:embed sprite.24.png
	sprite24File  []byte
	sprite24Image = engine.Must(engine.LoadImage24(sprite24File))
)
