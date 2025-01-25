package engine

import (
	"github.com/charmbracelet/x/ansi"
	"image"
	"image/color"
	"slices"
	"strings"
)

type Renderer interface {
	Name() string
	Ratio() (int, int)
	Render(sprites Sprites, width, height int, resizeWidth, resizeHeight int, paddingHorizontal, paddingVertical int) string
}

/* ******************* */
/* Half Block Bottom 8 */
/* ******************* */

type RendererHalfBlockBottom8 struct{}

func (r *RendererHalfBlockBottom8) Name() string {
	return "Half Block Bottom 8"
}

func (r *RendererHalfBlockBottom8) Ratio() (int, int) {
	return 1, 2
}

func (r *RendererHalfBlockBottom8) Render(sprites Sprites, width, height int, resizeWidth, resizeHeight int, paddingHorizontal, paddingVertical int) string {
	frame := image.NewPaletted(image.Rect(0, 0, width, height), nil)

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
	if resizeWidth != width || resizeHeight != height {
		frame = Image8Resize(frame, resizeWidth, resizeHeight)
	}

	widthRatio, heightRatio := r.Ratio()

	// Padding
	horizontal := strings.Repeat(" ", paddingHorizontal/widthRatio)
	vertical := strings.Repeat("\n", paddingVertical/heightRatio)

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

/* ******************** */
/* Half Block Bottom 24 */
/* ******************** */

type RendererHalfBlockBottom24 struct{}

func (r *RendererHalfBlockBottom24) Name() string {
	return "Half Block Bottom 24"
}

func (r *RendererHalfBlockBottom24) Ratio() (int, int) {
	return 1, 2
}

func (r *RendererHalfBlockBottom24) Render(sprites Sprites, width, height int, resizeWidth, resizeHeight int, paddingHorizontal, paddingVertical int) string {
	frame := image.NewNRGBA(image.Rect(0, 0, width, height))

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
	if resizeWidth != width || resizeHeight != height {
		frame = Image24Resize(frame, resizeWidth, resizeHeight)
	}

	widthRatio, heightRatio := r.Ratio()

	// Padding
	horizontal := strings.Repeat(" ", paddingHorizontal/widthRatio)
	vertical := strings.Repeat("\n", paddingVertical/heightRatio)

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
