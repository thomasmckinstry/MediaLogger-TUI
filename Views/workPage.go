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
}

type workKeyMap struct {
	TopLevelUp    key.Binding
	TopLevelDown  key.Binding
	TopLevelLeft  key.Binding
	TopLevelRight key.Binding
	Confirm       key.Binding
	Exit          key.Binding
}

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
	s = lipgloss.JoinHorizontal(lipgloss.Top, s, m.work.View().Content)

	v := tea.NewView(s)
	v.Cursor = c
	v.AltScreen = true
	return v
}
