package client

import (
	"fmt"
	"game/server"
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewSignup struct {
	viewBase
	usernameInput textinput.Model
	passwordInput textinput.Model
	nicknameInput textinput.Model
}

func (r *viewSignup) init() view {
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
	return r
}

func (r viewSignup) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewSignup) Update(msg tea.Msg) (cmd tea.Cmd) {
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

func (r viewSignup) String() string {
	return r.viewBase.view(fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		r.usernameInput.View(),
		r.passwordInput.View(),
		r.nicknameInput.View(),
	))
}

func (r *viewSignup) KeyBindings() KeyBindings {
	return KeyBindings{
		NewKeyBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		).SetHandler(func(currentview view) tea.Cmd {
			if r.nicknameInput.Focused() {
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
			if r.passwordInput.Focused() {
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
			goToView((viewSignin{}).Name())
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter"),
		).SetHandler(func(currentview view) tea.Cmd {
			requests[server.CmdSignupRequest](server.Data{
				Params: server.EventSignupParams{
					Username: r.usernameInput.Value(),
					Password: r.passwordInput.Value(),
					Nickname: r.nicknameInput.Value(),
				},
			})
			return nil
		}),
	}
}
