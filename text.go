package termloop

type Text struct {
	x      int
	y      int
	canvas []Cell
}

func NewText(x, y int, text string, fg, bg Attr) *Text {
	str := []rune(text)
	c := make([]Cell, len(str))
	for i := range c {
		c[i] = Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Text{x: x, y: y, canvas: c}
}

func (t *Text) Tick(ev Event) {}

func (t *Text) Draw(s *Screen) {
	for i := 0; i < Min(s.width-t.x, len(t.canvas)); i++ {
		if t.x+i >= 0 && t.y >= 0 && t.y < s.height {
			s.RenderCell(t.x+i, t.y, &t.canvas[i])
		}
	}
}

func (t *Text) Position() (int, int) {
	return t.x, t.y
}

func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}
