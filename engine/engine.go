package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"slices"
)

var (
	pause = false
)

func New(scene Scene) Engine {
	return Engine{
		scene:       scene,
		intersector: NewIntersector(),
		renderer:    &RendererHalfBlockBottom24{},
	}
}

type Engine struct {
	scene Scene
	// Window
	windowWidth, windowHeight int
	// Messages
	msgs []tea.Msg
	// Intersections
	intersector *Intersector
	// View
	renderer Renderer
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
		e.scene.Init(),
		tick(e.scene.FPS()),
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
			switch e.renderer.(type) {
			case *RendererHalfBlockBottom8:
				e.renderer = &RendererHalfBlockBottom24{}
			case *RendererHalfBlockBottom24:
				e.renderer = &RendererHalfBlockBottom8{}
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
		cmds = append(cmds, tick(e.scene.FPS()))
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
		x := (((msg.Mouse().X * e.renderer.WidthRatio()) - paddingHorizontal) * e.scene.Width()) / width
		y := (((msg.Mouse().Y * e.renderer.HeightRatio()) - paddingVertical) * e.scene.Height()) / height
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
			cmds = append(cmds, e.scene.Update(msg))
		}
		cmds = append(cmds, func() tea.Msg {
			return ModelUpdatedMsg{}
		})
		e.msgs = nil
	case ModelUpdatedMsg:
		cmds = append(cmds, e.scene.Update(msg))
		cmds = append(cmds, func() tea.Msg {
			e.intersector.Intersect(e.scene)
			return ModelIntersectedMsg{}
		})
	case ModelIntersectedMsg:
		resizeWidth, resizeHeight, paddingHorizontal, paddingVertical := e.size()
		e.view = e.renderer.Render(e.scene.Sprites(), e.scene.Width(), e.scene.Height(), resizeWidth, resizeHeight, paddingHorizontal, paddingVertical)
	}

	return e, tea.Batch(cmds...)
}

func (e Engine) size() (width, height int, paddingHorizontal, paddingVertical int) {
	sceneWidth, sceneHeight := e.scene.Width(), e.scene.Height()
	// Fit in window with optional padding
	if (e.windowWidth >= sceneWidth) && (e.windowHeight >= sceneHeight) {
		width, height = sceneWidth, sceneHeight
	} else {
		widthRatio := float64(e.windowWidth) / float64(sceneWidth)
		heightRatio := float64(e.windowHeight) / float64(sceneHeight)

		ratio := widthRatio
		if heightRatio < widthRatio {
			ratio = heightRatio
		}

		width = int(float64(sceneWidth) * ratio)
		height = int(float64(sceneHeight) * ratio)
	}

	return width, height,
		(e.windowWidth - width) / 2,
		(e.windowHeight - height) / 2
}

func (e Engine) View() string {
	return e.view
}
