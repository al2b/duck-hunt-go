package main

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/examples"
	"duck-hunt-go/scene"
	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var cli struct {
	Log      string `type:"path" short:"l" env:"LOG" help:"Log to file path." placeholder:"PATH"`
	Examples bool   `short:"e" env:"EXAMPLES" help:"Play examples."`
}

func main() {
	kong.Parse(&cli,
		kong.Name("duck-hunt"),
		kong.Description("A free adaptation of duck hunt in your terminal"),
		kong.UsageOnError(),
	)

	var options []engine.Option

	// Log
	if cli.Log != "" {
		options = append(options, engine.WithLogHandler(
			engine.MustNewLogFileHandler(cli.Log),
		))
	}

	// Scene
	var s engine.Scene
	if cli.Examples {
		s = examples.New()
	} else {
		s = scene.New()
	}

	program := tea.NewProgram(
		engine.New(s, options...),
	)

	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
