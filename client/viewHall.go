package client

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewHall struct {
	viewBase
	mode          string
	usernameInput textinput.Model
	passwordInput textinput.Model
	nicknameInput textinput.Model
}

func (r *viewHall) init() view {
	r.viewBase.init(r)

	r.usernameInput = textinput.New()
	r.usernameInput.Focus()
	r.usernameInput.Placeholder = "username"
	r.usernameInput.CharLimit = 156
	r.usernameInput.Width = 20

	r.passwordInput = textinput.New()
	r.passwordInput.Placeholder = "password"
	r.passwordInput.CharLimit = 156
	r.passwordInput.Width = 20

	r.nicknameInput = textinput.New()
	r.nicknameInput.Placeholder = "nickname"
	r.nicknameInput.CharLimit = 156
	r.nicknameInput.Width = 20

	r.mode = "signin"
	return r
}

func (r viewHall) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewHall) Update(msg tea.Msg) (cmd tea.Cmd) {
	if r.usernameInput.Focused() {
		r.usernameInput, cmd = r.usernameInput.Update(msg)
	}
	if r.passwordInput.Focused() {
		r.passwordInput, cmd = r.passwordInput.Update(msg)
	}
	if r.nicknameInput.Focused() {
		r.nicknameInput, cmd = r.nicknameInput.Update(msg)
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = r.viewBase.UpdateKeyMsg(msg)
		r.KeyBindings().Each(func(kb *KeyBinding, index int) bool {
			if key.Matches(msg, kb.Binding) {
				kb.Handler()(r)
			}
			return true
		})
	}
	return cmd
}

func (r viewHall) String() string {
	return r.viewBase.view(fmt.Sprintf(
		"%s\n\n%s",
		r.usernameInput.View(),
		r.passwordInput.View(),
	))
}

func (r *viewHall) KeyBindings() KeyBindings {
	return KeyBindings{
		NewKeyBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		).SetHandler(func(currentview view) tea.Cmd {
			if r.mode == "signup" && r.nicknameInput.Focused() {
				r.usernameInput.Blur()
				r.passwordInput.Focus()
				r.nicknameInput.Blur()
				return nil
			}
			r.usernameInput.Focus()
			r.passwordInput.Blur()
			r.nicknameInput.Blur()
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		).SetHandler(func(currentview view) tea.Cmd {
			if r.mode == "signup" && r.passwordInput.Focused() {
				r.usernameInput.Blur()
				r.passwordInput.Blur()
				r.nicknameInput.Focus()
				return nil
			}
			r.usernameInput.Blur()
			r.passwordInput.Focus()
			r.nicknameInput.Blur()
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "sign in / sign up"),
		).SetHandler(func(currentview view) tea.Cmd {
			if r.mode == "signup" {
				r.mode = "signin"
			} else {
				r.mode = "signup"
			}
			if r.mode == "signup" && r.nicknameInput.Focused() {
				r.usernameInput.Blur()
				r.passwordInput.Focus()
				r.nicknameInput.Blur()
				return nil
			}
			r.usernameInput.Focus()
			r.passwordInput.Blur()
			r.nicknameInput.Blur()
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter"),
		).SetHandler(func(currentview view) tea.Cmd {
			return nil
		}),
	}
}
