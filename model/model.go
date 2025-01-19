package model

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

func New() *Model {
	return &Model{
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

type Model struct {
	models []engine.Model
}

func (g *Model) Init() {
	for _, model := range g.models {
		model.Init()
	}
}

func (g *Model) Update(msgs []tea.Msg) {
	for _, model := range g.models {
		model.Update(msgs)
	}
}

func (g *Model) Bodies() (bodies engine.Bodies) {
	for _, model := range g.models {
		bodies = append(bodies, model.Bodies()...)
	}
	return bodies
}

func (g *Model) Sprites8() (sprites engine.Sprites8) {
	for _, model := range g.models {
		sprites = append(sprites, model.Sprites8()...)
	}

	return sprites
}

func (g *Model) Sprites24() (sprites engine.Sprites24) {
	for _, model := range g.models {
		sprites = append(sprites, model.Sprites24()...)
	}

	return sprites
}
