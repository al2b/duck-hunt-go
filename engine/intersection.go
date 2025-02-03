package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/solarlune/resolv"
	"reflect"
)

type Intersection struct {
	Collider, Receiver Body
	IntersectionSets   []resolv.IntersectionSet
}

type Intersections []Intersection

func (i Intersections) From(collider Body) (intersections Intersections) {
	t := reflect.TypeOf(collider)
	for _, intersection := range i {
		if reflect.TypeOf(intersection.Collider) == t {
			intersections = append(intersections, intersection)
		}
	}
	return
}

func (i Intersections) To(receiver any) (intersections Intersections) {
	t := reflect.TypeOf(receiver)
	for _, intersection := range i {
		if reflect.TypeOf(intersection.Receiver) == t {
			intersections = append(intersections, intersection)
		}
	}
	return
}

func (i Intersections) MTV() Vector {
	vec := resolv.NewVector(0, 0)

	for _, intersection := range i {
		for _, intersectionSet := range intersection.IntersectionSets {
			vec = vec.Add(intersectionSet.MTV)
		}
	}

	return Vector{
		X: vec.X,
		Y: vec.Y,
	}
}

func (i Intersections) Normal() Vector {
	return Vector{
		X: i[0].IntersectionSets[0].Intersections[0].Normal.X,
		Y: i[0].IntersectionSets[0].Intersections[0].Normal.Y,
	}
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

			collider := bodies[i]
			receiver := bodies[j]

			intersection := Intersection{
				Collider: collider,
				Receiver: receiver,
			}

			for _, colliderShape := range collider.Shape() {
				for _, receiverShape := range receiver.Shape() {
					receiverResolvShape := resolv.NewConvexPolygon(
						collider.X(),
						collider.Y(),
						colliderShape,
					)
					colliderResolvShape := resolv.NewConvexPolygon(
						receiver.X(),
						receiver.Y(),
						receiverShape,
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
