package client

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewBase struct {
	active  bool
	network component
	help    component
}

func (r *viewBase) init(v view) {
	help := &componentHelp{h: help.New()}
	help.init()
	help.setView(v)
	r.help = help
}

func (r *viewBase) UpdateWindowSizeMsg(msg tea.WindowSizeMsg) (cmd tea.Cmd) {
	help, _ := r.help.(*componentHelp)
	help.h.Width = msg.Width
	return nil
}

func (r *viewBase) UpdateKeyMsg(msg tea.KeyMsg) (cmd tea.Cmd) {
	help, _ := r.help.(*componentHelp)
	if key.Matches(msg, help.keyHelp.Binding) {
		return help.keyHelp.Handler()(nil)
	}
	if key.Matches(msg, help.keyQuit.Binding) {
		return help.keyQuit.Handler()(nil)
	}
	return nil
}

func (r *viewBase) SetActive(active bool) {
	help, _ := r.help.(*componentHelp)
	if !active {
		help.h.ShowAll = false
	}
	r.active = active
}

func (r viewBase) IsActive() bool {
	return r.active
}

func (r viewBase) view(innerView string) string {
	var style = lipgloss.NewStyle().
		Padding(2).
		Margin(1).
		Background(lipgloss.Color("12")).
		Foreground(lipgloss.Color("12"))
	return style.Render(fmt.Sprintf(
		"signin\n\n%s\n\n%s\n\n%s",
		"title",
		innerView,
		r.help.String(),
	))
}
