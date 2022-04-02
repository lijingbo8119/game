package client

import (
	"fmt"
	"game/server"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewModelSignup struct {
	viewModelBase
	inited        bool
	cursorPos     int
	loading       bool
	nameInput     textinput.Model
	nicknameInput textinput.Model
	passwordInput textinput.Model
	statusBar     viewComponentStatusBar
}

func (r *viewModelSignup) updateCursorPos(msg tea.KeyMsg) {
	switch msg.Type {
	case tea.KeyUp:
		if r.cursorPos > 0 {
			r.cursorPos--
		}
	case tea.KeyDown:
		if r.cursorPos < 2 {
			r.cursorPos++
		}
	default:
		return
	}
	r.nameInput.Blur()
	r.nicknameInput.Blur()
	r.passwordInput.Blur()
	switch r.cursorPos {
	case 0:
		r.nameInput.Focus()
	case 1:
		r.nicknameInput.Focus()
	case 2:
		r.passwordInput.Focus()
	}
}

func (r *viewModelSignup) FocusedInputUpdateCmd(msg tea.Msg) tea.Cmd {
	var (
		model textinput.Model
		cmd   tea.Cmd
	)
	switch r.cursorPos {
	case 0:
		model = r.nameInput
	case 1:
		model = r.nicknameInput
	case 2:
		model = r.passwordInput
	}
	model, cmd = model.Update(msg)
	switch r.cursorPos {
	case 0:
		r.nameInput = model
	case 1:
		r.nicknameInput = model
	case 2:
		r.passwordInput = model
	}
	return cmd
}

func (r *viewModelSignup) setStatusBarContent(s string) {
	r.loading = false
	r.statusBar.content = fmt.Sprintf(s)
}

func (r *viewModelSignup) submit() {
	r.setStatusBarContent("submitting")
	r.loading = true
	if err := requestSignup(server.EventSignupParams{
		Username: r.nameInput.Value(),
		Password: r.passwordInput.Value(),
		Nickname: r.nicknameInput.Value(),
	}); err != nil {
		r.setStatusBarContent(fmt.Sprintf("submit failed: %v", err))
	}
}

func (r *viewModelSignup) Init() tea.Cmd {
	if r.inited {
		return nil
	}
	r.inited = true

	r.nameInput = textinput.New()
	r.nameInput.Placeholder = "Name"
	r.nameInput.Focus()

	r.nicknameInput = textinput.New()
	r.nicknameInput.Placeholder = "Nickname"

	r.passwordInput = textinput.New()
	r.passwordInput.Placeholder = "Password"

	r.statusBar = viewComponentStatusBar{
		locationName: "Sign Up",
		ping:         1,
		content:      "",
	}
	return nil
}

func (r viewModelSignup) Name() string {
	return reflect.TypeOf(r).Name()
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r *viewModelSignup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if currentViewModel().Name() != r.Name() {
		return currentViewModel().Update(msg)
	}
	if r.loading {
		return r, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp, tea.KeyDown:
			r.updateCursorPos(msg)
		case tea.KeyEnter:
			r.submit()
			return r, nil
		}
	}
	return r, r.FocusedInputUpdateCmd(msg)
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (r viewModelSignup) View() string {
	doc := strings.Builder{}

	doc.WriteString("viewModelSignup" + time.Now().String())

	doc.WriteString("\n\n")
	doc.WriteString(r.nameInput.View())
	doc.WriteString("\n\n")
	doc.WriteString(r.nicknameInput.View())
	doc.WriteString("\n\n")
	doc.WriteString(r.passwordInput.View())
	doc.WriteString("\n\n")
	// Status bar
	doc.WriteString(r.statusBar.View())

	if (viewStyle{}.PhysicalWidth()) > 0 {
		docStyle = docStyle.MaxWidth(viewStyle{}.PhysicalWidth())
	}

	return docStyle.Render(doc.String())
}
