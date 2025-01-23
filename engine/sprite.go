package engine

type Sprite interface {
	Point
	Image
}

type Sprites []Sprite

type CoordinatedSprite struct {
	Coordinates
	Image
}
