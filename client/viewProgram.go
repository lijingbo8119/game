package client

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func viewProgramStart() {
	viewModelsInit()

	// Initialize our program
	p := tea.NewProgram(currentViewModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
