package mouse

import (
	"duck-hunt-go/engine"
	"embed"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const TickInterval = time.Second / 60

var (
	//go:embed assets/*.png
	assets embed.FS

	// Images
	imageMouse         = engine.MustLoad(engine.ImageLoader{assets, "assets/mouse.png"})
	imageMouseOutWhite = engine.ImageSlicer{imageMouse, engine.Point{0, 0}, engine.Size{11, 16}}.Image()
	imageMouseOutRed   = engine.ImageSlicer{imageMouse, engine.Point{0, 16}, engine.Size{11, 16}}.Image()
	imageMouseOutGreen = engine.ImageSlicer{imageMouse, engine.Point{0, 32}, engine.Size{11, 16}}.Image()
	imageMouseOutBlue  = engine.ImageSlicer{imageMouse, engine.Point{0, 48}, engine.Size{11, 16}}.Image()
	imageMouseInWhite  = engine.ImageSlicer{imageMouse, engine.Point{0, 64}, engine.Size{11, 16}}.Image()
	imageMouseInRed    = engine.ImageSlicer{imageMouse, engine.Point{0, 80}, engine.Size{11, 16}}.Image()
	imageMouseInGreen  = engine.ImageSlicer{imageMouse, engine.Point{0, 96}, engine.Size{11, 16}}.Image()
	imageMouseInBlue   = engine.ImageSlicer{imageMouse, engine.Point{0, 112}, engine.Size{11, 16}}.Image()
)

func New() *Mouse {
	return &Mouse{
		timerLeft:  engine.TimerPlayer{Timer: engine.Timer{Span: time.Second * 1}, OnEnd: engine.PlayerOnEndStopRewind},
		timerRight: engine.TimerPlayer{Timer: engine.Timer{Span: time.Second * 1}, OnEnd: engine.PlayerOnEndStopRewind},
	}
}

type Mouse struct {
	point                   engine.Point
	buttonLeft, buttonRight bool
	timerLeft, timerRight   engine.TimerPlayer
}

func (s *Mouse) String() string {
	return "Mouse"
}

func (s *Mouse) Size(_ engine.Size) engine.Size {
	return engine.Size{80, 50}
}

func (s *Mouse) Init() (cmd tea.Cmd) {
	s.point = engine.Point{0, 0}

	return tea.Batch(
		engine.StartTicker(TickInterval),
	)
}

func (s *Mouse) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		mouse := msg.Mouse()
		s.point = engine.Point{mouse.X, mouse.Y}
		switch msg := msg.(type) {
		case tea.MouseClickMsg:
			switch msg.Button {
			case tea.MouseLeft:
				s.buttonLeft = true
				s.timerLeft.Play()
			case tea.MouseRight:
				s.buttonRight = true
				s.timerRight.Play()
			}
		case tea.MouseReleaseMsg:
			switch msg.Button {
			case tea.MouseLeft:
				s.buttonLeft = false
			case tea.MouseRight:
				s.buttonRight = false
			}
		}
	case engine.TickMsg:
		s.timerLeft.Step(msg.Interval)
		s.timerRight.Step(msg.Interval)
	}

	return nil
}

func (s *Mouse) Draw(dst *engine.Image) {

	// Out
	imageMouseOut := imageMouseOutWhite
	switch true {
	case s.buttonLeft && s.buttonRight:
		imageMouseOut = imageMouseOutBlue
	case s.buttonLeft:
		imageMouseOut = imageMouseOutRed
	case s.buttonRight:
		imageMouseOut = imageMouseOutGreen
	}

	// In
	imageMouseIn := imageMouseInWhite
	switch true {
	case s.timerLeft.Playing() && s.timerRight.Playing():
		imageMouseIn = imageMouseInBlue
	case s.timerLeft.Playing():
		imageMouseIn = imageMouseInRed
	case s.timerRight.Playing():
		imageMouseIn = imageMouseInGreen
	}

	dst.Draw(
		engine.ImageDrawer{s.point, imageMouseOut},
		engine.ImageDrawer{s.point, imageMouseIn},
		engine.ImageDrawer{s.point.Add(engine.Point{0, 17}), engine.Text{
			fmt.Sprintf("%d,%d", s.point.X, s.point.Y),
			engine.Font5x5, engine.ColorWhite,
		}},
	)
}
