package main

import (
	"duck-hunt-go/engine"
	"duck-hunt-go/scene"
	"github.com/alecthomas/kong"
	tea "github.com/charmbracelet/bubbletea/v2"
)

var cli struct {
	Log string `type:"path" short:"l" env:"LOG" help:"Log to file path." placeholder:"PATH"`
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

	program := tea.NewProgram(
		engine.New(
			scene.New(), options...,
		),
	)

	if _, err := program.Run(); err != nil {
		panic(err)
	}
}
