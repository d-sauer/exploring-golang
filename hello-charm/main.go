package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const size = 3

var (
	ledStyle = lipgloss.NewStyle().
		Width(5).
		Height(2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("69"))
)

type model struct {
	ledMatrix [][]string // led matrix
}

func initialModel() model {
	ledMatrix := make([][]string, size)
	for i := 0; i < size; i++ {
		ledMatrix[i] = make([]string, size)

		// set default color
		for n := 0; n < size; n++ {
			ledMatrix[i][n] = "0"
			//lipgloss.Color().RGBA()
		}
	}

	return model{
		ledMatrix: ledMatrix,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "0":
			m.ledMatrix[0][0] = "-"

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	s := "Led matrix\n\n"

	// Iterate over led matrix
	for y, ledY := range m.ledMatrix {
		var sx = make([]string, len(ledY))
		for x, led := range ledY {
			sx[x] = ledStyle.Render(fmt.Sprintf("[%d-%d](%s) ", x, y, led))
		}

		s += lipgloss.JoinHorizontal(lipgloss.Left, sx...)
		s += fmt.Sprintf("\n")
	}

	return s
}

func led() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63")).
		Border(lipgloss.DoubleBorder(), true, true, true, true).
		Width(1).
		Height(1)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
