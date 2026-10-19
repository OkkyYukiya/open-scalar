package tui

import (
	"fmt"
	"strconv"

	"github.com/OkkyYukiya/openscalar/internal/agent/codex"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	table table.Model
}

func (m Model) Init () tea.Cmd {
	return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View () string {
	return m.table.View() + "\n  ↑/↓: move  q: quit\n"
}

func toRow(s codex.Session) table.Row {
	return  table.Row{
		s.ThreadName,
		strconv.FormatInt(s.SizeBytes, 10),
		strconv.Itoa(s.Turns),
		strconv.FormatInt(s.TotalTokens, 10),
		fmt.Sprintf("%.0f%%", s.ContextRatio()*100),
		strconv.Itoa(s.CompactCount),
	}
}


var columns = []table.Column{
	{Title: "Thread", Width: 24},
	{Title: "Size", Width: 10},
	{Title: "Turns", Width: 6},
	{Title: "Tokens", Width: 10},
	{Title: "Ctx%", Width: 6},
	{Title: "Compact", Width: 8},
}

func New (sessions []codex.Session) Model {

	rows := make([]table.Row, 0, len(sessions))
	for _, s := range sessions {
		rows = append(rows, toRow(s))
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	return Model{
		table: t,
	}
}