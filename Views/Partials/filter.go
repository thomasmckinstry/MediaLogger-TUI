package partials

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilterModel struct {
	headerText  string
	titleInput  textinput.Model
	selected    bool
	cursor      int
	forms       []any
	status      []string
	genres      []string
	themes      []string
	style       lipgloss.Style
	headerStyle lipgloss.Style
}

func (m FilterModel) toggleBorder() lipgloss.Style {
	if m.selected == true {
		return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
	}
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func InitialFilter(height int) FilterModel {
	titleInput := textinput.New()
	// TODO: I should probably have this colored differently or something to show that it's input instead of a descriptor
	// Also make it spaced properly so it's always taking up all the width of the column
	titleInput.Placeholder = "title"
	titleInput.ShowSuggestions = true
	titleInput.SetSuggestions([]string{"test1", "test2"})
	//titleInput.Blur()
	titleInput.CharLimit = 64
	titleInput.Width = 18
	titleInput.Focus() // TODO: This is here for proof of concept. Should be removed, focus should be set in update is cursor is on titleInput

	//status := []string{"Completed", "In Progress", "Started", "Pending", "Dropped"}
	forms := []any{titleInput}
	// forms is an array of all the forms that make up the filter box.
	// This is so I can index into each one as I navigate with the keyboard

	return FilterModel{
		headerText: "Filter",
		titleInput: titleInput,
		selected:   false,
		cursor:     0,
		forms:      forms,
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#6E3F00")).
			BorderTop(true).
			Width(18).
			Height(height).
			Align(lipgloss.Center),
		headerStyle: lipgloss.NewStyle().
			Align(lipgloss.Center),
	}
}

func (m FilterModel) Init() tea.Cmd {
	return textinput.Blink // TODO: Move this to update when textinput is selected
}

func (m FilterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style = m.style.Height(msg.Height - (9))
	case tea.KeyMsg:
		switch msg.String() {
		case "L", "H", "J", "K":
			m.style = m.toggleBorder()
			m.selected = !m.selected
		case "j", "down":
			if m.cursor > 0 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor < len(m.forms) {
				m.cursor--
			}
		default:
			m.titleInput, cmd = m.titleInput.Update(msg)
		}
	}
	return m, cmd
}

// TODO: Add styling to make it clear that a textbox is selected.
func (m FilterModel) View() string {
	//header:
	s := m.headerStyle.Render(m.headerText)
	// Text Input (Title):
	s = lipgloss.JoinVertical(lipgloss.Left, s, m.titleInput.View())
	// Status (Checkboxes):

	//Genres (Checkboxes):

	//Themes (Checkboxes):

	return m.style.Render(s)
}
