package engine

import (
	"fmt"
	"github.com/charmbracelet/x/ansi"
	"image"
	"image/color"
	"strings"
)

type Renderer interface {
	Ratio() (int, int)
	Render(img *image.NRGBA, paddingH, paddingV int) string
	fmt.Stringer
}

func NewRenderers() *Renderers {
	return &Renderers{
		renderers: []Renderer{
			NewRendererHalfBlockTop(),
			NewRendererHalfBlockBottom(),
			NewRendererHalfBlockTopExtended(),
			NewRendererHalfBlockBottomExtended(),
			NewRendererHalfBlockTopSystem(),
			NewRendererHalfBlockBottomSystem(),
			NewRendererHalfBlockTopGrayscale(),
			NewRendererHalfBlockBottomGrayscale(),
			NewRendererHalfBlockTopMonochrome(),
			NewRendererHalfBlockBottomMonochrome(),
		},
	}
}

type Renderers struct {
	renderers []Renderer
	current   int
}

func (r *Renderers) Current() Renderer {
	return r.renderers[r.current]
}

func (r *Renderers) Next() Renderer {
	r.current = (r.current + 1) % len(r.renderers)
	return r.Current()
}

func (r *Renderers) Previous() Renderer {
	r.current = (r.current - 1 + len(r.renderers)) % len(r.renderers)
	return r.Current()
}

/* **************** */
/* Half Block - Top */
/* **************** */

func NewRendererHalfBlockTop() *RendererHalfBlockTop {
	return &RendererHalfBlockTop{
		RendererHalfBlock: &RendererHalfBlock{
			Top: true,
		},
	}
}

type RendererHalfBlockTop struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockTop) String() string {
	return "Half Block - Top"
}

/* ******************* */
/* Half Block - Bottom */
/* ******************* */

func NewRendererHalfBlockBottom() *RendererHalfBlockBottom {
	return &RendererHalfBlockBottom{
		RendererHalfBlock: &RendererHalfBlock{
			Top: false,
		},
	}
}

type RendererHalfBlockBottom struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockBottom) String() string {
	return "Half Block - Bottom"
}

/* *************************** */
/* Half Block - Top - Extended */
/* *************************** */

func NewRendererHalfBlockTopExtended() *RendererHalfBlockTopExtended {
	return &RendererHalfBlockTopExtended{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    true,
			Binder: NewColorBinder(ColorBindingExtended()),
		},
	}
}

type RendererHalfBlockTopExtended struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockTopExtended) String() string {
	return "Half Block - Top - Extended"
}

/* ****************************** */
/* Half Block - Bottom - Extended */
/* ****************************** */

func NewRendererHalfBlockBottomExtended() *RendererHalfBlockBottomExtended {
	return &RendererHalfBlockBottomExtended{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    false,
			Binder: NewColorBinder(ColorBindingExtended()),
		},
	}
}

type RendererHalfBlockBottomExtended struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockBottomExtended) String() string {
	return "Half Block - Bottom - Extended"
}

/* ************************* */
/* Half Block - Top - System */
/* ************************* */

func NewRendererHalfBlockTopSystem() *RendererHalfBlockTopSystem {
	return &RendererHalfBlockTopSystem{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    true,
			Binder: NewColorBinder(ColorBindingSystem()),
		},
	}
}

type RendererHalfBlockTopSystem struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockTopSystem) String() string {
	return "Half Block - Top - System"
}

/* **************************** */
/* Half Block - Bottom - System */
/* **************************** */

func NewRendererHalfBlockBottomSystem() *RendererHalfBlockBottomSystem {
	return &RendererHalfBlockBottomSystem{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    false,
			Binder: NewColorBinder(ColorBindingSystem()),
		},
	}
}

type RendererHalfBlockBottomSystem struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockBottomSystem) String() string {
	return "Half Block - Bottom - System"
}

/* **************************** */
/* Half Block - Top - Grayscale */
/* **************************** */

func NewRendererHalfBlockTopGrayscale() *RendererHalfBlockTopGrayscale {
	return &RendererHalfBlockTopGrayscale{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    true,
			Binder: NewColorBinder(ColorBindingGrayscale()),
		},
	}
}

type RendererHalfBlockTopGrayscale struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockTopGrayscale) String() string {
	return "Half Block - Top - Grayscale"
}

/* ******************************* */
/* Half Block - Bottom - Grayscale */
/* ******************************* */

func NewRendererHalfBlockBottomGrayscale() *RendererHalfBlockBottomGrayscale {
	return &RendererHalfBlockBottomGrayscale{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    false,
			Binder: NewColorBinder(ColorBindingGrayscale()),
		},
	}
}

type RendererHalfBlockBottomGrayscale struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockBottomGrayscale) String() string {
	return "Half Block - Bottom - Grayscale"
}

/* ***************************** */
/* Half Block - Top - Monochrome */
/* ***************************** */

func NewRendererHalfBlockTopMonochrome() *RendererHalfBlockTopMonochrome {
	return &RendererHalfBlockTopMonochrome{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    true,
			Binder: NewColorBinder(ColorBindingMonochrome()),
		},
	}
}

type RendererHalfBlockTopMonochrome struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockTopMonochrome) String() string {
	return "Half Block - Top - Monochrome"
}

/* ******************************** */
/* Half Block - Bottom - Monochrome */
/* ******************************** */

func NewRendererHalfBlockBottomMonochrome() *RendererHalfBlockBottomMonochrome {
	return &RendererHalfBlockBottomMonochrome{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    false,
			Binder: NewColorBinder(ColorBindingMonochrome()),
		},
	}
}

type RendererHalfBlockBottomMonochrome struct {
	*RendererHalfBlock
}

func (r *RendererHalfBlockBottomMonochrome) String() string {
	return "Half Block - Bottom - Monochrome"
}

/* ********** */
/* Half Block */
/* ********** */

type RendererHalfBlock struct {
	Binder *ColorBinder
	Top    bool
}

func (r *RendererHalfBlock) Ratio() (int, int) {
	return 1, 2
}

func (r *RendererHalfBlock) Render(img *image.NRGBA, padH, padV int) string {
	// Padding
	ratioW, ratioH := r.Ratio()
	padHStr := strings.Repeat(" ", padH/ratioW)
	padVStr := strings.Repeat("\n", padV/ratioH)

	// Image bounds
	bounds := img.Bounds()

	// String
	str := strings.Builder{}
	str.WriteString(padVStr)

	for y := 0; y < bounds.Max.Y; y += 2 {
		str.WriteString(padHStr)
		var cct, ccb color.Color
		for x := 0; x < bounds.Max.X; x++ {
			var t ansi.Style
			// Color top
			var ct color.Color
			ct = img.NRGBAAt(x, y)
			if r.Binder != nil {
				ct = r.Binder.Bind(ct)
			}
			if ct != cct {
				cct = ct
				if r.Top {
					t = t.ForegroundColor(cct)
				} else {
					t = t.BackgroundColor(cct)
				}
			}
			//  Color bottom
			var cb color.Color
			cb = img.NRGBAAt(x, y+1)
			if r.Binder != nil {
				cb = r.Binder.Bind(cb)
			}
			if cb != ccb {
				ccb = cb
				if r.Top {
					t = t.BackgroundColor(ccb)
				} else {
					t = t.ForegroundColor(ccb)
				}
			}
			// Style
			if t != nil {
				str.WriteString(t.String())
			}
			// Half Block
			if r.Top {
				str.WriteString("▀")
			} else {
				str.WriteString("▄")
			}
		}
		str.WriteString(ansi.ResetStyle + "\n")
	}

	return str.String()
}
