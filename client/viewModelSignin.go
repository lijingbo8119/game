package client

import (
	"fmt"
	"game/server"
	"reflect"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type viewModelSignin struct {
	viewModelBase
	inited        bool
	cursorPos     int
	loading       bool
	tabs          viewComponentTabs
	nameInput     textinput.Model
	passwordInput textinput.Model
	statusBar     viewComponentStatusBar
}

func (r *viewModelSignin) updateCursorPos(msg tea.KeyMsg) {
	switch msg.Type {
	case tea.KeyUp:
		if r.cursorPos > 0 {
			r.cursorPos--
		}
	case tea.KeyDown:
		if r.cursorPos < 1 {
			r.cursorPos++
		}
	case tea.KeyTab:
		if r.cursorPos < 1 {
			r.cursorPos++
		} else {
			r.cursorPos = 0
		}
	default:
		return
	}
	r.nameInput.Blur()
	r.passwordInput.Blur()
	switch r.cursorPos {
	case 0:
		r.nameInput.Focus()
	case 1:
		r.passwordInput.Focus()
	}
}

func (r *viewModelSignin) FocusedInputUpdateNetCmd(msg tea.Msg) tea.Cmd {
	var (
		model textinput.Model
		cmd   tea.Cmd
	)
	switch r.cursorPos {
	case 0:
		model = r.nameInput
	case 1:
		model = r.passwordInput
	}
	model, cmd = model.Update(msg)
	switch r.cursorPos {
	case 0:
		r.nameInput = model
	case 1:
		r.passwordInput = model
	}
	return cmd
}

func (r *viewModelSignin) setStatusBarContent(s string) {
	r.loading = false
	r.statusBar.content = fmt.Sprintf(s)
}

func (r *viewModelSignin) submit() {
	r.setStatusBarContent("submitting")
	r.loading = true
	if err := requestSignin(server.EventSigninParams{
		Username: r.nameInput.Value(),
		Password: r.passwordInput.Value(),
	}); err != nil {
		r.setStatusBarContent(fmt.Sprintf("submit failed: %v", err))
	}
}

func (r *viewModelSignin) Init() tea.Cmd {
	if r.inited {
		return nil
	}
	r.inited = true

	r.tabs = newViewComponentTabs(
		viewComponentTab{Content: "Sign In", Active: true, Updater: func(msg tea.Msg) (tea.Model, tea.Cmd) {
			return getViewModel(viewModelSignup{}.Name()), nil
		}},
		viewComponentTab{Content: "Sign Up", Active: false, Updater: nil},
	)

	r.nameInput = textinput.New()
	r.nameInput.Placeholder = "Name"
	r.nameInput.Focus()

	r.passwordInput = textinput.New()
	r.passwordInput.Placeholder = "Password"

	r.statusBar = viewComponentStatusBar{
		locationName: "Sign In",
		ping:         1,
		content:      "",
	}
	return nil
}

func (r viewModelSignin) Name() string {
	return reflect.TypeOf(r).Name()
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r *viewModelSignin) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if getViewModel().Name() != r.Name() {
		return getViewModel().Update(msg)
	}
	if r.loading {
		return r, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			break
		case tea.KeyUp, tea.KeyDown:
			r.updateCursorPos(msg)
			break
		case tea.KeyTab:
			return r.tabs.Update(msg)
		case tea.KeyEnter:
			r.submit()
			return r, nil
		}
	}
	return r, r.FocusedInputUpdateNetCmd(msg)
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (r viewModelSignin) View() string {
	doc := strings.Builder{}

	doc.WriteString(r.tabs.View())
	doc.WriteString("\n\n")
	doc.WriteString(r.nameInput.View())
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
