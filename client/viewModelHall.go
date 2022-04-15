package client

import (
	"game/internal/game"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type viewModelHall struct {
	viewModelBase
	games []game.Game
}

func (r viewModelHall) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewModelHall) SetGames(games []game.Game) {
	r.games = games
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r viewModelHall) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (r viewModelHall) View() string {
	doc := strings.Builder{}

	doc.WriteString("viewModelHall")

	doc.WriteString("\n\n")
	// Status bar
	doc.WriteString(viewComponentStatusBar{}.View())

	return docStyle.MaxWidth(viewStyle{}.PhysicalWidth()).Render(doc.String())
}
