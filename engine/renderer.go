package engine

import (
	"github.com/charmbracelet/x/ansi"
	"image"
	"image/color"
	"slices"
	"strings"
)

func NewRenderer() *Renderer {
	return &Renderer{}
}

type Renderer struct {
	WindowWidth, WindowHeight int
}

func (r *Renderer) Render(model Model, paddingTop, paddingLeft int) string {
	switch mode {
	case Mode8:
		return r.render8(model, paddingTop, paddingLeft)
	case Mode24:
		return r.render24(model, paddingTop, paddingLeft)
	}
	return ""
}

func (r *Renderer) render8(model Model, paddingTop, paddingLeft int) string {
	frame := image.NewPaletted(image.Rect(0, 0, Width, Height), nil)

	sprites := model.Sprites8()

	// Sort by depth (z position)
	slices.SortStableFunc(sprites, func(a, b *Sprite8) int {
		return int(a.Position.Z - b.Position.Z)
	})

	for _, sprite := range sprites {
		if sprite.Image != nil {
			ImageDraw8(frame, sprite.Image, image.Point{X: int(sprite.Position.X), Y: int(sprite.Position.Y)})
		}
	}

	// Zoom
	if zoom != 1 {
		frame = ImageResize8(frame, Width/zoom, Height/zoom)
	}

	// Padding
	top := strings.Repeat("\n", paddingTop)
	left := strings.Repeat(" ", paddingLeft)

	bounds := frame.Bounds()

	str := strings.Builder{}

	str.WriteString(top)
	for y := 0; y < bounds.Max.Y; y += 2 {
		var cct, ccb uint8
		str.WriteString(left)
		for x := 0; x < bounds.Max.X; x++ {
			var t ansi.Style
			// Top
			ct := frame.ColorIndexAt(x, y)
			if ct != cct {
				cct = ct
				t = t.BackgroundColor(ansi.ExtendedColor(cct))
			}
			// Bottom
			cb := frame.ColorIndexAt(x, y+1)
			if cb != ccb {
				ccb = cb
				t = t.ForegroundColor(ansi.ExtendedColor(ccb))
			}
			// Style
			if t != nil {
				str.WriteString(t.String())
			}
			// Block
			str.WriteString("▄")
		}
		str.WriteString(ansi.ResetStyle)
		str.WriteString("\n")
	}

	return str.String()
}

func (r *Renderer) render24(model Model, paddingTop, paddingLeft int) string {
	frame := image.NewNRGBA(image.Rect(0, 0, Width, Height))

	sprites := model.Sprites24()

	// Sort by depth (z position)
	slices.SortStableFunc(sprites, func(a, b *Sprite24) int {
		return int(a.Position.Z - b.Position.Z)
	})

	for _, sprite := range sprites {
		if sprite.Image != nil {
			ImageDraw24(frame, sprite.Image, image.Point{X: int(sprite.Position.X), Y: int(sprite.Position.Y)})
		}
	}

	// Zoom
	if zoom != 1 {
		frame = ImageResize24(frame, Width/zoom, Height/zoom)
	}

	// Padding
	top := strings.Repeat("\n", paddingTop)
	left := strings.Repeat(" ", paddingLeft)

	bounds := frame.Bounds()

	str := strings.Builder{}

	str.WriteString(top)
	for y := 0; y < bounds.Max.Y; y += 2 {
		var cct, ccb color.NRGBA
		str.WriteString(left)
		for x := 0; x < bounds.Max.X; x++ {
			var t ansi.Style
			// Top
			ct := frame.NRGBAAt(x, y)
			ct.A = 0xff
			if ct != cct {
				cct = ct
				t = t.BackgroundColor(cct)
			}
			// Bottom
			cb := frame.NRGBAAt(x, y+1)
			cb.A = 0xff
			if cb != ccb {
				ccb = cb
				t = t.ForegroundColor(ccb)
			}
			// Style
			if t != nil {
				str.WriteString(t.String())
			}
			// Block
			str.WriteString("▄")
		}
		str.WriteString(ansi.ResetStyle)
		str.WriteString("\n")
	}

	return str.String()
}
