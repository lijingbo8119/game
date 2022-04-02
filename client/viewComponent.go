package client

import tea "github.com/charmbracelet/bubbletea"

type viewComponent viewModel

type viewComponentBase struct{}

func (r viewComponentBase) Init() tea.Cmd {
	return nil
}
