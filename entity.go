package termloop

// Provides a general Drawable to be rendered.
type Entity struct {
	canvas Canvas
	x      int
	y      int
	width  int
	height int
}

// NewEntity creates a new Entity, with position (x, y) and size
// (width, height).
// Returns a pointer to the new Entity.
func NewEntity(x, y, width, height int) *Entity {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, height)
	}
	e := Entity{x: x, y: y, width: width, height: height,
		canvas: canvas}
	return &e
}

// Draw draws the entity to its current position on the screen.
// This is usually called every frame.
func (e *Entity) Draw(s *Screen) {
	for i := 0; i < e.width; i++ {
		for j := 0; j < e.height; j++ {
			s.RenderCell(e.x+i, e.y+j, &e.canvas[i][j])
		}
	}
}

func (e *Entity) Tick(ev Event) {}

// Position returns the (x, y) coordinates of the Entity.
func (e *Entity) Position() (int, int) {
	return e.x, e.y
}

// Size returns the width and height of the entity, in characters.
func (e *Entity) Size() (int, int) {
	return e.width, e.height
}

// SetPosition sets the x and y coordinates of the Entity.
func (e *Entity) SetPosition(x, y int) {
	e.x = x
	e.y = y
}

// SetCell updates the attribute of the Cell at x, y to match those of c.
// The coordinates are relative to the entity itself, not the Screen.
func (e *Entity) SetCell(x, y int, c *Cell) {
	renderCell(&e.canvas[x][y], c)
}
