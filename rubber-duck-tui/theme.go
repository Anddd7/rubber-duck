package tui

import (
	"github.com/charmbracelet/lipgloss"
)

type Default struct {
}

func (t Default) Title(s string) string {
	return lipgloss.NewStyle().Margin(1, 2).Render(s)
}
