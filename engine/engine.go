package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"slices"
)

var (
	pause = false
)

func New(model Model, width, height int, fps int) Engine {
	return Engine{
		width:       width,
		height:      height,
		fps:         fps,
		model:       model,
		intersector: NewIntersector(),
		renderer:    NewRenderer(width, height),
	}
}

type Engine struct {
	width, height int
	fps           int
	// Window
	windowWidth, windowHeight int
	// Model
	model Model
	// Messages
	msgs []tea.Msg
	// Intersections
	intersector *Intersector
	// View
	renderer *Renderer
	view     string
}

func (e Engine) Init() (tea.Model, tea.Cmd) {
	return e, tea.Batch(
		// Force requesting window size again for certain terminal who
		// don't respond in time to the first automatic bubble tea request
		tea.RequestWindowSize(),
		// According to documentation, these should be enabled as a program option
		tea.EnterAltScreen,
		tea.EnableMouseAllMotion,
		e.model.Init(),
		tick(e.fps),
	)
}

func (e Engine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		e.windowWidth, e.windowHeight = msg.Width*e.renderer.WidthRatio(), msg.Height*e.renderer.HeightRatio()
		return e, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		// Quit
		case "enter", "q", "ctrl+c", "esc":
			return e, tea.Quit
		// Mode
		case "m":
			switch mode {
			case Mode8:
				mode = Mode24
			case Mode24:
				mode = Mode8
			}
			return e, nil
		// Debug
		case "d":
			debug = !debug
			return e, nil
		// Pause
		case "p":
			pause = !pause
			return e, nil
		}
	case TickMsg:
		cmds = append(cmds, tick(e.fps))
	}

	if pause {
		return e, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		e.msgs = append(e.msgs, msg)
	case tea.MouseMsg:
		mouse := msg.Mouse()
		// Mouse position
		width, height, paddingHorizontal, paddingVertical := e.size()
		x := (((msg.Mouse().X * e.renderer.WidthRatio()) - paddingHorizontal) * e.width) / width
		y := (((msg.Mouse().Y * e.renderer.HeightRatio()) - paddingVertical) * e.height) / height
		if mouse.Button != tea.MouseNone {
			e.msgs = append(e.msgs, tea.MouseClickMsg{X: x, Y: y, Button: mouse.Button})
		} else {
			// Remove previous mouse motion messages...
			e.msgs = slices.DeleteFunc(e.msgs, func(msg tea.Msg) bool {
				_, ok := msg.(tea.MouseMotionMsg)
				return ok
			})
			// ...to keep only the last one
			e.msgs = append(e.msgs, tea.MouseMotionMsg{X: x, Y: y})
		}
	case TickMsg:
		msgs := append(e.msgs, msg)
		for _, msg := range msgs {
			cmds = append(cmds, e.model.Update(msg))
		}
		cmds = append(cmds, func() tea.Msg {
			return ModelUpdatedMsg{}
		})
		e.msgs = nil
	case ModelUpdatedMsg:
		cmds = append(cmds, e.model.Update(msg))
		cmds = append(cmds, func() tea.Msg {
			e.intersector.Intersect(e.model)
			return ModelIntersectedMsg{}
		})
	case ModelIntersectedMsg:
		width, height, paddingHorizontal, paddingVertical := e.size()
		e.view = e.renderer.Render(e.model.Sprites(), width, height, paddingHorizontal, paddingVertical)
	}

	return e, tea.Batch(cmds...)
}

func (e Engine) size() (width, height int, paddingHorizontal, paddingVertical int) {
	// Fit in window with optional padding
	if (e.windowWidth >= e.width) && (e.windowHeight >= e.height) {
		width, height = e.width, e.height
	} else {
		widthRatio := float64(e.windowWidth) / float64(e.width)
		heightRatio := float64(e.windowHeight) / float64(e.height)

		ratio := widthRatio
		if heightRatio < widthRatio {
			ratio = heightRatio
		}

		width = int(float64(e.width) * ratio)
		height = int(float64(e.height) * ratio)
	}

	return width, height,
		(e.windowWidth - width) / 2,
		(e.windowHeight - height) / 2
}

func (e Engine) View() string {
	return e.view
}
