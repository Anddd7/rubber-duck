package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle()
	itemStyle         = lipgloss.NewStyle().PaddingLeft(2)
	focusItemStyle    = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("120")).Bold(true)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("120"))
	helpStyle         = lipgloss.NewStyle().PaddingTop(1).Foreground(lipgloss.Color("240"))
)

type item[T any] struct {
	label string
	value T
}

type SelectionModel[T any] struct {
	title      string
	items      []item[T]
	selected   map[int]struct{}
	cursor     int
	size       int
	singleMode bool
	quitting   bool
	autoclear  bool
}

func (m *SelectionModel[T]) Init() tea.Cmd {
	return nil
}

func (m *SelectionModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		// quit
		case "ctrl+c", "q":
			m.selected = map[int]struct{}{}
			m.quitting = true
			return m, tea.Quit
		case "c":
			m.quitting = true
			return m, tea.Quit

		// cursor movement
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = m.size - 1
			}
		case "down", "j":
			if m.cursor < m.size {
				m.cursor++
			} else {
				m.cursor = 0
			}

		// select
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				if m.singleMode {
					m.selected = map[int]struct{}{}
				}
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m *SelectionModel[T]) View() string {
	if m.quitting && m.autoclear {
		return ""
	}

	var output []string

	output = append(output, titleStyle.Render(m.title))

	for i, item := range m.items {
		style := itemStyle
		cursor := " "
		checked := " "

		if _, ok := m.selected[i]; ok {
			checked = "x"
			style = selectedItemStyle
		}
		if m.cursor == i {
			cursor = ">"
			style = focusItemStyle
		}

		output = append(output, style.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, item.label)))
	}

	output = append(output, helpStyle.Render(":Press q to quit, c to confirm."))

	return strings.Join(output, "\n")
}

func NewSelection[T any](title string, options []T, getlabel func(T) string) *SelectionModel[T] {
	_items := []item[T]{}
	for _, option := range options {
		_items = append(_items, item[T]{label: getlabel(option), value: option})
	}

	return &SelectionModel[T]{
		title:    title,
		items:    _items,
		selected: map[int]struct{}{},
		size:     len(options),
	}
}

func (m *SelectionModel[T]) Start() error {
	if _, err := tea.NewProgram(m).Run(); err != nil {
		return err
	}
	return nil
}

func (m *SelectionModel[T]) GetSelected() []T {
	selected := []T{}
	for i := range m.selected {
		selected = append(selected, m.items[i].value)
	}
	return selected
}

func (m *SelectionModel[T]) SetSingleMode() {
	m.singleMode = true
}

func (m *SelectionModel[T]) SetAutoclear() {
	m.autoclear = true
}

func NewStringSelection(title string, options []string) *SelectionModel[string] {
	return NewSelection(title, options, func(s string) string { return s })
}

// demo
func main() {
	// options := []string{"foo", "bar", "baz"}
	// model := NewStringSelection("Select:", options)
	// model.Start()
	// fmt.Println(model.GetSelected())

	type option struct {
		id   int
		name string
	}
	options := []option{
		{id: 10086, name: "foo"},
		{id: 20086, name: "bar"},
		{id: 30086, name: "baz"},
	}
	model := NewSelection("Select:", options, func(o option) string { return o.name })
	model.SetSingleMode()
	model.SetAutoclear()
	model.Start()
	fmt.Println(model.GetSelected())
}
