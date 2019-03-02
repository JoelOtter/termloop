package main

// A very simple image viewer, which uses Termloop's pixel mode

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
	"os"
)

type Image struct {
	e *tl.Entity
}

func NewImage(c *tl.Canvas) *Image {
	i := Image{e: tl.NewEntity(0, 0, len(*c), len((*c)[0]))}
	i.e.ApplyCanvas(c)
	return &i
}

func (i *Image) Draw(s *tl.Screen) { i.e.Draw(s) }

func (i *Image) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey {
		x, y := i.e.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			x -= 1
		case tl.KeyArrowLeft:
			x += 1
		case tl.KeyArrowUp:
			y += 1
		case tl.KeyArrowDown:
			y -= 1
		}
		i.e.SetPosition(x, y)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a filepath to open")
		return
	}

	g := tl.NewGame()
	g.Screen().SetFps(30)
	g.Screen().EnablePixelMode()
	c := tl.BackgroundCanvasFromFile(os.Args[1])
	g.Screen().AddEntity(NewImage(c))
	g.Start()
}
