package view

import "github.com/charmbracelet/bubbles/key"

type KeyBindingBooleanClosure = func(*keyBinding) bool

type KeyBindings []*keyBinding

func (r KeyBindings) First(closure ...KeyBindingBooleanClosure) *keyBinding {
	_closure := func() KeyBindingBooleanClosure {
		if len(closure) > 0 {
			return closure[0]
		}
		return func(c *keyBinding) bool { return true }
	}()
	for _, c := range r {
		if _closure(c) {
			return c
		}
	}
	return nil
}

func (r KeyBindings) toRawBindings() []key.Binding {
	res := []key.Binding{}
	for _, kb := range r {
		res = append(res, kb.Binding)
	}
	return res
}
