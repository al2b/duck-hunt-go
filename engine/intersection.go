package engine

import (
	"github.com/solarlune/resolv"
)

type Intersection struct {
	Body            *Body
	IntersectionSet resolv.IntersectionSet
}

type Intersections []Intersection

func NewIntersector() *Intersector {
	return &Intersector{}
}

type Intersector struct{}

func (r *Intersector) Intersect(model Model) {
	bodies := model.Bodies()

	// Empty bodies intersections
	for _, body := range bodies {
		body.Intersections = nil
	}

	for i := 0; i < len(bodies); i++ {
		for j := 0; j < len(bodies); j++ {
			if i == j {
				continue
			}
			body1 := bodies[i]
			body2 := bodies[j]
			if intersectionSet := body1.ResolvShape().Intersection(body2.ResolvShape()); !intersectionSet.IsEmpty() {
				body1.Intersections = append(body1.Intersections, Intersection{
					Body:            body2,
					IntersectionSet: intersectionSet,
				})
				if body1.onIntersect != nil {
					body1.onIntersect()
				}
			}
		}
	}
}
