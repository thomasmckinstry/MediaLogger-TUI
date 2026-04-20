package components

import (
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"fmt"
)

type CheckboxModel struct {
	cursor    int
	entries   []string
	entryVals []bool
	title     string
	selected  bool
	width     int

	entryStyle lipgloss.Style
}

func InitialCheckbox(entries []string, title string, width int) CheckboxModel {
	return CheckboxModel{
		cursor:    1,
		entries:   entries,
		width:     width,
		title:     title,
		entryVals: make([]bool, len(entries)),
		entryStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D17600")),
	}
}

func (*CheckboxModel) Init() tea.Cmd {
	return nil
}

func (m *CheckboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc": // Unfocus the component
			m.selected = false
		case "enter": // Add a tag from the current text input and empty the text input OR focus the component
			if !m.selected {
				m.selected = true
			} else {
				m.entryVals[m.cursor] = !m.entryVals[m.cursor]
			}
		case "j", "down": // Nav between tags
			if m.cursor < len(m.entries)-1 && m.selected {
				m.cursor++
			}
			cmd = func() tea.Msg { return NavMsg(!m.selected) }
		case "k", "up":
			if m.cursor > 0 && m.selected {
				m.cursor--
			}
			cmd = func() tea.Msg { return NavMsg(!m.selected) }

		}
	}
	return m, cmd
}

func (m *CheckboxModel) View() tea.View {
	var s string
	s = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, m.title)
	for i, medium := range m.entries {
		var entry string
		check := " "
		if m.entryVals[i] {
			check = "x"
		}
		entry = lipgloss.PlaceHorizontal(m.width-4, lipgloss.Center, medium)
		entry = lipgloss.JoinHorizontal(lipgloss.Center, fmt.Sprintf(" [%s] ", check), entry)
		if i == m.cursor && m.selected {
			entry = m.entryStyle.Render(entry)
		}
		s = lipgloss.JoinVertical(lipgloss.Left, s, entry)
	}
	return tea.NewView(s)
}
