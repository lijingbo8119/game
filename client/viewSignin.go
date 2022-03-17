package client

import (
	"fmt"
	"game/server"
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewSignin struct {
	viewBase
	usernameInput textinput.Model
	passwordInput textinput.Model
	message       component
}

func (r *viewSignin) init() view {
	r.viewBase.init(r)
	r.message = &componentMessage{}

	r.usernameInput = textinput.New()
	r.usernameInput.Focus()
	r.usernameInput.Placeholder = "username"
	r.usernameInput.CharLimit = 156
	r.usernameInput.Width = 20

	r.passwordInput = textinput.New()
	r.passwordInput.Placeholder = "password"
	r.passwordInput.CharLimit = 156
	r.passwordInput.Width = 20
	return r
}

func (r viewSignin) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewSignin) Update(msg tea.Msg) (cmd tea.Cmd) {
	if r.usernameInput.Focused() {
		r.usernameInput, cmd = r.usernameInput.Update(msg)
	}
	if r.passwordInput.Focused() {
		r.passwordInput, cmd = r.passwordInput.Update(msg)
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

func (r viewSignin) String() string {
	return r.viewBase.view(fmt.Sprintf(
		"%s\n\n%s\n%s",
		r.usernameInput.View(),
		r.passwordInput.View(),
		r.message.String(),
	))
}

func (r *viewSignin) KeyBindings() KeyBindings {
	return KeyBindings{
		NewKeyBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		).SetHandler(func(currentview view) tea.Cmd {
			r.usernameInput.Focus()
			r.passwordInput.Blur()
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		).SetHandler(func(currentview view) tea.Cmd {
			r.usernameInput.Blur()
			r.passwordInput.Focus()
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "sign in / sign up"),
		).SetHandler(func(currentview view) tea.Cmd {
			goToView((viewSignup{}).Name())
			return nil
		}),
		NewKeyBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter"),
		).SetHandler(func(currentview view) tea.Cmd {
			requests[server.CmdSigninRequest](server.Data{
				Params: server.EventSigninParams{
					Username: r.usernameInput.Value(),
					Password: r.passwordInput.Value(),
				},
			})
			return nil
		}),
	}
}
