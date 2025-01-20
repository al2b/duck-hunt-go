package layout

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
)

const (
	width  = engine.Width
	ground = 184
)

var (
	position      = &engine.Position{X: 0, Y: 0, Z: 100}
	skyPosition   = &engine.Position{X: 0, Y: 0, Z: -math.MaxFloat64}
	treePosition  = &engine.Position{X: 6, Y: 32, Z: 10}
	shrubPosition = &engine.Position{X: 193, Y: 122, Z: 20}
)

func New() *Layout {
	// Model
	m := &Layout{}

	// Bodies
	m.body = engine.NewBody(
		position,
		m.Intersect,
	).Shape(bodyShape)
	m.treeBodies = append(m.treeBodies,
		engine.NewBody(
			treePosition,
			m.Intersect,
		).Shape(treeBodyShapes[0]),
	)
	m.shrubBody = engine.NewBody(
		shrubPosition,
		m.Intersect,
	).Shape(shrubBodyShape)

	return m
}

type Layout struct {
	// Bodies
	body       *engine.Body
	treeBodies []*engine.Body
	shrubBody  *engine.Body
}

func (m *Layout) Init() {}

func (m *Layout) Update(_ []tea.Msg) {}

func (m *Layout) Bodies() (bodies engine.Bodies) {
	return bodies.Append(
		m.body,
		m.treeBodies[0],
		m.shrubBody,
	)
}

func (m *Layout) Intersect() {}

func (m *Layout) Sprites8() (sprites engine.Sprites8) {
	sprites.Append(
		sprite8,
		skySprite8,
		treeSprite8,
		shrubSprite8,
	)

	// Debug
	if engine.Debug() {
		for _, body := range m.Bodies() {
			sprites.Append(body.Sprite8())
		}
	}

	return sprites
}

func (m *Layout) Sprites24() (sprites engine.Sprites24) {
	sprites.Append(
		sprite24,
		skySprite24,
		treeSprite24,
		shrubSprite24,
	)

	// Debug
	if engine.Debug() {
		for _, body := range m.Bodies() {
			sprites.Append(body.Sprite24())
		}
	}

	return sprites
}
