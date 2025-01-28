package engine

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea/v2"
	"math"
	"slices"
	"strings"
	"time"
)

const consoleDuration = time.Second * 3

func NewConsole() *Console {
	return &Console{}
}

type Console struct {
	entries []ConsoleEntry
}

func (c *Console) Init() tea.Cmd {
	return nil
}

func (c *Console) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case ConsoleLogMsg:
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

func (c *Console) Sprites() Sprites {
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

	return Sprites{
		NewCoordinatedSprite(
			NewCoordinates(0, 0, math.MaxFloat64),
			NewText5x5(text.String(), ColorWhite).Image(),
		),
	}
}

func (c *Console) Bodies() Bodies {
	return nil
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
