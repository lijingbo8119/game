package client

import (
	"fmt"
	"game/internal/game"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type viewModelPokerDoudizhu struct {
	viewModelBase
	game *game.GamePokerDoudizhu
}

func (r viewModelPokerDoudizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewModelPokerDoudizhu) SetGame(g game.Game) error {
	_game, ok := g.(*game.GamePokerDoudizhu)
	if !ok {
		return fmt.Errorf("game type error")
	}
	r.game = _game
	return nil
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r viewModelPokerDoudizhu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

// Views return a string based on data in the model. That string which will be
// rendered to the terminal.
func (r viewModelPokerDoudizhu) View() string {
	doc := strings.Builder{}

	doc.WriteString("viewModelPokerDoudizhu")

	doc.WriteString("\n\n")
	// Status bar
	doc.WriteString(viewComponentStatusBar{}.View())

	return docStyle.MaxWidth(viewStyle{}.PhysicalWidth()).Render(doc.String())
}
