package client

import (
	tea "github.com/charmbracelet/bubbletea"
)

type viewModel interface {
	tea.Model
	Name() string
}

type viewModelBase struct{}

func (r viewModelBase) Init() tea.Cmd {
	return nil
}

type viewModelUpdater func(msg tea.Msg) (tea.Model, tea.Cmd)
