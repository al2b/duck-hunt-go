package engine

import (
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

func (r *Intersector) Intersections(bodies Bodies) (intersections Intersections) {
	for i := 0; i < len(bodies); i++ {
		for j := i + 1; j < len(bodies); j++ {
			body1 := bodies[i]
			body2 := bodies[j]
			if intersection1, ok1 := r.Intersection(body1, body2); ok1 {
				intersections = append(intersections, intersection1)
				if intersection2, ok2 := r.Intersection(body2, body1); ok2 {
					intersections = append(intersections, intersection2)
				}
			}
		}
	}

	return
}

func (r *Intersector) Intersection(collider, receiver Body) (intersection Intersection, ok bool) {
	intersection.Collider = collider
	intersection.Receiver = receiver

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
		ok = true
	}

	return
}
