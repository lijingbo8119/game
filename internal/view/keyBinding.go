package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyBinding struct {
	key.Binding
	handler func(View, *Views) tea.Cmd
}

func (r *keyBinding) SetHandler(handler func(currentView View, views *Views) tea.Cmd) *keyBinding {
	r.handler = handler
	return r
}

func newKeyBinding(opts ...key.BindingOpt) *keyBinding {
	return &keyBinding{
		Binding: key.NewBinding(opts...),
	}
}
