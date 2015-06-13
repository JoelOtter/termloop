package termloop

import "github.com/nsf/termbox-go"

type Entity struct {
	Canvas Canvas
	x      int
	y      int
	width  int
	height int
}

func NewEntity(x, y, width, height int) *Entity {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, height)
	}
	e := Entity{x: x, y: y, width: width, height: height,
		Canvas: canvas}
	return &e
}

func (e *Entity) Draw(s *Screen) {
	c := s.canvas
	for i := 0; i < Min(s.width-e.x, e.width); i++ {
		for j := 0; j < Min(s.height-e.y, e.height); j++ {
			if e.x+i >= 0 && e.y+j >= 0 {
				RenderCell(&c[e.x+i][e.y+j], &e.Canvas[i][j])
			}
		}
	}
}

func (e *Entity) Tick(ev termbox.Event) {}

func (e *Entity) Position() (int, int) {
	return e.x, e.y
}

func (e *Entity) SetPosition(x, y int) {
	e.x = x
	e.y = y
}
