package engine

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"time"
)

type TickMsg struct {
	time.Time
	Interval time.Duration
}

type startTickerMsg time.Duration

func StartTicker(interval time.Duration) tea.Cmd {
	return func() tea.Msg {
		return startTickerMsg(interval)
	}
}

type stopTickerMsg struct{}

func StopTicker() tea.Msg {
	return stopTickerMsg{}
}

type stepTickMsg time.Duration

func StepTick(interval time.Duration) tea.Cmd {
	return func() tea.Msg {
		return stepTickMsg(interval)
	}
}

type Ticker struct {
	active   bool
	interval time.Duration
}

func (t Ticker) Init() tea.Cmd {
	return nil
}

func (t Ticker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case startTickerMsg:
		t.interval = time.Duration(msg)
		if !t.active {
			t.active = true
			return t, t.tick(t.interval)
		}
	case stopTickerMsg:
		t.active = false
	case stepTickMsg:
		if !t.active {
			return t, t.tick(time.Duration(msg))
		}
	case TickMsg:
		if t.active {
			return t, t.tick(t.interval)
		}
	}

	return t, nil
}

func (t Ticker) tick(interval time.Duration) tea.Cmd {
	return tea.Tick(
		interval,
		func(time time.Time) tea.Msg {
			return TickMsg{
				Time:     time,
				Interval: interval,
			}
		},
	)
}
