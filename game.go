package main

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/models/background"
	"duck-hunt-go/models/duck"
	"duck-hunt-go/models/gun"
	"duck-hunt-go/models/scene"
	"duck-hunt-go/models/shrub"
	"duck-hunt-go/models/tree"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func NewGame() *Game {
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

func (g *Game) Init() {
	for _, model := range g.models {
		model.Init()
	}
}

func (g *Game) Update(msgs []tea.Msg) {
	for _, model := range g.models {
		model.Update(msgs)
	}
}

func (g *Game) Bodies() (bodies engine.Bodies) {
	for _, model := range g.models {
		bodies = append(bodies, model.Bodies()...)
	}
	return bodies
}

func (g *Game) Sprites8() (sprites engine.Sprites8) {
	for _, model := range g.models {
		sprites = append(sprites, model.Sprites8()...)
	}

	return sprites
}

func (g *Game) Sprites24() (sprites engine.Sprites24) {
	for _, model := range g.models {
		sprites = append(sprites, model.Sprites24()...)
	}

	return sprites
}
