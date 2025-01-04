package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"slices"
)

const (
	Width  = 256
	Height = 240
	Fps    = 60
)

var (
	zoom  = 1
	pause = false
)

func New(model Model) Engine {
	return Engine{
		model:       model,
		intersector: NewIntersector(),
		renderer:    NewRenderer(),
	}
}

type Engine struct {
	// Window
	width, height int
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

func (m Engine) Init() (tea.Model, tea.Cmd) {
	m.model.Init()
	return m, tick()
}

func (m Engine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		// Quit
		case "enter", "q", "ctrl+c", "esc":
			return m, tea.Quit
		// Mode
		case "m":
			switch mode {
			case Mode8:
				mode = Mode24
			case Mode24:
				mode = Mode8
			}
			return m, nil
		// Debug
		case "d":
			debug = !debug
			return m, nil
		// Pause
		case "p":
			pause = !pause
			return m, nil
		// Zoom in
		case "i":
			if zoom > 1 {
				zoom = zoom / 2
			}
			return m, nil
		// Zoom out
		case "o":
			if zoom < 8 {
				zoom = zoom * 2
			}
			return m, nil
		}
	case TickMsg:
		cmds = append(cmds, tick())
	}

	if pause {
		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		m.msgs = append(m.msgs, msg)
	case tea.MouseMsg:
		// Remove previous mouse motion messages...
		m.msgs = slices.DeleteFunc(m.msgs, func(msg tea.Msg) bool {
			_, ok := msg.(tea.MouseMotionMsg)
			return ok
		})
		// ...to keep only the last one
		top, left := m.padding()
		m.msgs = append(m.msgs, tea.MouseMotionMsg{
			X: (msg.Mouse().X - left) * zoom,
			Y: ((msg.Mouse().Y - top) * 2) * zoom,
		})
	case TickMsg:
		msgs := append(m.msgs, msg)
		cmds = append(cmds, func() tea.Msg {
			m.model.Update(msgs)
			return ModelUpdatedMsg{}
		})
		m.msgs = nil
	case ModelUpdatedMsg:
		cmds = append(cmds, func() tea.Msg {
			m.intersector.Intersect(m.model)
			return ModelIntersectedMsg{}
		})
	case ModelIntersectedMsg:
		// Render
		top, left := m.padding()
		m.view = m.renderer.Render(m.model, top, left)
	}

	return m, tea.Batch(cmds...)
}

func (m Engine) padding() (top int, left int) {
	if m.width > (Width / zoom) {
		left = (m.width - (Width / zoom)) / 2
	}
	if m.height > ((Height / zoom) / 2) {
		top = (m.height - ((Height / zoom) / 2)) / 2
	}
	return
}

func (m Engine) View() string {
	return m.view
}
