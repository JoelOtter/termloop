package termloop

import "github.com/nsf/termbox-go"

type Entity struct {
	canvas Canvas
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
		canvas: canvas}
	return &e
}

func (e *Entity) Draw(s *Screen) {
	for i := 0; i < min(s.width-e.x, e.width); i++ {
		for j := 0; j < min(s.height-e.y, e.height); j++ {
			if e.x+i >= 0 && e.y+j >= 0 {
				s.RenderCell(e.x+i, e.y+j, &e.canvas[i][j])
			}
		}
	}
}

func (e *Entity) Tick(ev termbox.Event) {}

func (e *Entity) Position() (int, int) {
	return e.x, e.y
}

func (e *Entity) Size() (int, int) {
	return e.width, e.height
}

func (e *Entity) SetPosition(x, y int) {
	e.x = x
	e.y = y
}

func (e *Entity) SetCell(x, y int, c *Cell) {
	renderCell(&e.canvas[x][y], c)
}
