package main

import tl "github.com/joelotter/termloop"

type CollRec struct {
	r    *tl.Rectangle
	move bool
	px   int
	py   int
}

// Implement necessary methods for Physical interface
func (r *CollRec) Draw(s *tl.Screen)    { r.r.Draw(s) }
func (r *CollRec) Size() (int, int)     { return r.r.Size() }
func (r *CollRec) Position() (int, int) { return r.r.Position() }

func (r *CollRec) Tick(ev tl.Event) {
	// Enable arrow key movement
	if ev.Type == tl.EventKey && r.move {
		r.px, r.py = r.r.Position()
		switch ev.Key {
		case tl.KeyArrowRight:
			r.r.SetPosition(r.px+1, r.py)
			break
		case tl.KeyArrowLeft:
			r.r.SetPosition(r.px-1, r.py)
			break
		case tl.KeyArrowUp:
			r.r.SetPosition(r.px, r.py-1)
			break
		case tl.KeyArrowDown:
			r.r.SetPosition(r.px, r.py+1)
			break
		}
	}
}

func (r *CollRec) Collide(p tl.Physical) {
	// Check if it's a CollRec we're colliding with
	if _, ok := p.(*CollRec); ok && r.move {
		r.r.SetColor(tl.ColorBlue)
		r.r.SetPosition(r.px, r.py)
	}
}

func main() {
	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
	})
	l.AddEntity(&CollRec{
		r:    tl.NewRectangle(3, 3, 3, 3, tl.ColorRed),
		move: true,
	})
	l.AddEntity(&CollRec{
		r:    tl.NewRectangle(7, 4, 3, 3, tl.ColorGreen),
		move: false,
	})
	g.SetLevel(l)
	g.AddEntity(tl.NewFpsText(0, 0, tl.ColorRed, tl.ColorDefault, 0.5))
	g.Start()
}
