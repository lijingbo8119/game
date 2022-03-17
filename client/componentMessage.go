package client

import "github.com/charmbracelet/lipgloss"

type cMessageType uint8

const (
	cMessageTWarning cMessageType = iota + 1
	cMessageTError
	cMessageTSuccess
	cMessageTInfo
)

type componentMessage struct {
	t       cMessageType
	message string
}

func (r *componentMessage) init() component {
	return r
}

func (r componentMessage) String() string {
	return ""
	var style = lipgloss.NewStyle().
		Width(24).
		Border(lipgloss.ThickBorder(), true, true).
		Bold(true).                        // make it bold
		UnsetBold().                       // jk don't make it bold
		Background(lipgloss.Color("227")). // yellow background
		UnsetBackground().                 // never mind
		Foreground(lipgloss.Color("63"))
	block := lipgloss.PlaceVertical(30, lipgloss.Bottom, style.Render("Hello, kitty."))
	return block
}
