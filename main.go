package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := new(Store)
	if err := store.Init(); err != nil {
		log.Fatalln(err)
	}

	m := NewModel(store)
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
