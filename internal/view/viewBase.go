package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type viewBase struct {
	active bool
	help   componentHelp
}

func (r *viewBase) init() {
	r.help = componentHelp{h: help.New()}
	r.help.keyHelp = newKeyBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	).SetHandler(func(currentView View, views *Views) tea.Cmd {
		r.help.h.ShowAll = !r.help.h.ShowAll
		return nil
	})
	r.help.keyQuit = newKeyBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	).SetHandler(func(currentView View, views *Views) tea.Cmd {
		// TODO
		return nil
	})
}

func (r *viewBase) UpdateWindowSizeMsg(msg tea.WindowSizeMsg) (cmd tea.Cmd) {
	r.help.h.Width = msg.Width
	return nil
}

func (r *viewBase) UpdateKeyMsg(msg tea.KeyMsg, views *Views) (cmd tea.Cmd) {
	if key.Matches(msg, r.help.keyHelp.Binding) {
		r.help.keyHelp.handler(nil, views)
	}
	if key.Matches(msg, r.help.keyQuit.Binding) {
		r.help.keyQuit.handler(nil, views)
	}
	return nil
}

func (r *viewBase) SetActive(active bool) {
	r.active = active
}

func (r viewBase) IsActive() bool {
	return r.active
}
