package views

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	partials "github.com/thomasmckinstry/MediaLogger-TUI/Views/Partials"

	"log"
	"os"
)

var height int

type homeKeyMap struct {
	TopLevelUp    key.Binding
	TopLevelDown  key.Binding
	TopLevelLeft  key.Binding
	TopLevelRight key.Binding
	SidebarNav    key.Binding
	Confirm       key.Binding
}

var defaultHomeKeyMap = homeKeyMap{
	TopLevelUp: key.NewBinding(
		key.WithKeys("K"),
		key.WithHelp("K", "Move up between sections"),
	),
	TopLevelDown: key.NewBinding(
		key.WithKeys("J"),
		key.WithHelp("J", "Move down between sections"),
	),
	TopLevelLeft: key.NewBinding(
		key.WithKeys("H"),
		key.WithHelp("H", "Move left between sections"),
	),
	TopLevelRight: key.NewBinding(
		key.WithKeys("L"),
		key.WithHelp("L", "Move right between sections"),
	),
	SidebarNav: key.NewBinding(
		key.WithKeys("k", "up", "j", "down"),
		key.WithHelp("k/↑", "Move up within a section"),
		key.WithHelp("j/↓", "Move down within a section"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "Confirm an input or focus a component"),
	),
}

type HomeModel struct {
	sidebarCursor int
	mainCursor    int
	sidebarViews  []tea.Model
	listModel     tea.Model
}

type ViewMsg int

func InitialHome(width int, height int) *HomeModel {
	list := partials.InitialList(width-19, height)
	add := partials.InitialAdd() // height = 1 Note: I think each side of the border adds 1
	filter := partials.InitialFilter(height - (7))
	sort := partials.InitialSort(3)

	sidebarList := []tea.Model{}               //make([]tea.Model, 3)
	sidebarList = append(sidebarList, &add)    //[0] = add
	sidebarList = append(sidebarList, &filter) //[1] = filter
	sidebarList = append(sidebarList, &sort)   //[2] = sort

	return &HomeModel{
		sidebarViews:  sidebarList,
		listModel:     &list,
		sidebarCursor: 0,
		mainCursor:    0,
	}
}

func (m *HomeModel) Init() tea.Cmd {
	return nil
}

func (m *HomeModel) Update(msg tea.Msg) (*HomeModel, tea.Cmd) {
	var cmds tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		_, cmd = m.listModel.Update(msg)
		cmds = tea.Batch(cmds, cmd)
		_, cmd = m.sidebarViews[1].Update(msg)
		cmds = tea.Batch(cmds, cmd)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, defaultHomeKeyMap.TopLevelUp):
			if m.mainCursor == 0 && m.sidebarCursor > 0 {
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
				m.sidebarCursor--
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
			}
		case key.Matches(msg, defaultHomeKeyMap.TopLevelDown):
			if m.mainCursor == 0 && m.sidebarCursor < 2 {
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
				m.sidebarCursor++
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
			}
		case key.Matches(msg, defaultHomeKeyMap.TopLevelLeft):
			if m.mainCursor > 0 {
				m.mainCursor--
				m.listModel, cmd = m.listModel.Update(msg)
				cmds = tea.Batch(cmds, cmd)
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
			}
		case key.Matches(msg, defaultHomeKeyMap.TopLevelRight):
			if m.mainCursor < 1 {
				m.mainCursor++
				m.listModel, cmd = m.listModel.Update(msg)
				cmds = tea.Batch(cmds, cmd)
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
			}
		case key.Matches(msg, defaultHomeKeyMap.SidebarNav):
			if m.mainCursor == 1 {
				m.listModel, cmd = m.listModel.Update(msg)
				cmds = tea.Batch(cmds, cmd)
			} else {
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
			}
		case key.Matches(msg, defaultHomeKeyMap.Confirm):
			if m.sidebarCursor == 0 && m.mainCursor == 0 {
				if len(os.Getenv("DEBUG")) > 0 {
					log.Println("homePage sending AddMsg")
				}
				cmds = tea.Batch(cmds, func() tea.Msg { return (ViewMsg(1)) })
			} else {
				_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
				cmds = tea.Batch(cmds, cmd)
				if cmd != nil {
					msg, ok := cmd().(partials.FilterMsg)
					if ok {
						_, cmd = m.listModel.Update(msg)
						cmds = tea.Batch(cmds, cmd)
					}
				}
			}
		default:
			_, cmd = m.sidebarViews[m.sidebarCursor].Update(msg)
			cmds = tea.Batch(cmds, cmd)
			if cmd != nil {
				msg, ok := cmd().(partials.SortMsg)
				if ok {
					_, cmd = m.listModel.Update(msg)
					cmds = tea.Batch(cmds, cmd)
				}
			}
		}
	}

	return m, cmds
}

func (m *HomeModel) View() tea.View {
	var c *tea.Cursor
	s := ""
	sidebarContent := []string{}
	for _, form := range m.sidebarViews {
		formView := form.View()
		sidebarContent = append(sidebarContent, formView.Content)

		if formView.Cursor != nil {
			c = formView.Cursor
			c.Y += lipgloss.Height(s)
		}
	}
	sidebar := lipgloss.JoinVertical(lipgloss.Center, sidebarContent...)

	list := m.listModel.View()
	s = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, list.Content)

	// Send the UI for rendering
	view := tea.NewView(s)
	view.Cursor = c
	view.AltScreen = true
	return view
}
