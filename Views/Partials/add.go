package partials

import "github.com/charmbracelet/bubbletea"

type AddModel struct {
	text string
}

func InitialAdd() AddModel {
	return AddModel{
		text: "Add",
	}
}

func (m AddModel) Init() tea.Cmd {
	return nil
}

func (m AddModel) Update(msg tea.Msg) (AddModel, tea.Cmd) {
	return m, nil
}

func (m AddModel) View() string {
	return m.text
}
