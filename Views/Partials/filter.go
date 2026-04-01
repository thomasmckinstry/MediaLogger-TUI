package partials

import (
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/thomasmckinstry/Bubbletea-Tutorial/Views/Components"
)

// TODO: I can probably sub out most of this file for a huh? component
// Don't want to do that because it's more hands off
// TODO: Need a different way to index into the components because of different Model types.
type FilterModel struct {
	headerText     string
	titleInput     tea.Model
	genreInput     textinput.Model
	themeInput     textinput.Model // TODO: These need additional displays for previous entries
	selected       bool
	cursor         int
	forms          []tea.Model // Can I get this to use pointers to the actual models? I think right now I'm copying them
	status         []string
	genres         []string
	themes         []string
	style          lipgloss.Style
	headerStyle    lipgloss.Style
	textinputStyle lipgloss.Style

	errorMsg string
}

// TODO: This should be a utils
func (m FilterModel) toggleBorder() lipgloss.Style {
	if m.selected == true {
		return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
	}
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func InitialFilter(height int) FilterModel {
	/*titleInput := textinput.New()
	// TODO: I should probably have this colored differently or something to show that it's input instead of a descriptor
	// Also make it spaced properly so it's always taking up all the width of the column
	titleInput.Placeholder = "title"
	titleInput.CharLimit = 64
	titleInput.SetWidth(14)
	//titleInput.Blur()*/
	titleInput := components.InitialInput(3, "title", "Title", 14)

	genreInput := textinput.New()
	// TODO: I should probably have this colored differently or something to show that it's input instead of a descriptor
	// Also make it spaced properly so it's always taking up all the width of the column
	genreInput.Placeholder = "genre"
	genreInput.ShowSuggestions = true
	genreInput.SetSuggestions([]string{"Fantasy", "Drama"}) // TODO: Suggestions will be subbed in with an array of every work.
	genreInput.CharLimit = 64
	genreInput.SetWidth(14)
	//genreInput.Blur()

	//status := []string{"Completed", "In Progress", "Started", "Pending", "Dropped"}
	forms := []tea.Model{titleInput} // TODO: Figure out how to have null pointers to each form
	// forms is an array of all the forms that make up the filter box.
	// This is so I can index into each one as I navigate with the keyboard

	return FilterModel{
		headerText: "Filter",
		titleInput: titleInput,
		genreInput: genreInput,
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
			Align(lipgloss.Center).
			Width(16),
		textinputStyle: lipgloss.NewStyle().
			Align(lipgloss.Center).
			MarginTop(1).
			Width(16),
	}
}

func (m FilterModel) Init() tea.Cmd {
	return nil
}

func (m FilterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.errorMsg = ""

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style = m.style.Height(msg.Height - (9))
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter":
			m.forms[m.cursor], cmd = m.forms[m.cursor].Update(msg)
		case "L", "H", "J", "K":
			m.style = m.toggleBorder()
			m.selected = !m.selected
		case "j", "down": // TODO: Make these check for focused inputs before moving the cursor
			if m.cursor < len(m.forms)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		default:
			/*if field, ok := m.forms[m.cursor].(textinput.Model); ok {
				if field.Focused() {
					m.forms[m.cursor], cmd = field.Update(msg)
				}
			}*/
			m.forms[m.cursor], cmd = m.forms[m.cursor].Update(msg)
		}
	}
	return m, cmd
}

// TODO: Add styling to make it clear that a textbox is selected.
// TODO: Iterate over m.forms instead of having a bunch of different conditional blocks
func (m FilterModel) View() tea.View {
	//header:
	s := m.headerStyle.Render(m.headerText)
	// Text Input (Title):
	//if field, ok := m.forms[0].(textinput.Model); ok {
	//	titleInput := textinput.Model(field)
	s = lipgloss.JoinVertical(lipgloss.Center, s, m.textinputStyle.Render(m.forms[m.cursor].View().Content))
	//}

	//Genres (text -> tags):
	/*if field, ok := m.forms[1].(textinput.Model); ok {
		genreInput := textinput.Model(field)
		s = lipgloss.JoinVertical(lipgloss.Center, s, m.textinputStyle.Render(genreInput.View()))
	}*/
	// Status (Checkboxes):

	//Themes (text -> tags):

	s = lipgloss.JoinVertical(lipgloss.Center, s, m.errorMsg)
	v := tea.NewView(m.style.Render(s))
	return v
}
