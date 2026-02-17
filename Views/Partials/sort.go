package partials

import "github.com/charmbracelet/bubbletea"

type SortModel struct {
	text string
}

func InitialSort() SortModel {
	return SortModel{
		text: "Sort",
	}
}

func (m SortModel) Init() tea.Cmd {
	return nil
}

func (m SortModel) Update(msg tea.Msg) (SortModel, tea.Cmd) {
	return m, nil
}

func (m SortModel) View() string {
	return m.text
}
