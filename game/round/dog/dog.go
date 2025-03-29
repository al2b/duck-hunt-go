package dog

import (
	"duck-hunt-go/engine"
	"embed"
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

const frameDuration = time.Second / 60

//go:embed assets/*.png
var assets embed.FS

var (
	imageTrack1    = engine.Must(engine.LoadImage(assets, "assets/dog.track.1.png"))
	imageTrack2    = engine.Must(engine.LoadImage(assets, "assets/dog.track.2.png"))
	imageTrack3    = engine.Must(engine.LoadImage(assets, "assets/dog.track.3.png"))
	imageTrack4    = engine.Must(engine.LoadImage(assets, "assets/dog.track.4.png"))
	imageSniff     = engine.Must(engine.LoadImage(assets, "assets/dog.sniff.png"))
	imagePant      = engine.Must(engine.LoadImage(assets, "assets/dog.pant.png"))
	animationTrack = engine.SequenceAnimation{
		engine.RepeatAnimation{
			engine.SequenceAnimation{
				engine.RepeatAnimation{
					engine.Animation{
						{imageTrack1, frameDuration * 7},
						{imageTrack2, frameDuration * 7},
						{imageTrack3, frameDuration * 7},
						{imageTrack4, frameDuration * 7},
					},
					4,
				},
				engine.Animation{
					{imageTrack1, frameDuration * 14},
				},
				engine.RepeatAnimation{
					engine.Animation{
						{imageSniff, frameDuration * 9},
						{imageTrack1, frameDuration * 9},
					},
					2,
				},
				engine.Animation{
					{imageSniff, frameDuration * 10},
				},
			},
			2,
		},
		engine.Animation{
			{imageTrack1, frameDuration * 7},
			{imageTrack2, frameDuration * 1},
			{imagePant, frameDuration * 18},
		},
	}
	pathTrack = engine.ChainPath{
		engine.FixedPath{Position: engine.Vec(2, 141)},
		engine.StepPath{Delta: engine.Vec(2, 0), Span: frameDuration * 7, Count: 4 * 4},
		engine.FixedPath{Span: 60 * frameDuration},
		engine.StepPath{Delta: engine.Vec(2, 0), Span: frameDuration * 7, Count: 4 * 4},
		engine.FixedPath{Span: 60 * frameDuration},
		engine.FixedPath{Span: 26 * frameDuration},
	}
)

func New() *Dog {
	return &Dog{}
}

type Dog struct {
	animationTrack engine.AnimationPlayer
	pathTrack      engine.PathPlayer
}

func (m *Dog) Init() tea.Cmd {
	m.animationTrack = engine.AnimationPlayer{Animation: animationTrack, Loop: true}
	m.pathTrack = engine.PathPlayer{Path: pathTrack, Loop: true}
	return nil
}

func (m *Dog) Update(msg tea.Msg) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case engine.TickMsg:
		m.animationTrack.Step(msg.Duration)
		m.pathTrack.Step(msg.Duration)
	}
	return nil
}

func (m *Dog) Draw(scene *engine.Image) {
	scene.Draw(
		engine.DrawImage(m.pathTrack.Position().Point(), m.animationTrack.Image()),
	)
}
