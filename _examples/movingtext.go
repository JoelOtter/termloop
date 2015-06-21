package main

import (
	"fmt"
	tl "github.com/joelotter/termloop"
	"os"
)

type MovingText struct {
	text *tl.Text
}

func (m *MovingText) Draw(s *tl.Screen) {
	m.text.Draw(s)
}

func (m *MovingText) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := m.text.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x += 1
			break
		case tl.KeyArrowLeft:
			x -= 1
			break
		case tl.KeyArrowUp:
			y -= 1
			break
		case tl.KeyArrowDown:
			y += 1
			break
		}
		m.text.SetPosition(x, y)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string as first argument")
		return
	}
	g := tl.NewGame()
	g.AddEntity(&MovingText{
		text: tl.NewText(0, 0, os.Args[1], tl.ColorWhite, tl.ColorBlue),
	})
	g.Start()
}
