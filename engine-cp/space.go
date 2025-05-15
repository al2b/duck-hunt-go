package cp

import (
	"duck-hunt-go/engine"
	"github.com/jakecoffman/cp/v2"
	"image/color"
	"math"
)

type SpaceDrawer struct {
	Space *cp.Space
}

func (s SpaceDrawer) Draw(dst *engine.Image) {
	cp.DrawSpace(s.Space, newSpaceDrawer(dst))
}

func newSpaceDrawer(img *engine.Image) spaceDrawer {
	return spaceDrawer{
		img:                 img,
		flags:               cp.DRAW_SHAPES | cp.DRAW_CONSTRAINTS | cp.DRAW_COLLISION_POINTS,
		outlineColor:        cp.FColor{R: 0, G: 1, B: 0, A: 1},
		constraintColor:     cp.FColor{R: 0, G: 0, B: 1, A: 1},
		collisionPointColor: cp.FColor{R: 1, G: 0, B: 0, A: 1},
		data:                nil,
	}
}

type spaceDrawer struct {
	img                                                *engine.Image
	flags                                              uint
	outlineColor, constraintColor, collisionPointColor cp.FColor
	data                                               interface{}
}

func (d spaceDrawer) cpVectorToPoint(v cp.Vector) engine.Point {
	return engine.Pt(int(v.X), int(v.Y))
}

func (d spaceDrawer) cpFColorToColor(fc cp.FColor) color.Color {
	return color.RGBA{R: uint8(fc.R * 255), G: uint8(fc.G * 255), B: uint8(fc.B * 255), A: uint8(fc.A * 255)}
}

func (d spaceDrawer) DrawCircle(pos cp.Vector, angle, radius float64, outline, fill cp.FColor, data interface{}) {
	d.img.Draw(
		engine.Circle{
			d.cpVectorToPoint(pos),
			int(radius),
			d.cpFColorToColor(outline),
		},
		engine.Segment{
			d.cpVectorToPoint(pos),
			d.cpVectorToPoint(pos.Add(cp.ForAngle(angle).Mult(radius))),
			d.cpFColorToColor(outline),
		},
	)
}

func (d spaceDrawer) DrawSegment(a, b cp.Vector, fill cp.FColor, data interface{}) {
	d.img.Draw(
		engine.Segment{
			d.cpVectorToPoint(a),
			d.cpVectorToPoint(b),
			d.cpFColorToColor(fill),
		},
	)
}

func (d spaceDrawer) DrawFatSegment(a, b cp.Vector, radius float64, outline, fill cp.FColor, data interface{}) {
	d.img.Draw(
		engine.Segment{
			d.cpVectorToPoint(a),
			d.cpVectorToPoint(b),
			d.cpFColorToColor(outline),
		},
	)
}

func (d spaceDrawer) DrawPolygon(count int, verts []cp.Vector, radius float64, outline, fill cp.FColor, data interface{}) {
	for i := 0; i < len(verts); i++ {
		d.img.Draw(
			engine.Segment{
				d.cpVectorToPoint(verts[i]),
				d.cpVectorToPoint(verts[(i+1)%len(verts)]),
				d.cpFColorToColor(outline),
			},
		)
	}
}

func (d spaceDrawer) DrawDot(size float64, pos cp.Vector, fill cp.FColor, data interface{}) {
	d.img.Draw(
		engine.Dot{
			d.cpVectorToPoint(pos),
			d.cpFColorToColor(fill),
		},
	)
}

func (d spaceDrawer) Flags() uint {
	return d.flags
}

func (d spaceDrawer) OutlineColor() cp.FColor {
	return d.outlineColor
}

func (d spaceDrawer) ShapeColor(shape *cp.Shape, data interface{}) cp.FColor {
	if shape.Sensor() {
		return cp.FColor{R: 1, G: 1, B: 1, A: .1}
	}

	body := shape.Body()

	if body.IsSleeping() {
		return cp.FColor{R: .2, G: .2, B: .2, A: 1}
	}

	if body.IdleTime() > shape.Space().SleepTimeThreshold {
		return cp.FColor{R: .66, G: .66, B: .66, A: 1}
	}

	val := shape.HashId()

	// scramble the bits up using Robert Jenkins' 32-bit integer hash function
	val = (val + 0x7ed55d16) + (val << 12)
	val = (val ^ 0xc761c23c) ^ (val >> 19)
	val = (val + 0x165667b1) + (val << 5)
	val = (val + 0xd3a2646c) ^ (val << 9)
	val = (val + 0xfd7046c5) + (val << 3)
	val = (val ^ 0xb55a4f09) ^ (val >> 16)

	r := float32((val >> 0) & 0xFF)
	g := float32((val >> 8) & 0xFF)
	b := float32((val >> 16) & 0xFF)

	rgbMax := float32(math.Max(math.Max(float64(r), float64(g)), float64(b)))
	rgbMin := float32(math.Min(math.Min(float64(r), float64(g)), float64(b)))
	var intensity float32
	if body.GetType() == cp.BODY_STATIC {
		intensity = 0.15
	} else {
		intensity = 0.75
	}

	if rgbMin == rgbMax {
		return cp.FColor{R: intensity, A: 1}
	}

	coef := intensity / (rgbMax - rgbMin)
	return cp.FColor{
		R: (r - rgbMin) * coef,
		G: (g - rgbMin) * coef,
		B: (b - rgbMin) * coef,
		A: 1,
	}
}

func (d spaceDrawer) ConstraintColor() cp.FColor {
	return d.constraintColor
}

func (d spaceDrawer) CollisionPointColor() cp.FColor {
	return d.collisionPointColor
}

func (d spaceDrawer) Data() interface{} {
	return d.data
}
