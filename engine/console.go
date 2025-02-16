package engine

import (
	"duck-hunt-go/engine/log"
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"image"
	"slices"
	"strings"
	"time"
)

const consoleDuration = time.Second * 3

func NewConsole() *Console {
	return &Console{}
}

type Console struct {
	AbsolutePosition
	entries []ConsoleEntry
}

func (c *Console) Init() tea.Cmd {
	return log.Info("Console initialized")
}

func (c *Console) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case ConsoleLogMsg:
		// Stack entries
		c.entries = append(c.entries, ConsoleEntry{
			Text:       msg.Text,
			Expiration: time.Now().Add(consoleDuration),
		})
	case TickMsg:
		// Purge old entries
		c.entries = slices.DeleteFunc(c.entries, func(entry ConsoleEntry) bool {
			return entry.Expiration.Before(time.Now())
		})
	}
	return nil
}

func (c *Console) Image() *Image {
	// No entries
	if len(c.entries) == 0 {
		return nil
	}

	// Compute text
	var text strings.Builder
	for i, entry := range c.entries {
		if i != 0 {
			text.WriteString("\n")
		}
		text.WriteString(entry.Text)
	}

	return NewText5x5(text.String(), ColorWhite).Image()
}

func (c *Console) Draw(img *Image) {
	src := c.Image()
	if src != nil {
		position := c.Position()
		img.DrawImage(image.Pt(
			int(position.X),
			int(position.Y),
		), src)
	}
}

type ConsoleEntry struct {
	Text       string
	Expiration time.Time
}

func ConsoleLog(format string, a ...any) tea.Cmd {
	return func() tea.Msg {
		return ConsoleLogMsg{
			Text: fmt.Sprintf(format, a...),
		}
	}
}

type ConsoleLogMsg struct {
	Text string
}
