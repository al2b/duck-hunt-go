package game

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/model/game/background"
	"duck-hunt-go/model/game/duck"
	"duck-hunt-go/model/game/gun"
	"duck-hunt-go/model/game/scene"
	"duck-hunt-go/model/game/shrub"
	"duck-hunt-go/model/game/tree"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func New() *Game {
	return &Game{
		models: []engine.Model{
			background.New(),
			duck.New(),
			tree.New(),
			shrub.New(),
			scene.New(),
			gun.New(),
		},
	}
}

type Game struct {
	models []engine.Model
}

func (m Game) Init() {
	for _, model := range m.models {
		model.Init()
	}
}

func (m Game) Update(msgs []tea.Msg) {
	for _, model := range m.models {
		model.Update(msgs)
	}
}

func (m Game) Bodies() (bodies engine.Bodies) {
	for _, model := range m.models {
		bodies = append(bodies, model.Bodies()...)
	}
	return bodies
}

func (m Game) Sprites8() (sprites engine.Sprites8) {
	for _, model := range m.models {
		sprites = append(sprites, model.Sprites8()...)
	}
	return sprites
}

func (m Game) Sprites24() (sprites engine.Sprites24) {
	for _, model := range m.models {
		sprites = append(sprites, model.Sprites24()...)
	}
	return sprites
}
