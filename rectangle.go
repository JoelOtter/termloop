package termloop

import "github.com/nsf/termbox-go"

type Rectangle struct {
	x      int
	y      int
	width  int
	height int
	color  Attr
}

func NewRectangle(x, y, w, h int, color Attr) *Rectangle {
	r := Rectangle{x: x, y: y, width: w, height: h, color: color}
	return &r
}

func (r *Rectangle) Draw(s *Screen) {
	for i := 0; i < Min(r.width, s.width-r.x); i++ {
		for j := 0; j < Min(r.height, s.height-r.y); j++ {
			if r.x+i >= 0 && r.y+j >= 0 {
				s.canvas[r.x+i][r.y+j] = Cell{Bg: r.color}
			}

		}
	}
}

func (r *Rectangle) Tick(ev termbox.Event) {}
