package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	ent   *tl.Entity
	px    int
	py    int
	level *tl.BaseLevel
}

func (p *Player) Draw(s *tl.Screen) {
	sw, sh := s.Size()
	x, y := p.ent.Position()
	p.level.SetOffset(sw/2-x, sh/2-y)
	p.ent.Draw(s)
}

func (p *Player) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey { // Is it a keyboard event?
		p.px, p.py = p.ent.Position()
		switch ev.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			p.ent.SetPosition(p.px+1, p.py)
			break
		case tl.KeyArrowLeft:
			p.ent.SetPosition(p.px-1, p.py)
			break
		case tl.KeyArrowUp:
			p.ent.SetPosition(p.px, p.py-1)
			break
		case tl.KeyArrowDown:
			p.ent.SetPosition(p.px, p.py+1)
			break
		}
	}
}

func (p *Player) Size() (int, int)     { return p.ent.Size() }
func (p *Player) Position() (int, int) { return p.ent.Position() }

func (p *Player) Collide(c tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := c.(*tl.Rectangle); ok {
		p.ent.SetPosition(p.px, p.py)
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
		ent:   tl.NewEntity(1, 1, 1, 1),
		level: l,
	}
	// Set the character at position (0, 0) on the entity.
	p.ent.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	l.AddEntity(&p)
	g.SetLevel(l)
	g.Start()
}
