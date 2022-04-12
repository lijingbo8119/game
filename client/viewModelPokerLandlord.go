package client

import (
	"fmt"
	"game/internal/room"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type viewModelPokerDoudizhu struct {
	viewModelBase
	room *room.RoomPokerDoudizhu
}

func (r viewModelPokerDoudizhu) Name() string {
	return reflect.TypeOf(r).Name()
}

func (r *viewModelPokerDoudizhu) SetRoom(ro room.Room) error {
	_room, ok := ro.(*room.RoomPokerDoudizhu)
	if !ok {
		return fmt.Errorf("room type error")
	}
	r.room = _room
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
