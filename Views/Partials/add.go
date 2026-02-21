package partials

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AddModel struct {
	text     string
	selected bool
	style    lipgloss.Style
}

func (m AddModel) toggleBorder() lipgloss.Style {
	if m.selected == true {
		return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
	}
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func InitialAdd() AddModel {
	return AddModel{
		text:     "Add",
		selected: true,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(lipgloss.Color("#D17600")).
			Width(18).
			Height(1).
			Align(lipgloss.Center).
			MarginLeft(1),
	}
}

func (m AddModel) Init() tea.Cmd {
	return nil
}

func (m AddModel) Update(msg tea.Msg) (AddModel, tea.Cmd) {
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

func (m AddModel) View() string {
	return m.style.Render(m.text)
}
