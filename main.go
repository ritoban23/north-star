package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

// model defines the application's state.
type model struct {
	msg string 
}

// Init is called when the program starts.
func (m model) Init() tea.Cmd{
	return nil
}

// Update is called when a message is received.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){

	case tea.KeyMsg:
		switch msg.String(){
		case "ctrl+c", "q":
			fmt.Println("User clicked ctrl+c or q to quit")
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the UI.
func (m model) View() string {
	return m.msg
}

// initializeMode initializes the model with default values.
func initializeMode() model{
	return model{
		msg :"baingan",
	}
}

func main() {

	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Println("Alas ther has been a error: %v", err)
		os.Exit(1)
	}
	fmt.Println("hello")
}
