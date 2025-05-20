package duck

import (
	"duck-hunt-go/engine"
	enginecp "duck-hunt-go/engine-cp"
	"duck-hunt-go/game/config"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"math"
	"math/rand/v2"
)

func New(space *cp.Space, discriminator any) *Duck {
	return &Duck{
		space:         space,
		discriminator: discriminator,
	}
}

type Duck struct {
	state         state
	space         *cp.Space
	discriminator any
	body          *cp.Body
	animationFly  engine.AnimationPlayer
	cinematicFall engine.Cinematic2DPlayer
	engine.ImageDrawer
}

func (m *Duck) Init() tea.Cmd {
	// State
	m.state = stateFly

	// Space
	bodyMass := 1.0
	shapeRadius := 18.0

	m.body = m.space.AddBody(cp.NewBody(bodyMass, cp.MomentForCircle(bodyMass, 0, shapeRadius, cp.Vector{})))
	m.body.UserData = m.discriminator
	m.body.SetPosition(cp.Vector{
		85 + math.Round(rand.Float64()*85),
		config.Ground - 120,
	})
	m.body.SetVelocityVector(cp.
		ForAngle(engine.Radians(235 + (rand.Float64() * 90))).
		Normalize().
		Mult(100),
	)

	shape := m.space.AddShape(cp.NewCircle(m.body, shapeRadius, cp.Vector{}))
	shape.SetFilter(cp.NewShapeFilter(cp.NO_GROUP, config.ShapeCategoryDuck, config.ShapeCategoryLayout|config.ShapeCategoryDuck))
	shape.SetElasticity(1)
	shape.SetFriction(0)

	// Animation
	m.animationFly.Animation = animationFly{enginecp.VelocityVelociter{m.body}}
	m.animationFly.OnEnd = engine.PlayerOnEndLoop
	m.animationFly.Play()

	// Drawer
	m.ImageDrawer.Pointer = engine.PointAdder{
		enginecp.PositionPointer{m.body},
		engine.Pt(-19, -19),
	}
	m.ImageDrawer.Imager = &m.animationFly

	return nil
}

func (m *Duck) Update(msg tea.Msg) tea.Cmd {
	switch m.state {

	// Fly
	case stateFly:
		switch msg := msg.(type) {
		case engine.TickMsg:
			// Animation
			m.animationFly.Step(msg.Interval)
		case ShotMsg:
			// State
			m.state = stateFall

			// Space
			m.body.EachShape(func(shape *cp.Shape) {
				m.space.RemoveShape(shape)
			})
			m.space.RemoveBody(m.body)

			// Cinematic
			position := m.body.Position()
			m.cinematicFall.Cinematic = cinematicFall(engine.Vec2D(position.X, position.Y))
			m.cinematicFall.OnEnd = engine.PlayerOnEndLoop
			m.cinematicFall.Play()

			// Drawer
			m.ImageDrawer.Pointer = engine.PointAdder{
				engine.Position2DPointer{&m.cinematicFall},
				engine.Pt(-19, -19),
			}
			m.ImageDrawer.Imager = &m.cinematicFall
		}

	// Fall
	case stateFall:
		switch msg := msg.(type) {
		case engine.TickMsg:
			// Cinematic
			m.cinematicFall.Step(msg.Interval)
		}

	}

	return nil
}

type state int

const (
	stateIdle state = iota
	stateFly
	stateFall
)
