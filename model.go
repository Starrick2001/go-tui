package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type AppModel struct {
	state     uint
	store     Store
	notes     []Note
	currNote  Note
	listIndex int
	textinput textinput.Model
	textarea  textarea.Model
}

func (a *AppModel) Init() tea.Cmd {
	return nil
}

func (a *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	a.textarea, cmd = a.textarea.Update(msg)
	cmds = append(cmds, cmd)

	a.textinput, cmd = a.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch a.state {
		case listView:
			switch key {
			case "ctrl+c", "q":
				return a, tea.Quit
			case "n":
				a.textinput.SetValue("")
				a.textinput.Focus()
				a.currNote = Note{}
				a.state = titleView
			case "up", "k":
				if a.listIndex > 0 {
					a.listIndex--
				}
			case "down", "j":
				if a.listIndex < len(a.notes)-1 {
					a.listIndex++
				}
			case "enter":
				a.currNote = a.notes[a.listIndex]
				a.state = bodyView
				a.textarea.SetValue(a.currNote.Body)
				a.textarea.Focus()
				a.textarea.CursorEnd()
			}
		case titleView:
			switch key {
			case "enter":
				title := a.textinput.Value()
				if title != "" {
					a.currNote.Title = title
					a.textarea.SetValue(a.currNote.Body)
					a.textarea.Focus()
					a.textarea.CursorEnd()

					a.state = bodyView
				}
			case "esc":
				a.state = listView
			}
		case bodyView:
			switch key {
			case "enter", "ctrl+s":
				a.currNote.Body = a.textarea.Value()
				var err error
				err = a.store.SaveNote(a.currNote)
				if err != nil {
					log.Println(err)
					return a, tea.Quit
				}

				a.notes, err = a.store.GetNotes()
				if err != nil {
					log.Println(err)
					return a, tea.Quit
				}

				a.currNote = Note{}
				a.state = listView
			case "esc":
				a.state = listView
			}
		}
	}
	return a, tea.Batch(cmds...)
}

func NewModel(store *Store) *AppModel {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalln(err)
	}

	return &AppModel{state: listView, store: *store, notes: notes, textinput: textinput.New(), textarea: textarea.New()}
}
