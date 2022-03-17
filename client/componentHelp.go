package client

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type componentHelp struct {
	h       help.Model
	v       view
	keyHelp *KeyBinding
	keyQuit *KeyBinding
}

func (r *componentHelp) init() component {
	r.keyHelp = NewKeyBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "toggle help"),
	).SetHandler(func(currentview view) tea.Cmd {
		r.h.ShowAll = !r.h.ShowAll
		return nil
	})
	r.keyQuit = NewKeyBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	).SetHandler(func(currentview view) tea.Cmd {
		return tea.Quit
	})
	return r
}

func (r *componentHelp) setView(v view) {
	r.v = v
}

func (r componentHelp) String() string {
	return r.h.View(r)
}

func (r componentHelp) ShortHelp() []key.Binding {
	return []key.Binding{
		r.keyHelp.Binding,
		r.keyQuit.Binding,
	}
}

func (r componentHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		(r.v).KeyBindings().ToRawBindings(),
		{},
		{
			r.keyHelp.Binding,
			r.keyQuit.Binding,
		},
	}
}
