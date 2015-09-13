package termloop

// Text represents a string that can be drawn to the screen.
type Text struct {
	x      int
	y      int
	fg     Attr
	bg     Attr
	text   []rune
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
	return &Text{
		x:      x,
		y:      y,
		fg:     fg,
		bg:     bg,
		text:   str,
		canvas: c,
	}
}

func (t *Text) Tick(ev Event) {}

// Draw draws the Text to the Screen s.
func (t *Text) Draw(s *Screen) {
	w, _ := t.Size()
	for i := 0; i < w; i++ {
		s.RenderCell(t.x+i, t.y, &t.canvas[i])
	}
}

// Position returns the (x, y) coordinates of the Text.
func (t *Text) Position() (int, int) {
	return t.x, t.y
}

// Size returns the width and height of the Text.
func (t *Text) Size() (int, int) {
	return len(t.text), 1
}

// SetPosition sets the coordinates of the Text to be (x, y).
func (t *Text) SetPosition(x, y int) {
	t.x = x
	t.y = y
}

// Text returns the text of the Text.
func (t *Text) Text() string {
	return string(t.text)
}

// SetText sets the text of the Text to be text.
func (t *Text) SetText(text string) {
	t.text = []rune(text)
	c := make([]Cell, len(t.text))
	for i := range c {
		c[i] = Cell{Ch: t.text[i], Fg: t.fg, Bg: t.bg}
	}
	t.canvas = c
}

// Color returns the (foreground, background) colors of the Text.
func (t *Text) Color() (Attr, Attr) {
	return t.fg, t.bg
}

// SetColor sets the (foreground, background) colors of the Text
// to fg, bg respectively.
func (t *Text) SetColor(fg, bg Attr) {
	t.fg = fg
	t.bg = bg
	for i := range t.canvas {
		t.canvas[i].Fg = fg
		t.canvas[i].Bg = bg
	}
}
