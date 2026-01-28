package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// model defines the application's state.
type model struct {
	newFileInput textinput.Model
	createFileInputVisible bool
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
			return m, tea.Quit
	    case "ctrl+n":
			m.createFileInputVisible = true
		}
	}
	return m, nil
}

// View renders the UI.
func (m model) View() string {

	var style =lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		Margin(1, 2).
		Bold(true)

	welcome := style.Render("welcome to North Star!")

	//TODO: add dynamic view content here
    view := ""
	if m.createFileInputVisible {
		view += "Create New File:\n"
		view += m.newFileInput.View() + "\n"
	}

	//TODO: style the help text
	help := "Press q or ctrl+c to quit."

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome,view, help)
}

// initializeModel initializes the model with default values.
func initializeModel() model{

	//initialize new file input 
    ti := textinput.New()
	ti.Placeholder = "Enter new file name"
	ti.Focus()
	ti.CharLimit = 156

	return model{
		newFileInput: ti,
		createFileInputVisible: false,
	}
}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas there has been an error: %v\n", err)
		os.Exit(1)
	}
}
