package components

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type TagInputModel struct {
	textInput  textinput.Model
	tags       []string
	tagsCursor int
	title      string
	Selected   bool

	errorMsg string
}

func InitialInput(tagCnt int, placeholder string, title string, width int) TagInputModel {
	tags := make([]string, tagCnt)

	input := textinput.New()
	input.Placeholder = placeholder
	input.SetVirtualCursor(false)
	input.Focus()
	input.CharLimit = 64
	input.SetWidth(width)

	return TagInputModel{
		tags:       tags,
		textInput:  input,
		tagsCursor: 0,
		title:      title,
		Selected:   false,
	}
}

func (m TagInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TagInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.errorMsg = ""

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		m.errorMsg = msg.String()
		switch msg.String() {
		case "up", "down":
		case "delete":
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TagInputModel) View() tea.View {
	var s string
	s = m.title
	s = lipgloss.JoinVertical(lipgloss.Left, s, m.textInput.View())
	//s = lipgloss.JoinVertical(lipgloss.Center, s, m.errorMsg)
	return tea.NewView(s)
}
