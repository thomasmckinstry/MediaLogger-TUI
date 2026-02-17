package partials

import "github.com/charmbracelet/bubbletea"

type FilterModel struct {
	text string
}

func InitialFilter() FilterModel {
	return FilterModel{
		text: "Filter",
	}
}

func (m FilterModel) Init() tea.Cmd {
	return nil
}

func (m FilterModel) Update(msg tea.Msg) (FilterModel, tea.Cmd) {
	return m, nil
}

func (m FilterModel) View() string {
	return m.text
}
