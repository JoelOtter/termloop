package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	ent *tl.Entity
}

// Here, Draw simply tells the Entity ent to handle its own drawing.
// We don't need to do anything.
func (p *Player) Draw(s *tl.Screen) { p.ent.Draw(s) }

func (p *Player) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey { // Is it a keyboard event?
		x, y := p.ent.Position()
		switch ev.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			p.ent.SetPosition(x+1, y)
			break
		case tl.KeyArrowLeft:
			p.ent.SetPosition(x-1, y)
			break
		case tl.KeyArrowUp:
			p.ent.SetPosition(x, y-1)
			break
		case tl.KeyArrowDown:
			p.ent.SetPosition(x, y+1)
			break
		}
	}
}

func main() {
	g := tl.NewGame()
	l := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	l.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	p := Player{
		ent: tl.NewEntity(1, 1, 1, 1),
	}
	// Set the character at position (0, 0) on the entity.
	p.ent.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	l.AddEntity(&p)
	g.SetLevel(l)
	g.Start()
}
