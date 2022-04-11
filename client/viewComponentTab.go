package client

import (
	"fmt"
	"game/util"
	"reflect"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

var (
	tabStyle = lipgloss.NewStyle().
			Border(tabBorder, true).
			BorderForeground(highlight).
			Padding(0, 1)

	activeTabStyle = tabStyle.Copy().Border(activeTabBorder, true)

	tabStyleGap = tabStyle.Copy().
			BorderTop(false).
			BorderLeft(false).
			BorderRight(false)
)

type viewComponentTab struct {
	Content string
	Active  bool
	Updater viewModelUpdater
}

func (r viewComponentTab) View() string {
	if r.Active {
		return activeTabStyle.Render(r.Content)
	}
	return tabStyle.Render(r.Content)
}

type viewComponentTabs struct {
	viewComponentBase
	tabs []viewComponentTab
}

func (r viewComponentTabs) Name() string {
	return fmt.Sprintf("%s", reflect.TypeOf(r).Name())
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (r viewComponentTabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	tab, ok := lo.Find(r.tabs, func(t viewComponentTab) bool {
		return t.Active == true
	})
	if ok {
		return tab.Updater(msg)
	}
	return nil, nil
}

func (r viewComponentTabs) View() string {
	tabViews := []string{}
	lo.ForEach(r.tabs, func(t viewComponentTab, i int) {
		tabViews = append(tabViews, t.View())
	})
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabViews...,
	)
	gap := tabStyleGap.Render(strings.Repeat(" ", util.MaxInt(0, viewStyle{}.PhysicalWidth()-lipgloss.Width(row)-2)))
	return lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
}

func newViewComponentTabs(tabs ...viewComponentTab) viewComponentTabs {
	return viewComponentTabs{
		tabs: tabs,
	}
}
