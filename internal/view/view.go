package view

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type View interface {
	fmt.Stringer
	init()
	UpdateWindowSizeMsg(msg tea.WindowSizeMsg) (cmd tea.Cmd)
	Update(msg tea.Msg, views *Views) (cmd tea.Cmd)
	Name() string
	SetActive(bool)
	IsActive() bool
	Factory() View
	keyBindings() KeyBindings
}
