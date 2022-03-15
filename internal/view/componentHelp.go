package view

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type componentHelp struct {
	h       help.Model
	v       View
	keyHelp *keyBinding
	keyQuit *keyBinding
}

func (r *componentHelp) SetView(v View) {
	r.v = v
}

func (r componentHelp) ShortHelp() []key.Binding {
	return []key.Binding{
		r.keyHelp.Binding,
		r.keyQuit.Binding,
	}
}

func (r componentHelp) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		r.v.keyBindings().toRawBindings(),
		{
			r.keyHelp.Binding,
			r.keyQuit.Binding,
		},
	}
}
