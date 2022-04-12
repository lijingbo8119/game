package client

import (
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/samber/lo"
)

type viewModel interface {
	tea.Model
	Name() string
}

type viewModelBase struct{}

func (r viewModelBase) Init() tea.Cmd {
	return nil
}

type viewModelUpdater func(msg tea.Msg) (tea.Model, tea.Cmd)

var (
	_viewModelsMux sync.RWMutex

	viewModels []viewModel = []viewModel{
		&viewModelSignin{},
		&viewModelSignup{},
		&viewModelHall{},
		&viewModelPokerDoudizhu{},
	}

	_currentviewModelName string
)

func getViewModel(name ...string) viewModel {
	_viewModelsMux.RLock()
	defer _viewModelsMux.RUnlock()
	_currentviewModelName = func() string {
		if len(name) > 0 {
			return name[0]
		}
		return _currentviewModelName
	}()
	if m, ok := lo.Find(viewModels, func(t viewModel) bool {
		return t.Name() == _currentviewModelName
	}); ok {
		m.Init()
		return m
	}
	return nil
}

func viewModelsInit() {
	if _currentviewModelName == "" {
		_currentviewModelName = viewModelSignin{}.Name()
	}
}
