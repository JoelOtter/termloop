package termloop

import "time"
import "github.com/nsf/termbox-go"

type Game struct {
	screen *Screen
	input  *input
}

func NewGame() *Game {
	g := Game{screen: NewScreen(), input: newInput()}
	return &g
}

func (g *Game) SetScreen(s *Screen) {
	g.screen = s
}

func (g *Game) CreateLevel(bg Attr) {
	g.screen.level = NewBaseLevel(Cell{Bg: bg})
}

func (g *Game) SetLevel(l Level) {
	g.screen.level = l
}

func (g *Game) AddEntity(d Drawable) {
	g.screen.entities = append(g.screen.entities, d)
}

func (g *Game) Start() {
	// Init Termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	g.screen.resize(termbox.Size())

	// Init input
	g.input.start()
	defer g.input.stop()
	clock := time.Now()

mainloop:
	for {
		select {
		case ev := <-g.input.eventQ:
			if ev.Key == g.input.endKey {
				break mainloop
			} else if EventType(ev.Type) == EventResize {
				g.screen.resize(ev.Width, ev.Height)
			}
			g.screen.Tick(convertEvent(ev))
		default:
			g.screen.Tick(Event{Type: EventNone})
		}
		update := time.Now()
		g.screen.delta = update.Sub(clock).Seconds()
		clock = update
		g.screen.Draw()
	}
}
