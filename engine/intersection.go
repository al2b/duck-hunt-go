package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/solarlune/resolv"
)

type Intersection struct {
	Body            Body
	IntersectionSet resolv.IntersectionSet
}

type IntersectionMsg struct {
	Intersection
}

type Intersections []Intersection

func NewIntersector() *Intersector {
	return &Intersector{}
}

type Intersector struct{}

func (r *Intersector) Intersect(model Model) tea.Cmd {
	var cmds []tea.Cmd

	bodies := model.Bodies()

	// Empty bodies intersections

	for i := 0; i < len(bodies); i++ {
		for j := 0; j < len(bodies); j++ {
			if i == j {
				continue
			}
			body1 := bodies[i]
			body2 := bodies[j]
			if intersectionSet := r.resolvShape(body1).Intersection(r.resolvShape(body2)); !intersectionSet.IsEmpty() {
				cmds = append(cmds,
					body1.Update(IntersectionMsg{
						Intersection{
							Body:            body2,
							IntersectionSet: intersectionSet,
						},
					}))
			}
		}
	}

	return tea.Batch(cmds...)
}

func (r *Intersector) resolvShape(body Body) *resolv.ConvexPolygon {
	var points []float64

	for _, sp := range body.Shape() {
		points = append(points, sp.X, sp.Y)
	}

	return resolv.NewConvexPolygon(
		body.X(),
		body.Y(),
		points,
	)
}
