package partials

import (
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type ListModel struct {
	style lipgloss.Style
	table table.Model
}

func (m ListModel) selectView() lipgloss.Style {
	return m.style.BorderForeground(lipgloss.Color("#D17600"))
}

func (m ListModel) deselectView() lipgloss.Style {
	return m.style.BorderForeground(lipgloss.Color("#6E3F00"))
}

func InitialList(width int, height int, rows []table.Row) ListModel {

	var columns = []table.Column{
		{Title: "Title", Width: width / 4},
		{Title: "Medium", Width: width / 8},
		{Title: "Status", Width: width / 8},
		{Title: "Tags", Width: width / 3},
		{Title: "Released", Width: width / 6},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithWidth(width),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6E3F00")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#D17600")).
		Bold(false)
	t.SetStyles(s)

	return ListModel{
		style: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderLeft(true).
			BorderForeground(lipgloss.Color("#6E3F00")).
			PaddingTop(1).
			Width(width).
			Height(height),
		table: t,
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.style = m.style.Height(msg.Height).Width(msg.Width - 18)
		width := msg.Width - 29
		m.table.SetColumns([]table.Column{
			{Title: "Title", Width: width / 4},
			{Title: "Medium", Width: width / 8},
			{Title: "Status", Width: width / 8},
			{Title: "Tags", Width: width / 3},
			{Title: "Released", Width: width / 6},
		})
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "L":
			m.style = m.selectView()
			m.table.Focus()
		case "H":
			m.style = m.deselectView()
			m.table.Blur()
		case "j", "k", "up", "down":
			m.table, cmd = m.table.Update(msg)
		}
	}
	return m, cmd
}

func (m ListModel) View() tea.View {
	return tea.NewView(m.style.Render(m.table.View()))
}
