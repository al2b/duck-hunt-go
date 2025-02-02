package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/solarlune/resolv"
)

type Intersection struct {
	Body1, Body2     Body
	IntersectionSets []resolv.IntersectionSet
}

type Intersections []Intersection

type IntersectionsMsg struct {
	Intersections
}

func NewIntersector() *Intersector {
	return &Intersector{}
}

type Intersector struct{}

func (r *Intersector) Intersect(bodies Bodies) tea.Cmd {
	var intersections Intersections

	for i := 0; i < len(bodies); i++ {
		for j := 0; j < len(bodies); j++ {
			if i == j {
				continue
			}

			body1 := bodies[i]
			body2 := bodies[j]

			intersection := Intersection{
				Body1: body1,
				Body2: body2,
			}

			for _, shape1 := range body1.Shape() {
				for _, shape2 := range body2.Shape() {
					resolvShape1 := resolv.NewConvexPolygon(
						body1.X(),
						body1.Y(),
						shape1,
					)
					resolvShape2 := resolv.NewConvexPolygon(
						body2.X(),
						body2.Y(),
						shape2,
					)
					if intersectionSet := resolvShape1.Intersection(resolvShape2); !intersectionSet.IsEmpty() {
						intersection.IntersectionSets = append(intersection.IntersectionSets, intersectionSet)
					}
				}
			}

			if len(intersection.IntersectionSets) > 0 {
				intersections = append(intersections, intersection)
			}
		}
	}

	if len(intersections) == 0 {
		return nil
	}

	return func() tea.Msg {
		return IntersectionsMsg{Intersections: intersections}
	}
}
