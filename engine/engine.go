package engine

import (
	"context"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"log/slog"
	"slices"
)

//go:embed assets/*
var assets embed.FS

var (
	pause = false
)

func New(scene Scene, options ...Option) Engine {
	engine := Engine{
		scene:          scene,
		console:        NewConsole(),
		consoleEnabled: false,
		renderers:      NewRenderers(),
		logHandler:     slog.DiscardHandler,
	}

	// Options
	for _, option := range options {
		option(&engine)
	}

	return engine
}

type Engine struct {
	scene Scene
	// Console
	console        *Console
	consoleEnabled bool
	// Window
	windowSize Size
	// Messages
	msgs []tea.Msg
	// View
	renderers *Renderers
	view      string
	// Log
	logHandler slog.Handler
}

func (e Engine) Init() (tea.Model, tea.Cmd) {
	return e, tea.Batch(
		LogInfo("Initialize engine..."),
		// Force requesting window size again for certain terminal who
		// don't respond in time to the first automatic bubble tea request
		tea.RequestWindowSize(),
		tea.EnterAltScreen,
		tea.EnableMouseAllMotion,
		e.renderers.Init(),
		ConsoleLog("Renderer: %s", e.renderers.Current()),
		e.console.Init(),
		e.scene.Init(),
		Tick(e.scene.FPS()),
	)
}

func (e Engine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Renderers
	cmds = append(cmds, e.renderers.Update(msg))

	// Console
	cmds = append(cmds, e.console.Update(msg))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ratio := e.renderers.Current().Ratio()
		e.windowSize = Size{
			Width:  msg.Width * ratio.Width,
			Height: msg.Height * ratio.Height,
		}
		return e, ConsoleLog("Window size: %s", e.windowSize)
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyF1:
			e.consoleEnabled = !e.consoleEnabled
		}
		switch msg.String() {
		// Quit
		case "enter", "q", "ctrl+c", "esc":
			return e, tea.Quit
		// Mode
		case "l":
			return e, ConsoleLog("Renderer: %s", e.renderers.Previous())
		case "m":
			return e, ConsoleLog("Renderer: %s", e.renderers.Next())
		// Pause
		case "p":
			pause = !pause
			return e, nil
		}
	case TickMsg:
		cmds = append(cmds, Tick(e.scene.FPS()))
	case LogMsg:
		_ = e.logHandler.Handle(context.Background(), msg.Record)
	}

	// Pause
	if pause {
		return e, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		e.msgs = append(e.msgs, msg)
	case tea.MouseMsg:
		mouse := msg.Mouse()
		// Mouse position
		screenSize, screenPadding := e.screenDimensions()
		ratio := e.renderers.Current().Ratio()
		sceneSize := e.scene.Size(e.windowSize)
		x := (((msg.Mouse().X * ratio.Width) - screenPadding.Width) * sceneSize.Width) / screenSize.Width
		y := (((msg.Mouse().Y * ratio.Height) - screenPadding.Height) * sceneSize.Height) / screenSize.Height
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
		cmds = append(cmds, func() tea.Msg {
			return ModelIntersectedMsg{}
		})
	case ModelIntersectedMsg:
		screenSize, screenPadding := e.screenDimensions()

		sceneSize := e.scene.Size(e.windowSize)

		scene := NewImage(sceneSize)

		e.scene.Draw(scene)

		// Console
		if e.consoleEnabled {
			scene.Draw(e.console)
		}

		e.view = e.renderers.Current().Render(
			scene.Resize(screenSize),
			screenPadding,
		)
	default:
		cmds = append(cmds, e.scene.Update(msg))
	}

	return e, tea.Batch(cmds...)
}

func (e Engine) screenDimensions() (size, padding Size) {
	sceneSize := e.scene.Size(e.windowSize)

	// Fit scene in window with optional padding
	if (e.windowSize.Width >= sceneSize.Width) && (e.windowSize.Height >= sceneSize.Height) {
		size = sceneSize
	} else {
		ratioWidth := float64(e.windowSize.Width) / float64(sceneSize.Width)
		ratioHeight := float64(e.windowSize.Height) / float64(sceneSize.Height)

		ratio := ratioWidth
		if ratioHeight < ratioWidth {
			ratio = ratioHeight
		}

		size.Width = int(float64(sceneSize.Width) * ratio)
		size.Height = int(float64(sceneSize.Height) * ratio)
	}

	padding.Width = (e.windowSize.Width - size.Width) / 2
	padding.Height = (e.windowSize.Height - size.Height) / 2

	return
}

func (e Engine) View() string {
	return e.view
}
