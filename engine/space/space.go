package space

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
	"time"
)

func NewSpace() *Space {
	return &Space{
		cpSpace: cp.NewSpace(),
	}
}

type Space struct {
	cpSpace *cp.Space
}

func (s *Space) Init() (cmd tea.Cmd) {
	return nil
}

func (s *Space) Update(_ tea.Msg) tea.Cmd {
	return nil
}

func (s *Space) Clear() {
	gravity := s.cpSpace.Gravity()
	*s = *NewSpace()
	s.cpSpace.SetGravity(gravity)
}

func (s *Space) Step(delta time.Duration) {
	s.cpSpace.Step(delta.Seconds())
}

func (s *Space) Draw(dst *engine.Image) {
	cp.DrawSpace(s.cpSpace, NewDrawer(dst))
}

func (s *Space) SetGravity(gravity engine.Vector2D) *Space {
	s.cpSpace.SetGravity(cp.Vector{
		X: gravity.X,
		Y: gravity.Y,
	})
	return s
}

func (s *Space) AddNewSegment(position0, position1 engine.Vector2D, radius float64) Shape {
	cpShape := s.cpSpace.AddShape(
		cp.NewSegment(
			s.cpSpace.StaticBody,
			cp.Vector{X: position0.X, Y: position0.Y},
			cp.Vector{X: position1.X, Y: position1.Y},
			radius,
		),
	)

	return Shape{
		cpShape: cpShape,
	}
}

func (s *Space) AddNewBody(mass float64) Body {
	cpBody := s.cpSpace.AddBody(
		cp.NewBody(mass, 0),
	)

	return Body{
		cpSpace: s.cpSpace,
		cpBody:  cpBody,
	}
}

func (s *Space) AddNewPositionableBody(positioner engine.Positioner2D) Body {
	cpBody := s.cpSpace.AddBody(
		cp.NewKinematicBody(),
	)

	cpBody.SetPositionUpdateFunc(func(cpBody *cp.Body, deltaTime float64) {
		position := positioner.Position()
		cpBody.SetPosition(
			cp.Vector{
				X: position.X,
				Y: position.Y,
			},
		)
	})

	return Body{
		cpSpace: s.cpSpace,
		cpBody:  cpBody,
	}
}
