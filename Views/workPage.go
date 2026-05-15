package views

import (
	key "charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	components "github.com/thomasmckinstry/MediaLogger-TUI/Views/Components"
	. "github.com/thomasmckinstry/MediaLogger-TUI/utils"
)

type WorkPageModel struct {
	work          *components.WorkFormModel
	width, height int
	tabCursor     int
	mainCursor    int
	headerCursor  int
	rightCursor   int
	tabStyle      lipgloss.Style
	tabsStyle     lipgloss.Style
	buttonStyle   lipgloss.Style
	displayStyle  lipgloss.Style
	detailsStyle  lipgloss.Style
}

type workKeyMap struct {
	TopLevelUp    key.Binding
	TopLevelDown  key.Binding
	TopLevelLeft  key.Binding
	TopLevelRight key.Binding
	Confirm       key.Binding
	Exit          key.Binding
}

var (
	unfocused = lipgloss.Color("#6E3F00")
	focused   = lipgloss.Color("#D17600")
)

var defaultWorkMap = workKeyMap{
	TopLevelUp:    key.NewBinding(key.WithKeys("K")),
	TopLevelDown:  key.NewBinding(key.WithKeys("J")),
	TopLevelLeft:  key.NewBinding(key.WithKeys("H")),
	TopLevelRight: key.NewBinding(key.WithKeys("L")),
	Confirm:       key.NewBinding(key.WithKeys("enter")),
	Exit:          key.NewBinding(key.WithKeys("esc")),
}

func InitialWorkPage(width, height int) *WorkPageModel {
	workForm := components.InitialWorkFormModel(22, height)
	return &WorkPageModel{
		work:   workForm,
		width:  width,
		height: height,
		tabStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(unfocused),
		tabsStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			MarginLeft(1).
			MarginRight(1).
			BorderForeground(unfocused),
		detailsStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderRight(true).
			BorderForeground(unfocused),
		buttonStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(unfocused),
		displayStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.DoubleBorder()).
			Width(width - 23).
			BorderTop(true).
			BorderForeground(unfocused),
	}
}

func (m *WorkPageModel) Init() tea.Cmd {
	return nil
}

func (m *WorkPageModel) Update(msg tea.Msg) (*WorkPageModel, tea.Cmd) {
	var (
		cmd, cmds tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		_, cmd = m.work.Update(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultWorkMap.Exit):
			cmd = func() tea.Msg { return ViewMsg(0) }
		}
	}

	cmds = tea.Batch(cmds, cmd)
	return m, cmds
}

func (m *WorkPageModel) View() tea.View {
	var (
		c *tea.Cursor
		s string
	)
	s = lipgloss.JoinHorizontal(lipgloss.Top, s, m.detailsStyle.Render(m.work.View().Content))

	header := m.buttonStyle.Render(lipgloss.PlaceHorizontal(10, lipgloss.Center, "ADD"))

	notes := m.tabStyle.Render(lipgloss.PlaceHorizontal(10, lipgloss.Center, "NOTES"))
	reviews := m.tabStyle.Render(lipgloss.PlaceHorizontal(10, lipgloss.Center, "REVIEWS"))

	tabs := m.tabsStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, notes, " ", reviews))
	header = lipgloss.JoinHorizontal(lipgloss.Top, tabs, header)
	display := m.displayStyle.Render("")

	rightSide := lipgloss.JoinVertical(lipgloss.Left, header, display)

	s = lipgloss.JoinHorizontal(lipgloss.Top, s, rightSide)

	v := tea.NewView(s)
	v.Cursor = c
	v.AltScreen = true
	return v
}
