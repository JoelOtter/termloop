package termloop

// A type representing a 2D rectangle, with position, size and color.
type Rectangle struct {
	x      int
	y      int
	width  int
	height int
	color  Attr
}

// NewRectangle creates a new Rectangle at position (x, y), with size
// (width, height) and color color.
// Returns a pointer to the new Rectangle.
func NewRectangle(x, y, w, h int, color Attr) *Rectangle {
	r := Rectangle{x: x, y: y, width: w, height: h, color: color}
	return &r
}

// Draws the Rectangle r onto Screen s.
func (r *Rectangle) Draw(s *Screen) {
	for i := 0; i < r.width; i++ {
		for j := 0; j < r.height; j++ {
			s.RenderCell(r.x+i, r.y+j, &Cell{Bg: r.color, Ch: ' '})
		}
	}
}

func (r *Rectangle) Tick(ev Event) {}

// Size returns the width and height in characters of the Rectangle.
func (r *Rectangle) Size() (int, int) {
	return r.width, r.height
}

// Position returns the x and y coordinates of the Rectangle.
func (r *Rectangle) Position() (int, int) {
	return r.x, r.y
}

// SetPosition sets the coordinates of the Rectangle to be x and y.
func (r *Rectangle) SetPosition(x, y int) {
	r.x = x
	r.y = y
}

// SetSize sets the width and height of the Rectangle to be w and h.
func (r *Rectangle) SetSize(w, h int) {
	r.width = w
	r.height = h
}

// Color returns the color of the Rectangle.
func (r *Rectangle) Color() Attr {
	return r.color
}

// SetColor sets the color of the Rectangle.
func (r *Rectangle) SetColor(color Attr) {
	r.color = color
}
