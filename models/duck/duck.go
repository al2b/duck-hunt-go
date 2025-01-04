package duck

import (
	"duck-hunt-go/engine"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
	"math/rand/v2"
)

func New() *Duck {
	// Model
	m := &Duck{
		position:  engine.NewPosition(),
		direction: engine.NewDirection(),
	}

	// Body
	m.body = engine.NewBody(
		m.position,
		m.Intersect,
	).Shape(bodyShape)

	return m
}

type Duck struct {
	// Position
	position *engine.Position
	// Direction
	direction *engine.Direction
	// Animation
	animationFrame    int
	animationVelocity int
	// Body
	body *engine.Body
}

func (m *Duck) Init() {
	// Position
	m.position.X = 85 + (rand.Float64() * 85)
	m.position.Y = 160
	m.position.Y = 100
	m.position.Z = 5 + (rand.Float64() * 20)
	// Direction
	m.direction.Angle = (315 + rand.IntN(90)) % 360
	m.direction.Velocity = 1
	// Animation
	m.animationFrame = 0
	m.animationVelocity = 6
}

func (m *Duck) Update(msgs []tea.Msg) {
	// Messages
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case tea.MouseMotionMsg:
			m.position.X = float64(msg.X)
			m.position.Y = float64(msg.Y)
		case tea.KeyMsg:
			switch key := msg.Key(); key.Code {
			case tea.KeyRight:
				m.direction.RotateUp(10)
			case tea.KeyLeft:
				m.direction.RotateDown(10)
			case tea.KeyUp:
				m.position.DepthUp(10)
			case tea.KeyDown:
				m.position.DepthDown(10)
			}
		}
	}

	// Position
	m.position.Move(m.direction)

	// Animation
	m.animationFrame = (m.animationFrame + 1) % (len(animationFrames[m.animation()]) * m.animationVelocity)
}

func (m *Duck) Bodies() (bodies engine.Bodies) {
	return bodies.Append(m.body)
}

func (m *Duck) Intersect() {}

func (m *Duck) Sprites8() (sprites engine.Sprites8) {
	frame := animationFrames[m.animation()][m.animationFrame/m.animationVelocity]

	img := engine.ImageCrop8(sprites8Image, image.Rect(
		frame.X*spriteWidth,
		frame.Y*spriteHeight,
		(frame.X+1)*spriteWidth,
		(frame.Y+1)*spriteHeight,
	))

	if frame.FlipH {
		img = engine.ImageFlipH8(img)
	}

	if frame.FlipV {
		img = engine.ImageFlipV8(img)
	}

	sprites.Append(&engine.Sprite8{
		Position: m.position,
		Image:    img,
	})

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite8())
	}

	return sprites
}

func (m *Duck) Sprites24() (sprites engine.Sprites24) {
	frame := animationFrames[m.animation()][m.animationFrame/m.animationVelocity]

	img := engine.ImageCrop24(sprites24Image, image.Rect(
		frame.X*spriteWidth,
		frame.Y*spriteHeight,
		(frame.X+1)*spriteWidth,
		(frame.Y+1)*spriteHeight,
	))

	if frame.FlipH {
		img = engine.ImageFlipH24(img)
	}

	if frame.FlipV {
		img = engine.ImageFlipV24(img)
	}

	sprites.Append(&engine.Sprite24{
		Position: m.position,
		Image:    img,
	})

	// Debug
	if engine.Debug() {
		sprites.Append(m.body.Sprite24())
	}

	return sprites
}

func (m *Duck) animation() engine.Animation {
	angle := m.direction.Angle

	switch true {
	default:
		return animationFlyTop
	case 23 <= angle && angle <= 67:
		return animationFlyTopRight
	case 68 <= angle && angle <= 112:
		return animationFlyRight
	case 113 <= angle && angle <= 157:
		return animationFlyBottomRight
	case 158 <= angle && angle <= 202:
		return animationFlyBottom
	case 203 <= angle && angle <= 247:
		return animationFlyBottomLeft
	case 248 <= angle && angle <= 292:
		return animationFlyLeft
	case 293 <= angle && angle <= 337:
		return animationFlyTopLeft
	}
}
