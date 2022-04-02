package client

import (
	"sync"

	"github.com/samber/lo"
)

var viewModelsMux sync.RWMutex

var viewModels []viewModel = []viewModel{
	&viewModelSignin{},
	&viewModelSignup{},
	&viewModelHall{},
}

var currentviewModelName string

func currentViewModel(name ...string) viewModel {
	viewModelsMux.RLock()
	defer viewModelsMux.RUnlock()
	currentviewModelName = func() string {
		if len(name) > 0 {
			return name[0]
		}
		return currentviewModelName
	}()
	if m, ok := lo.Find(viewModels, func(t viewModel) bool {
		return t.Name() == currentviewModelName
	}); ok {
		m.Init()
		return m
	}
	return nil
}

func viewModelsInit() {
	if currentviewModelName == "" {
		currentviewModelName = viewModelSignup{}.Name()
	}
}
