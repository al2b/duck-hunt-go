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
		position: engine.NewPosition(),
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
	// Movement
	movement engine.Vector
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
	m.position.Z = 5 + (rand.Float64() * 20)
	// Movement
	m.movement = engine.VectorFromAngle(
		45 + (rand.Float64() * 90),
	).Scale(1)
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
		case tea.KeyPressMsg:
			switch key := msg.Key(); key.Code {
			case tea.KeyRight:
				m.movement = m.movement.Rotate(-10)
			case tea.KeyLeft:
				m.movement = m.movement.Rotate(10)
			case tea.KeyUp:
				m.position.DepthUp(10)
			case tea.KeyDown:
				m.position.DepthDown(10)
			}
			switch msg.String() {
			case "r":
				m.Init()
			}
		}
	}

	// Move position
	m.position.Move(m.movement)

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
	a := m.movement.Angle()

	switch true {
	default:
		return animationFlyRight
	case 23 <= a && a <= 67:
		return animationFlyTopRight
	case 68 <= a && a <= 112:
		return animationFlyTop
	case 113 <= a && a <= 157:
		return animationFlyTopLeft
	case 158 <= a && a <= 202:
		return animationFlyLeft
	case 203 <= a && a <= 247:
		return animationFlyBottomLeft
	case 248 <= a && a <= 292:
		return animationFlyBottom
	case 293 <= a && a <= 337:
		return animationFlyBottomRight
	}
}
