package engine

type Shape [][]float64

func (s Shape) Width() int {
	var minX, maxX float64
	for i, _ := range s {
		for j := 0; j < len(s[i]); j += 2 {
			if s[i][j] < minX {
				minX = s[i][j]
			}
			if s[i][j] > maxX {
				maxX = s[i][j]
			}
		}
	}
	return int(maxX-minX) + 1
}

func (s Shape) Height() int {
	var minY, maxY float64
	for i, _ := range s {
		for j := 1; j < len(s[i]); j += 2 {
			if s[i][j] < minY {
				minY = s[i][j]
			}
			if s[i][j] > maxY {
				maxY = s[i][j]
			}
		}
	}
	return int(maxY-minY) + 1
}

func NewRectangleShape(x1, y1, x2, y2 float64) RectangleShape {
	return RectangleShape{
		x1: x1, y1: y1,
		x2: x2, y2: y2,
	}
}

type RectangleShape struct {
	x1, y1 float64
	x2, y2 float64
}

func (s RectangleShape) Shape() Shape {
	return Shape{
		{
			s.x1, s.y1,
			s.x2, s.y1,
			s.x2, s.y2,
			s.x1, s.y2,
		},
	}
}

func NewPolygonShape(points ...float64) PolygonShape {
	return PolygonShape{
		points: points,
	}
}

type PolygonShape struct {
	points []float64
}

func (s PolygonShape) Shape() Shape {
	return Shape{
		s.points,
	}
}
