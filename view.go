package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle           = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	faintStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	listEnumratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (a *AppModel) View() string {
	s := appStyle.Render("NOTES App") + "\n\n"
	if a.state == titleView {
		s += "Note title:\n\n"
		s += a.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - save • esc - discard")
	}

	if a.state == bodyView {
		s += "Note:\n\n"
		s += a.textarea.View() + "\n\n"
		s += faintStyle.Render("ctrl+s/enter - save • esc - discard")
	}
	if a.state == listView {
		for i, n := range a.notes {
			prefix := " "
			if i == a.listIndex {
				prefix = ">"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += listEnumratorStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}
		s += faintStyle.Render("n - new note, q/ctrl + c - quit")
	}

	return s
}
