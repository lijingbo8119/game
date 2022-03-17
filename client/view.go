package client

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

type view interface {
	fmt.Stringer
	init() view
	UpdateWindowSizeMsg(msg tea.WindowSizeMsg) (cmd tea.Cmd)
	Update(msg tea.Msg) (cmd tea.Cmd)
	Name() string
	SetActive(bool)
	IsActive() bool
	KeyBindings() KeyBindings
}

func viewInit() {
	viewProgram = tea.NewProgram(&viewModel{})
}

type viewModel struct{}

func (m *viewModel) Init() tea.Cmd {
	return nil
}

func (m *viewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		lo.ForEach(views, func(v view, i int) {
			if v.IsActive() {
				cmd = v.UpdateWindowSizeMsg(msg)
			} else {
				v.UpdateWindowSizeMsg(msg)
			}
		})
	}
	if activeView() != nil {
		cmd = activeView().Update(msg)
	}
	return m, cmd
}

func (m viewModel) View() string {
	return activeView().String()
}

var viewProgram *tea.Program
