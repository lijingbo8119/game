package view

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewLogin struct {
	viewBase
	mode          string
	usernameInput textinput.Model
	passwordInput textinput.Model
	nicknameInput textinput.Model
}

func (r *ViewLogin) init() {
	fmt.Println("ViewLogin init")
	r.viewBase.init()
	r.help.SetView(r)

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
}

func (r ViewLogin) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *ViewLogin) Update(msg tea.Msg, views *Views) (cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		r.viewBase.UpdateKeyMsg(msg, views)
		keyBinding := r.keyBindings().First(func(kb *keyBinding) bool {
			return key.Matches(msg, kb.Binding)
		})
		if keyBinding != nil {
			keyBinding.handler(r, views)
		}
	}
	if r.usernameInput.Focused() {
		r.usernameInput, cmd = r.usernameInput.Update(msg)
	}
	if r.passwordInput.Focused() {
		r.passwordInput, cmd = r.passwordInput.Update(msg)
	}
	if r.nicknameInput.Focused() {
		r.nicknameInput, cmd = r.nicknameInput.Update(msg)
	}
	return cmd
}

func (r ViewLogin) String() string {
	if r.mode == "signup" {
		return fmt.Sprintf(
			"What’s your favorite Pokémon?\n\n%s\n\n%s\n\n%s\n\n%s",
			r.usernameInput.View(),
			r.passwordInput.View(),
			r.nicknameInput.View(),
			r.help.h.View(r.help),
		)
	}
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s\n\n%s",
		r.usernameInput.View(),
		r.passwordInput.View(),
		r.help.h.View(r.help),
	)
}

func (r *ViewLogin) keyBindings() KeyBindings {
	return KeyBindings{
		newKeyBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		).SetHandler(func(currentView View, views *Views) tea.Cmd {
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
		newKeyBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		).SetHandler(func(currentView View, views *Views) tea.Cmd {
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
		newKeyBinding(
			key.WithKeys("ctrl+b", "ctrl+b"),
			key.WithHelp("ctrl+b", "sign in / sign up"),
		).SetHandler(func(currentView View, views *Views) tea.Cmd {
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
	}
}

func (r ViewLogin) Factory() View {
	v := &ViewLogin{}
	v.init()
	return v
}
