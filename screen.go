package termloop

import "github.com/nsf/termbox-go"

type Canvas [][]Cell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, height)
	}
	return canvas
}

type Screen struct {
	canvas   Canvas
	level    Level
	entities []Drawable
	width    int
	height   int
	delta    float64
}

func NewScreen() *Screen {
	s := Screen{entities: make([]Drawable, 0)}
	return &s
}

func (s *Screen) Tick(ev Event) {
	// TODO implement ticks using worker pools
	if s.level != nil {
		s.level.Tick(ev)
	}
	for _, e := range s.entities {
		e.Tick(ev)
	}
}

func (s *Screen) Draw() {
	// Update termloop canvas
	s.canvas = NewCanvas(s.width, s.height)
	if s.level != nil {
		s.level.DrawBackground(s)
		s.level.Draw(s)
	}
	for _, e := range s.entities {
		e.Draw(s)
	}
	// Draw to terminal
	termbox.Clear(0, 0)
	for i, row := range s.canvas {
		for j, cell := range row {
			termbox.SetCell(i, j, cell.Ch,
				termbox.Attribute(cell.Fg),
				termbox.Attribute(cell.Bg))
		}
	}
	termbox.Flush()
}

func (s *Screen) Resize(w, h int) {
	s.width = w
	s.height = h
	c := NewCanvas(w, h)
	// Copy old data that fits
	for i := 0; i < Min(w, len(s.canvas)); i++ {
		for j := 0; j < Min(h, len(s.canvas[0])); j++ {
			c[i][j] = s.canvas[i][j]
		}
	}
	s.canvas = c
}

func (s *Screen) Size() (int, int) {
	return s.width, s.height
}

func (s *Screen) TimeDelta() float64 {
	return s.delta
}

func (s *Screen) RenderCell(x, y int, c *Cell) {
	RenderCell(&s.canvas[x][y], c)
}

func RenderCell(old, new_ *Cell) {
	if new_.Ch != 0 {
		old.Ch = new_.Ch
	}
	if new_.Bg != 0 {
		old.Bg = new_.Bg
	}
	if new_.Fg != 0 {
		old.Fg = new_.Fg
	}
}
