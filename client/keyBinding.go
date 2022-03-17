package client

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyBinding struct {
	key.Binding
	handler func(view) tea.Cmd
}

func (r *KeyBinding) SetHandler(handler func(currentview view) tea.Cmd) *KeyBinding {
	r.handler = handler
	return r
}

func (r *KeyBinding) Handler() func(currentview view) tea.Cmd {
	return r.handler
}

func NewKeyBinding(opts ...key.BindingOpt) *KeyBinding {
	return &KeyBinding{
		Binding: key.NewBinding(opts...),
	}
}
