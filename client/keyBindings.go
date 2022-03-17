package client

import "github.com/charmbracelet/bubbles/key"

type KeyBindingBooleanIndexClosure = func(*KeyBinding, int) bool

type KeyBindings []*KeyBinding

func (r KeyBindings) Each(closure KeyBindingBooleanIndexClosure) {
	for i, c := range r {
		if !closure(c, i) {
			break
		}
	}
}

func (r KeyBindings) ToRawBindings() []key.Binding {
	res := []key.Binding{}
	for _, kb := range r {
		res = append(res, kb.Binding)
	}
	return res
}
