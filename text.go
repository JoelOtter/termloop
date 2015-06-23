package termloop

// Text represents a string that can be drawn to the screen.
type Text struct {
	x      int
	y      int
	canvas []Cell
}

// NewText creates a new Text, at position (x, y). It sets the Text's
// background and foreground colors to fg and bg respectively, and sets the
// Text's text to be text.
// Returns a pointer to the new Text.
func NewText(x, y int, text string, fg, bg Attr) *Text {
	str := []rune(text)
	c := make([]Cell, len(str))
	for i := range c {
		c[i] = Cell{Ch: str[i], Fg: fg, Bg: bg}
	}
	return &Text{x: x, y: y, canvas: c}
}

func (t *Text) Tick(ev Event) {}

// Draw draws the Text to the Screen s.
func (t *Text) Draw(s *Screen) {
	for i := 0; i < min(s.width-t.x, len(t.canvas)); i++ {
		if t.x+i >= 0 && t.y >= 0 && t.y < s.height {
			s.RenderCell(t.x+i, t.y, &t.canvas[i])
		}
	}
}

// Position returns the (x, y) coordinates of the Text.
func (t *Text) Position() (int, int) {
	return t.x, t.y
}

// Size returns the width and height of the Text.
func (t *Text) Size() (int, int) {
	return len(t.canvas), 1
}

// SetPosition sets the coordinates of the Text to be (x, y).
func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}
