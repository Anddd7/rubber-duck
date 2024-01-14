package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectionModel struct {
	list list.Model
}

func (m SelectionModel) Init() {
}

func (m SelectionModel) Update(msg tea.Msg) (SelectionModel, tea.Cmd) {
	return m, nil
}

func (m SelectionModel) View() string {
	return ""
}
