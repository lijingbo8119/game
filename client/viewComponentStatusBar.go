package client

import (
	"fmt"
	"reflect"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewComponentStatusBar struct {
	viewComponentBase
	locationName string
	ping         int
	content      string
}

func (r viewComponentStatusBar) Name() string {
	return reflect.TypeOf(r).Name()
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r viewComponentStatusBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r viewComponentStatusBar) View() string {
	var (
		// Status Bar.
		statusNugget = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFDF5")).
				Padding(0, 1)

		statusBarStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
				Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

		statusStyle = lipgloss.NewStyle().
				Inherit(statusBarStyle).
				Foreground(lipgloss.Color("#FFFDF5")).
				Background(lipgloss.Color("#FF5F87")).
				Padding(0, 1).
				MarginRight(1)

		encodingStyle = statusNugget.Copy().
				Background(lipgloss.Color("#A550DF")).
				Align(lipgloss.Right)

		statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

		fishCakeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))
	)
	w := lipgloss.Width

	locationName := statusStyle.Render(r.locationName)
	encoding := encodingStyle.Render(time.Now().Format("2006-01-02 15:04:05"))
	fishCake := fishCakeStyle.Render(fmt.Sprintf("%s %d", "PING:", r.ping))
	statusVal := statusText.Copy().
		Width(viewStyle{}.PhysicalWidth() - w(locationName) - w(encoding) - w(fishCake)).
		Render(r.content)

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		locationName,
		statusVal,
		encoding,
		fishCake,
	)

	return statusBarStyle.Width(viewStyle{}.PhysicalWidth()).Render(bar)
}
