package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	ledStyle = lipgloss.NewStyle().
		Width(5).
		Height(2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("99"))
)

type model struct {
	ledMatrixColor [][]string // led matrix color
}

func initialModel(size int) model {
	ledMatrixColor := make([][]string, size)
	for i := 0; i < size; i++ {
		ledMatrixColor[i] = make([]string, size)

		// set default color
		for n := 0; n < size; n++ {
			ledMatrixColor[i][n] = "0"
		}
	}

	return model{
		ledMatrixColor: ledMatrixColor,
	}
}

func (m model) Init() tea.Cmd {
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
			m.ledMatrixColor[0][0] = RgbToColor(252, 3, 173)

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
	for y, ledY := range m.ledMatrixColor {
		var sx = make([]string, len(ledY))
		for x, _ := range ledY {
			sx[x] = ledColor(m.ledMatrixColor[y][x]).Render(fmt.Sprintf(""))
		}

		s += lipgloss.JoinHorizontal(lipgloss.Left, sx...)
		s += fmt.Sprintf("\n")
	}

	return s
}

func ledColor(color string) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(3).
		Height(1).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(-1).
		BorderForeground(lipgloss.Color("99")).
		Background(lipgloss.Color(color))
}

func RgbToColor(r int, g int, b int) string {
	c := NewRGBColor(r, g, b)
	return c.ToHex()
}

func main() {
	p := tea.NewProgram(initialModel(3))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
