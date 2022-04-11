package client

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

type viewComponentButton struct {
	viewComponentBase
	Content string
	Active  bool
	Updater viewModelUpdater
}

func (r viewComponentButton) Name() string {
	return fmt.Sprintf("%s:%s", reflect.TypeOf(r).Name(), r.Content)
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r viewComponentButton) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r.Updater(msg)
}

func (r viewComponentButton) View() string {
	return lo.Switch[bool, lipgloss.Style](r.Active).
		Case(false, buttonStyle).
		Default(activeButtonStyle).
		Copy().
		Render(r.Content)
}
