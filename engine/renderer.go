package engine

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-ciede2000"
	"image/color"
	"math"
	"strings"
)

type Renderer interface {
	fmt.Stringer
	Support(profile colorprofile.Profile) bool
	Ratio() Size
	Render(img *Image, padding Size) string
}

func NewRenderers() *Renderers {
	mixed := NewRendererMixedBlockAscii(ColorWhite, ColorBlack)

	return &Renderers{
		discard: RendererDiscard{},
		mixed:   mixed,
		available: []Renderer{
			NewRendererHalfBlockTrueColor(true),
			NewRendererHalfBlockTrueColor(false),
			NewRendererHalfBlockANSI256(true, ColorBindingANSI256{}),
			NewRendererHalfBlockANSI256(false, ColorBindingANSI256{}),
			NewRendererHalfBlockANSI256(true, ColorBindingANSI256Grayscale{}),
			NewRendererHalfBlockANSI256(false, ColorBindingANSI256Grayscale{}),
			NewRendererHalfBlockANSI256(true, ColorBindingANSI256BlackAndWhite{}),
			NewRendererHalfBlockANSI256(false, ColorBindingANSI256BlackAndWhite{}),
			NewRendererHalfBlockANSI(true, ColorBindingANSI{}),
			NewRendererHalfBlockANSI(false, ColorBindingANSI{}),
			NewRendererHalfBlockANSI(true, ColorBindingANSIGrayscale{}),
			NewRendererHalfBlockANSI(false, ColorBindingANSIGrayscale{}),
			NewRendererHalfBlockANSI(true, ColorBindingANSIBlackAndWhite{}),
			NewRendererHalfBlockANSI(false, ColorBindingANSIBlackAndWhite{}),
			NewRendererAscii(".,:-=i|%O#@$X"),
			mixed,
		},
	}
}

type Renderers struct {
	discard   Renderer
	mixed     *RendererMixedBlockAscii
	available []Renderer
	enabled   []Renderer
	current   int
}

func (r *Renderers) Init() tea.Cmd {
	return tea.Batch(
		LogInfo("Initialize renderers..."),
		tea.RequestForegroundColor,
		tea.RequestBackgroundColor,
	)
}

func (r *Renderers) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.ColorProfileMsg:
		return r.updateProfile(msg.Profile)
	case tea.ForegroundColorMsg:
		r.mixed.SetForeground(msg.Color)
	case tea.BackgroundColorMsg:
		r.mixed.SetBackground(msg.Color)
	}
	return nil
}

func (r *Renderers) updateProfile(profile colorprofile.Profile) tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds,
		LogInfo("Update color profile", "profile", profile),
	)

	// Enable supported renderers
	for _, renderer := range r.available {
		if renderer.Support(profile) {
			cmds = append(cmds,
				LogInfo("Enable renderer", "renderer", renderer),
			)
			r.enabled = append(r.enabled, renderer)
		} else {
			cmds = append(cmds,
				LogInfo("Disable renderer", "renderer", renderer),
			)
		}
	}

	return tea.Batch(cmds...)
}

func (r *Renderers) Current() Renderer {
	if len(r.enabled) == 0 {
		return r.discard
	}

	return r.enabled[r.current]
}

func (r *Renderers) Next() Renderer {
	r.current = (r.current + 1) % len(r.enabled)
	return r.Current()
}

func (r *Renderers) Previous() Renderer {
	r.current = (r.current - 1 + len(r.enabled)) % len(r.enabled)
	return r.Current()
}

/* ******* */
/* Discard */
/* ******* */

type RendererDiscard struct{}

func (r RendererDiscard) String() string                      { return "Discard" }
func (r RendererDiscard) Support(_ colorprofile.Profile) bool { return true }
func (r RendererDiscard) Ratio() Size                         { return Size{Width: 1, Height: 2} }
func (r RendererDiscard) Render(_ *Image, _ Size) string      { return "" }

/* ********** */
/* Half Block */
/* ********** */

type RendererHalfBlock struct {
	Binder *ColorBinder
	Top    bool
}

func (r *RendererHalfBlock) String() string {
	return fmt.Sprintf("Half Block - %s",
		map[bool]string{true: "Top", false: "Bottom"}[r.Top],
	)
}

func (r *RendererHalfBlock) Ratio() Size { return Size{Width: 1, Height: 2} }

func (r *RendererHalfBlock) Render(img *Image, padding Size) string {
	// Padding
	ratio := r.Ratio()
	paddingWidthStr := strings.Repeat(" ", padding.Width/ratio.Width)
	paddingHeightStr := strings.Repeat("\n", padding.Height/ratio.Height)

	// Image bounds
	bounds := img.Bounds()

	// String
	str := strings.Builder{}
	str.WriteString(paddingHeightStr)

	for y := 0; y < bounds.Max.Y; y += 2 {
		str.WriteString(paddingWidthStr)
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
			// Block
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

/* ********************** */
/* Half Block - TrueColor */
/* ********************** */

func NewRendererHalfBlockTrueColor(top bool) *RendererHalfBlockTrueColor {
	return &RendererHalfBlockTrueColor{
		RendererHalfBlock: &RendererHalfBlock{
			Top: top,
		},
	}
}

type RendererHalfBlockTrueColor struct{ *RendererHalfBlock }

func (r *RendererHalfBlockTrueColor) String() string {
	return fmt.Sprintf("%s - %s",
		r.RendererHalfBlock,
		colorprofile.TrueColor,
	)
}

func (r *RendererHalfBlockTrueColor) Support(profile colorprofile.Profile) bool {
	return profile == colorprofile.TrueColor
}

/* ******************** */
/* Half Block - ANSI256 */
/* ******************** */

func NewRendererHalfBlockANSI256(top bool, binding ColorBinding) *RendererHalfBlockANSI256 {
	return &RendererHalfBlockANSI256{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    top,
			Binder: NewColorBinder(binding),
		},
	}
}

type RendererHalfBlockANSI256 struct{ *RendererHalfBlock }

func (r *RendererHalfBlockANSI256) String() string {
	return fmt.Sprintf("%s - %s",
		r.RendererHalfBlock,
		r.RendererHalfBlock.Binder.Binding,
	)
}

func (r *RendererHalfBlockANSI256) Support(profile colorprofile.Profile) bool {
	return profile <= colorprofile.ANSI256
}

/* ***************** */
/* Half Block - ANSI */
/* ***************** */

func NewRendererHalfBlockANSI(top bool, binding ColorBinding) *RendererHalfBlockANSI {
	return &RendererHalfBlockANSI{
		RendererHalfBlock: &RendererHalfBlock{
			Top:    top,
			Binder: NewColorBinder(binding),
		},
	}
}

type RendererHalfBlockANSI struct{ *RendererHalfBlock }

func (r *RendererHalfBlockANSI) String() string {
	return fmt.Sprintf("%s - %s",
		r.RendererHalfBlock,
		r.RendererHalfBlock.Binder.Binding,
	)
}

func (r *RendererHalfBlockANSI) Support(profile colorprofile.Profile) bool {
	return profile <= colorprofile.ANSI
}

/* ******************* */
/* Mixed Block - Ascii */
/* ******************* */

func NewRendererMixedBlockAscii(foreground, background color.Color) *RendererMixedBlockAscii {
	return &RendererMixedBlockAscii{
		foreground: foreground,
		background: background,
		cache:      map[color.Color]color.Color{},
	}
}

type RendererMixedBlockAscii struct {
	foreground color.Color
	background color.Color
	cache      map[color.Color]color.Color
}

func (r *RendererMixedBlockAscii) String() string {
	return "Mixed Block - " + colorprofile.Ascii.String()
}

func (r *RendererMixedBlockAscii) Support(profile colorprofile.Profile) bool {
	return profile <= colorprofile.Ascii
}

func (r *RendererMixedBlockAscii) Ratio() Size { return Size{Width: 1, Height: 2} }

func (r *RendererMixedBlockAscii) SetForeground(c color.Color) {
	if c != r.foreground {
		r.foreground = c
		// Clear cache
		r.cache = map[color.Color]color.Color{}
	}
}

func (r *RendererMixedBlockAscii) SetBackground(c color.Color) {
	if c != r.background {
		r.background = c
		// Clear cache
		r.cache = map[color.Color]color.Color{}
	}
}

func (r *RendererMixedBlockAscii) bind(in color.Color) (out color.Color) {
	var ok bool
	if out, ok = r.cache[in]; !ok {
		distForeground := ciede2000.Diff(in, r.foreground)
		distBackground := ciede2000.Diff(in, r.background)
		if distForeground <= distBackground {
			out = r.foreground
		} else {
			out = r.background
		}
		r.cache[in] = out
	}

	return out
}

func (r *RendererMixedBlockAscii) Render(img *Image, padding Size) string {
	// Padding
	ratio := r.Ratio()
	paddingWidthStr := strings.Repeat(" ", padding.Width/ratio.Width)
	paddingHeightStr := strings.Repeat("\n", padding.Height/ratio.Height)

	// Image bounds
	bounds := img.Bounds()

	// String
	str := strings.Builder{}
	str.WriteString(paddingHeightStr)

	for y := 0; y < bounds.Max.Y; y += 2 {
		str.WriteString(paddingWidthStr)
		for x := 0; x < bounds.Max.X; x++ {
			// Color top
			ct := r.bind(img.NRGBAAt(x, y))
			//  Color bottom
			cb := r.bind(img.NRGBAAt(x, y+1))
			// Block
			if ct == cb {
				if ct == r.foreground {
					str.WriteString("█")
				} else {
					str.WriteString(" ")
				}
			} else {
				if ct == r.foreground {
					str.WriteString("▀")
				} else {
					str.WriteString("▄")
				}
			}
		}
		str.WriteString(ansi.ResetStyle + "\n")
	}

	return str.String()
}

/* ***** */
/* Ascii */
/* ***** */

func NewRendererAscii(density string) *RendererAscii {
	return &RendererAscii{
		density: density,
	}
}

type RendererAscii struct {
	density string
}

func (r *RendererAscii) String() string {
	return "Ascii"
}

func (r *RendererAscii) Support(profile colorprofile.Profile) bool {
	return profile <= colorprofile.Ascii
}

func (r *RendererAscii) Ratio() Size { return Size{Width: 1, Height: 1} }

func (r *RendererAscii) character(c color.NRGBA) string {

	//avgF := float64(int(r)+int(g)+int(b)) / 3.0
	//avg := uint8(math.Round(avgF))

	avg := uint8(math.Round(float64(c.R+c.G+c.B) / 3))

	//len := len(cd)

	l := len(r.density)

	//i := int(mapValue(avg, 0, 255, 0, uint8(len)))

	finalValue := avg
	if avg > 255 {
		finalValue = 255
	} else if avg < 0 {
		finalValue = 0
	}

	aa := float64(255-finalValue) / float64(255-0)
	bb := aa * float64(l-0)
	cc := math.Round(bb)

	if cc < 0 {
		cc = 0
	} else if cc > 255 {
		cc = 255
	}

	i := l - int(cc)

	// ---
	/*
			func mapValue(
				value uint8,
				minIn uint8,
				maxIn uint8,
				minOut uint8,
				maxOut uint8,
		) uint8 {
				finalValue := value

				if value > maxIn {
				finalValue = maxIn
			} else if value < minIn {
				finalValue = minIn
			}

				a := float64(maxIn-finalValue) / float64(maxIn-minIn)

				b := a * float64(maxOut-minOut)

				c := math.Round(b)

				if c < 0 {
				c = 0
			} else if c > 255 {
				c = 255
			}

				return maxOut - uint8(c)
			}
	*/
	// ---

	/*
		if i >= len {
			i = len - 1
		}
	*/

	if i >= l {
		i = l - 1
	}

	//return string(cd[int(i)])
	return string(r.density[i])
}

func (r *RendererAscii) Render(img *Image, padding Size) string {
	// Padding
	ratio := r.Ratio()
	paddingWidthStr := strings.Repeat(" ", padding.Width/ratio.Width)
	paddingHeightStr := strings.Repeat("\n", padding.Height/ratio.Height)

	// Image bounds
	bounds := img.Bounds()

	// String
	str := strings.Builder{}
	str.WriteString(paddingHeightStr)

	for y := 0; y < bounds.Max.Y; y++ {
		str.WriteString(paddingWidthStr)
		for x := 0; x < bounds.Max.X; x++ {

			/*
				co := img.At(x, y)
				rr, gg, bb, _ := co.RGBA()
				r := uint8(rr)
				g := uint8(gg)
				b := uint8(bb)
				hex := rgbToHex(r, g, b)
			*/

			c := img.NRGBAAt(x, y)

			//c := characterFromRgb(r, g, b, config.CharacterDensity)

			char := r.character(c)

			/*
				if hex == "#000000" {
					s := style.
						Foreground(lipgloss.Color("#FFFFFF"))
					if config.SetRandomBlank {
						res += s.Render(getRandomToken(config.CharacterDensity))
					} else {
						res += s.Render(string(config.CharacterDensity[0]))
					}
				} else {
					complementaryHex := rgbToHex(255-r, 255-g, 255-b)
					s := style.
						Background(lipgloss.Color(hex)).
						Foreground(lipgloss.Color(complementaryHex))
					res += s.Render(c)
				}
			*/

			var t ansi.Style
			t = t.BackgroundColor(c)
			t = t.ForegroundColor(color.NRGBA{R: 255 - c.R, G: 255 - c.G, B: 255 - c.B})

			str.WriteString(t.String() + char + ansi.ResetStyle)

		}
		//str.WriteString(ansi.ResetStyle + "\n")
		str.WriteString("\n")
	}

	return str.String()
}
