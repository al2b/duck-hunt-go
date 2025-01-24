package engine

import (
	"github.com/charmbracelet/x/ansi"
	"image"
	"image/color"
	"slices"
	"strings"
)

func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		width:  width,
		height: height,
	}
}

type Renderer struct {
	width, height int
}

func (r *Renderer) WidthRatio() int {
	return 1
}

func (r *Renderer) HeightRatio() int {
	return 2
}

func (r *Renderer) Render(sprites Sprites, width, height int, paddingHorizontal, paddingVertical int) string {
	switch mode {
	case Mode8:
		return r.render8(sprites, width, height, paddingHorizontal, paddingVertical)
	case Mode24:
		return r.render24(sprites, width, height, paddingHorizontal, paddingVertical)
	}
	return ""
}

func (r *Renderer) render8(sprites Sprites, width, height int, paddingHorizontal, paddingVertical int) string {
	frame := image.NewPaletted(image.Rect(0, 0, r.width, r.height), nil)

	// Sort by depth (z coordinate)
	slices.SortStableFunc(sprites, func(a, b Sprite) int {
		return int(a.Z() - b.Z())
	})

	for _, sprite := range sprites {
		img := sprite.Image8()
		if img != nil {
			Image8Draw(frame, img, image.Point{X: int(sprite.X()), Y: int(sprite.Y())})
		}
	}

	// Resize
	if width != r.width || height != r.height {
		frame = Image8Resize(frame, width, height)
	}

	// Padding
	horizontal := strings.Repeat(" ", paddingHorizontal/r.WidthRatio())
	vertical := strings.Repeat("\n", paddingVertical/r.HeightRatio())

	bounds := frame.Bounds()

	str := strings.Builder{}

	str.WriteString(vertical)
	for y := 0; y < bounds.Max.Y; y += 2 {
		var cct, ccb uint8
		str.WriteString(horizontal)
		for x := 0; x < bounds.Max.X; x++ {
			var t ansi.Style
			// Block top
			ct := frame.ColorIndexAt(x, y)
			if ct != cct {
				cct = ct
				t = t.BackgroundColor(ansi.ExtendedColor(cct))
			}
			// Block bottom
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

func (r *Renderer) render24(sprites Sprites, width, height int, paddingHorizontal, paddingVertical int) string {
	frame := image.NewNRGBA(image.Rect(0, 0, r.width, r.height))

	// Sort by depth (z coordinate)
	slices.SortStableFunc(sprites, func(a, b Sprite) int {
		return int(a.Z() - b.Z())
	})

	for _, sprite := range sprites {
		img := sprite.Image24()
		if img != nil {
			Image24Draw(frame, img, image.Point{X: int(sprite.X()), Y: int(sprite.Y())})
		}
	}

	// Resize
	if width != r.width || height != r.height {
		frame = Image24Resize(frame, width, height)
	}

	// Padding
	horizontal := strings.Repeat(" ", paddingHorizontal/r.WidthRatio())
	vertical := strings.Repeat("\n", paddingVertical/r.HeightRatio())

	bounds := frame.Bounds()

	str := strings.Builder{}

	str.WriteString(vertical)
	for y := 0; y < bounds.Max.Y; y += 2 {
		var cct, ccb color.NRGBA
		str.WriteString(horizontal)
		for x := 0; x < bounds.Max.X; x++ {
			var t ansi.Style
			// Block top
			ct := frame.NRGBAAt(x, y)
			ct.A = 0xff
			if ct != cct {
				cct = ct
				t = t.BackgroundColor(cct)
			}
			// Block bottom
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
