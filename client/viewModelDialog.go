package client

import (
	"fmt"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewModelSigninDialog struct {
	viewModelBase
	name         string
	Content      string
	OkButton     viewComponentButton
	CancelButton viewComponentButton
}

func (r *viewModelSigninDialog) updateCursorPos(msg tea.KeyMsg) {
	switch msg.Type {
	case tea.KeyLeft, tea.KeyRight, tea.KeyTab:
		r.OkButton.Active = !r.OkButton.Active
		r.CancelButton.Active = !r.CancelButton.Active
	default:
		return
	}
}

func (r *viewModelSigninDialog) submit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if r.OkButton.Active {
		return r.OkButton.Updater(msg)
	}
	if r.CancelButton.Active {
		return r.CancelButton.Updater(msg)
	}
	return nil, nil
}

func (r *viewModelSigninDialog) Init() tea.Cmd {
	return nil
}

func (r viewModelSigninDialog) Name() string {
	if r.name != "" {
		return fmt.Sprintf("%s_%s", reflect.TypeOf(r).Name(), r.name)
	}
	return fmt.Sprintf("%s_%s", reflect.TypeOf(r).Name(), r.Content)
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r *viewModelSigninDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeySpace:
			break
		case tea.KeyLeft, tea.KeyRight, tea.KeyTab:
			r.updateCursorPos(msg)
			break
		case tea.KeyEnter:
			return r.submit(msg)
		}
	}
	return r, nil
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (r viewModelSigninDialog) View() string {
	doc := strings.Builder{}

	content := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(r.Content)

	buttons := lipgloss.JoinHorizontal(lipgloss.Top, r.OkButton.View(), r.CancelButton.View())
	ui := lipgloss.JoinVertical(lipgloss.Center, content, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	doc.WriteString(dialog + "\n\n")

	return docStyle.Render(doc.String())
}
