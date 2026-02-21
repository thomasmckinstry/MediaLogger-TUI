package partials

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilterModel struct {
	text     string
	selected bool
	style    lipgloss.Style
}

func (m FilterModel) toggleBorder() lipgloss.Style {
	if m.selected == true {
		return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
	}
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func InitialFilter(height int) FilterModel {
	return FilterModel{
		text:     "Filter",
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

func (m FilterModel) Init() tea.Cmd {
	return nil
}

func (m FilterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style = m.style.Height(msg.Height - (9))
	case tea.KeyMsg:
		switch msg.String() {
		case "L", "H", "J", "K":
			m.style = m.toggleBorder()
			m.selected = !m.selected
		}
	}
	return m, nil
}

func (m FilterModel) View() string {
	return m.style.Render(m.text)
}
