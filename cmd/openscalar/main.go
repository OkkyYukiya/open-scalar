package main

import (
	"fmt"
	"os"

	"github.com/OkkyYukiya/openscalar/internal/agent/codex"
	"github.com/OkkyYukiya/openscalar/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	sessions := []codex.Session{
		{ThreadName: "refactor auth", SizeBytes: 29_000_000, Turns: 120, TotalTokens: 210_000, ContextWindow: 258_400, CompactCount: 3},
		{ThreadName: "fix tui bug", SizeBytes: 1_200_000, Turns: 8, TotalTokens: 40_000, ContextWindow: 258_400, CompactCount: 0},
		{ThreadName: "write tests", SizeBytes: 5_400_000, Turns: 32, TotalTokens: 130_000, ContextWindow: 258_400, CompactCount: 1},
	}

	p := tea.NewProgram(tui.New(sessions), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "run tui:", err)
		os.Exit(1)
	}

}