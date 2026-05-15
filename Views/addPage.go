package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/thomasmckinstry/MediaLogger-TUI/Views/Components"
)

type AddModel struct {
	form          *components.WorkFormModel
	style         lipgloss.Style
	width, height int
}

func InitialAddModel(width, height int) *AddModel {
	form := components.InitialWorkFormModel(width, height)
	return &AddModel{
		form: form,
		style: lipgloss.NewStyle().
			Height(height).
			Align(lipgloss.Center).
			PaddingLeft(1).
			PaddingRight(1).
			BorderStyle(lipgloss.DoubleBorder()),
	}
}

func (m *AddModel) Init() tea.Cmd {
	return nil
}

func (m *AddModel) Update(msg tea.Msg) (*AddModel, tea.Cmd) {
	var cmds tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style = m.style.Height(msg.Height - (7))
		m.form.Update(msg)
	default:
		_, cmd = m.form.Update(msg)
		cmds = tea.Batch(cmds, cmd)
	}
	return m, cmds
}

func (m *AddModel) View() tea.View {
	var c *tea.Cursor

	formView := m.form.View()
	c = formView.Cursor
	if c != nil {
		c.Y += 1
		c.X += 2
	}
	v := tea.NewView(m.style.Render(formView.Content))
	v.Cursor = c
	v.AltScreen = true
	return v
}
