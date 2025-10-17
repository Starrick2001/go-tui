package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(NewModel())
	_, err := p.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

type AppModel struct {
	title     string
	textInput textinput.Model
}

func (a *AppModel) Init() tea.Cmd {
	//
	return textinput.Blink
}

func (a *AppModel) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return a, nil
}

func (a *AppModel) View() string {
	return a.textInput.View()
}

func NewModel() tea.Model {
	tI := textinput.New()
	tI.Placeholder = "Enter text"
	tI.Focus()
	return &AppModel{title: "Title", textInput: textinput.New()}
}
