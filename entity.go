package termloop

// Entity provides a general Drawable to be rendered.
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
	canvas := NewCanvas(width, height)
	e := Entity{x: x, y: y, width: width, height: height,
		canvas: canvas}
	return &e
}

// NewEntityFromCanvas returns a pointer to a new Entity, with
// position (x, y) and Canvas c. Width and height are calculated
// using the Canvas.
func NewEntityFromCanvas(x, y int, c Canvas) *Entity {
	e := Entity{
		x:      x,
		y:      y,
		canvas: c,
		width:  len(c),
		height: len(c[0]),
	}
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

// Tick needs to be implemented to satisfy the Drawable interface.
// It updates the Entity based on the Screen's FPS
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

// Fill fills the canvas of the Entity with
// a Cell c.
func (e *Entity) Fill(c *Cell) {
	for i := range e.canvas {
		for j := range e.canvas[i] {
			renderCell(&e.canvas[i][j], c)
		}
	}
}

// ApplyCanvas takes a pointer to a Canvas, c, and applies this canvas
// over the top of the Entity's canvas. Any new values in c will overwrite
// those in the entity.
func (e *Entity) ApplyCanvas(c *Canvas) {
	for i := 0; i < min(len(e.canvas), len(*c)); i++ {
		for j := 0; j < min(len(e.canvas[0]), len((*c)[0])); j++ {
			renderCell(&e.canvas[i][j], &(*c)[i][j])
		}
	}
}

// SetCanvas takes a pointer to a Canvas and replaces the Entity's canvas with
// the pointer's. It also updates the Entity's dimensions.
func (e *Entity) SetCanvas(c *Canvas) {
	e.width = len(*c)
	e.height = len((*c)[0])
	e.canvas = *c
}
