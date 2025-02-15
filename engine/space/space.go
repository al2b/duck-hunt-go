package space

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/jakecoffman/cp/v2"
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

func (s *Space) Step(deltaTime float64) {
	s.cpSpace.Step(deltaTime)
}

func (s *Space) Draw(scene *engine.Image) {
	cp.DrawSpace(s.cpSpace, NewDrawer(scene))
}

func (s *Space) SetGravity(vector engine.Vector) *Space {
	s.cpSpace.SetGravity(cp.Vector{
		X: vector.X,
		Y: vector.Y,
	})
	return s
}

func (s *Space) AddNewSegment(p0, p1 engine.Position, radius float64) Shape {
	cpShape := s.cpSpace.AddShape(
		cp.NewSegment(
			s.cpSpace.StaticBody,
			cp.Vector{X: p0.X, Y: p0.Y},
			cp.Vector{X: p1.X, Y: p1.Y},
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

func (s *Space) AddNewPositionableBody(m engine.Positioner) Body {
	cpBody := s.cpSpace.AddBody(
		cp.NewKinematicBody(),
	)

	cpBody.SetPositionUpdateFunc(func(cpBody *cp.Body, dt float64) {
		position := m.Position()
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
