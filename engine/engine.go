package engine

import (
	"context"
	tea "github.com/charmbracelet/bubbletea/v2"
	"log/slog"
)

func New(scene Scene, options ...Option) Engine {
	engine := Engine{
		scene:          scene,
		console:        NewConsole(),
		consoleEnabled: false,
		ticker:         Ticker{},
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
	// Ticker
	ticker tea.Model
	// View
	renderers *Renderers
	view      string
	// Log
	logHandler slog.Handler
}

func (e Engine) Init() tea.Cmd {
	return tea.Batch(
		LogInfo("Initialize engine..."),
		// Force requesting window size again for certain terminal who
		// don't respond in time to the first automatic bubble tea request
		tea.RequestWindowSize,
		tea.EnterAltScreen,
		tea.EnableMouseAllMotion,
		e.renderers.Init(),
		ConsoleLog("Renderer: %s", e.renderers.Current()),
		e.console.Init(),
		e.scene.Init(),
		e.ticker.Init(),
	)
}

func (e Engine) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	// Ticker
	e.ticker, cmd = e.ticker.Update(msg)
	cmds = append(cmds, cmd)

	// Renderers
	cmds = append(cmds, e.renderers.Update(msg))

	// Console
	cmds = append(cmds, e.console.Update(msg))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ratio := e.renderers.Current().Ratio()
		e.windowSize = Size{msg.Width * ratio.Width, msg.Height * ratio.Height}
		return e, ConsoleLog("Window size: %s", e.windowSize)
	case tea.KeyPressMsg:
		switch key := msg.Key(); key.Code {
		case tea.KeyF1:
			e.consoleEnabled = !e.consoleEnabled
			return e, nil
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
		}
	case LogMsg:
		_ = e.logHandler.Handle(context.Background(), msg.Record)
	}

	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Cannot determine the mouse position if the window size is zero
		if e.windowSize.IsZero() {
			break
		}

		sceneSize := e.scene.Size(e.windowSize)
		screenSize, screenPadding := e.screenDimensions(e.windowSize, sceneSize)

		// Mouse position
		mouse := msg.Mouse()
		ratio := e.renderers.Current().Ratio()

		x := (((mouse.X * ratio.Width) - screenPadding.Width) * sceneSize.Width) / screenSize.Width
		y := (((mouse.Y * ratio.Height) - screenPadding.Height) * sceneSize.Height) / screenSize.Height

		switch msg := msg.(type) {
		case tea.MouseClickMsg:
			if msg.Button == tea.MouseNone {
				cmds = append(cmds, e.scene.Update(
					tea.MouseMotionMsg{X: x, Y: y, Button: msg.Button, Mod: msg.Mod},
				))
			} else {
				msg.X, msg.Y = x, y
				cmds = append(cmds, e.scene.Update(msg))
			}
		case tea.MouseReleaseMsg:
			msg.X, msg.Y = x, y
			cmds = append(cmds, e.scene.Update(msg))
		case tea.MouseWheelMsg:
			msg.X, msg.Y = x, y
			cmds = append(cmds, e.scene.Update(msg))
		case tea.MouseMotionMsg:
			msg.X, msg.Y = x, y
			cmds = append(cmds, e.scene.Update(msg))
		}
	case TickMsg:
		cmds = append(cmds, e.scene.Update(msg))

		// Cannot draw the scene if the window size is zero
		if e.windowSize.IsZero() {
			break
		}

		sceneSize := e.scene.Size(e.windowSize)
		screenSize, screenPadding := e.screenDimensions(e.windowSize, sceneSize)

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

func (e Engine) screenDimensions(windowSize, sceneSize Size) (size, padding Size) {
	// Fit the scene in the window with optional padding
	if (windowSize.Width >= sceneSize.Width) && (windowSize.Height >= sceneSize.Height) {
		size = sceneSize
	} else {
		ratioWidth := float64(windowSize.Width) / float64(sceneSize.Width)
		ratioHeight := float64(windowSize.Height) / float64(sceneSize.Height)

		ratio := ratioWidth
		if ratioHeight < ratioWidth {
			ratio = ratioHeight
		}

		size.Width = int(float64(sceneSize.Width) * ratio)
		size.Height = int(float64(sceneSize.Height) * ratio)
	}

	padding.Width = (windowSize.Width - size.Width) / 2
	padding.Height = (windowSize.Height - size.Height) / 2

	return
}

func (e Engine) View() string {
	return e.view
}

type Option func(engine *Engine)

func WithLogHandler(handler slog.Handler) Option {
	return func(engine *Engine) {
		engine.logHandler = handler
	}
}
