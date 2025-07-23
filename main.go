package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github/SamedArslan28/statix/model"
	"log"
)

func main() {

	if _, err := tea.NewProgram(model.InitialModel()).Run(); err != nil {
		log.Fatal(err)
	}

}
