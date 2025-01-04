package duck

import (
	"duck-hunt-go/engine"
	_ "embed"
)

const (
	spriteWidth  = 32
	spriteHeight = 32
)

var (
	//go:embed sprites.8.png
	sprites8File  []byte
	sprites8Image = engine.Must(engine.LoadImage8(sprites8File))
)

var (
	//go:embed sprites.24.png
	sprites24File  []byte
	sprites24Image = engine.Must(engine.LoadImage24(sprites24File))
)
