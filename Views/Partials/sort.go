package partials

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SortModel struct {
	text     string
	selected bool
	style    lipgloss.Style
}

func (m SortModel) toggleBorder() lipgloss.Style {
	if m.selected == true {
		return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
	}
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func InitialSort(height int) SortModel {
	return SortModel{
		text:     "Sort",
		selected: false,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#6E3F00")).
			BorderTop(true).
			Width(18).
			Height(height).
			Align(lipgloss.Center),
	}
}

func (m SortModel) Init() tea.Cmd {
	return nil
}

func (m SortModel) Update(msg tea.Msg) (SortModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "L", "H", "J", "K":
			m.style = m.toggleBorder()
			m.selected = !m.selected
		}
	}
	return m, nil
}

func (m SortModel) View() string {
	return m.style.Render(m.text)
}
