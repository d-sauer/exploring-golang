package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlue)
	s.SetStyle(defStyle)

	s.SetContent(0, 0, 'H', nil, defStyle)
	s.SetContent(1, 0, 'i', nil, defStyle)
	s.SetContent(2, 0, '!', nil, defStyle)
	s.Show()
	time.Sleep(2 * time.Second)

	s.SetContent(0, 0, 'H', nil, tcell.StyleDefault.Foreground(tcell.ColorRed))
	s.Show()
	time.Sleep(3 * time.Second)

	// Clear screen
	s.Clear()
}
