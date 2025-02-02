package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/solarlune/resolv"
	"reflect"
	"slices"
)

type Intersection struct {
	Receiver, Collider Body
	IntersectionSets   []resolv.IntersectionSet
}

type Intersections []Intersection

func (i Intersections) From(receiver Body) Intersections {
	t := reflect.TypeOf(receiver)
	return slices.DeleteFunc(i, func(i Intersection) bool {
		return reflect.TypeOf(i.Receiver) != t
	})
}

func (i Intersections) To(collider any) Intersections {
	t := reflect.TypeOf(collider)
	return slices.DeleteFunc(i, func(i Intersection) bool {
		return reflect.TypeOf(i.Collider) != t
	})
}

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

			receiver := bodies[i]
			collider := bodies[j]

			intersection := Intersection{
				Receiver: receiver,
				Collider: collider,
			}

			for _, receiverShape := range receiver.Shape() {
				for _, colliderShape := range collider.Shape() {
					receiverResolvShape := resolv.NewConvexPolygon(
						receiver.X(),
						receiver.Y(),
						receiverShape,
					)
					colliderResolvShape := resolv.NewConvexPolygon(
						collider.X(),
						collider.Y(),
						colliderShape,
					)
					if intersectionSet := receiverResolvShape.Intersection(colliderResolvShape); !intersectionSet.IsEmpty() {
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
