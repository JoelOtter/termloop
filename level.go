package termloop

// Level interface represents a Drawable with a separate background
// that is drawn first. It can also contain Drawables of its own.
type Level interface {
	DrawBackground(*Screen)
	AddEntity(Drawable)
	RemoveEntity(Drawable)
	Draw(*Screen)
	Tick(Event)
}

// BaseLevel type represents a Level with a background defined as a Cell,
// which is tiled. The background is drawn first, then all Entities.
type BaseLevel struct {
	Entities []Drawable
	bg       Cell
	offsetx  int
	offsety  int
}

// NewBaseLevel creates a new BaseLevel with background bg.
// Returns a pointer to the new BaseLevel.
func NewBaseLevel(bg Cell) *BaseLevel {
	level := BaseLevel{Entities: make([]Drawable, 0), bg: bg}
	return &level
}

// Tick handles any collisions between Physicals in the level's entities,
// and processes any input.
func (l *BaseLevel) Tick(ev Event) {
	// Handle input
	for _, e := range l.Entities {
		e.Tick(ev)
	}

	// Handle collisions
	colls := make([]Physical, 0)
	dynamics := make([]DynamicPhysical, 0)
	for _, e := range l.Entities {
		if p, ok := interface{}(e).(Physical); ok {
			colls = append(colls, p)
		}
		if p, ok := interface{}(e).(DynamicPhysical); ok {
			dynamics = append(dynamics, p)
		}

	}
	jobs := make(chan DynamicPhysical, len(dynamics))
	results := make(chan int, len(dynamics))
	for w := 0; w <= len(dynamics)/3; w++ {
		go checkCollisionsWorker(colls, jobs, results)
	}
	for _, p := range dynamics {
		jobs <- p
	}
	close(jobs)
	for r := 0; r < len(dynamics); r++ {
		<-results
	}
}

// DrawBackground draws the background Cell bg to each Cell of the Screen s.
func (l *BaseLevel) DrawBackground(s *Screen) {
	for i, row := range s.canvas {
		for j := range row {
			s.canvas[i][j] = l.bg
		}
	}
}

// Draw draws the level's entities to the Screen s.
func (l *BaseLevel) Draw(s *Screen) {
	offx, offy := s.offset()
	s.setOffset(l.offsetx, l.offsety)
	for _, e := range l.Entities {
		e.Draw(s)
	}
	s.setOffset(offx, offy)
}

// AddEntity adds Drawable d to the level's entities.
func (l *BaseLevel) AddEntity(d Drawable) {
	l.Entities = append(l.Entities, d)
}

// RemoveEntity removes Drawable d from the level's entities.
func (l *BaseLevel) RemoveEntity(d Drawable) {
	for i, elem := range l.Entities {
		if elem == d {
			l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
			return
		}
	}
}

// Offset returns the level's drawing offset.
func (l *BaseLevel) Offset() (int, int) {
	return l.offsetx, l.offsety
}

// SetOffset sets the level's drawing offset to be (x, y).
// The drawing offset can be used to simulate moving the level, or
// moving the 'camera'.
func (l *BaseLevel) SetOffset(x, y int) {
	l.offsetx, l.offsety = x, y
}

func checkCollisionsWorker(ps []Physical, jobs <-chan DynamicPhysical, results chan<- int) {
	for p := range jobs {
		for _, c := range ps {
			if c == p {
				continue
			}
			px, py := p.Position()
			cx, cy := c.Position()
			pw, ph := p.Size()
			cw, ch := c.Size()
			if px < cx+cw && px+pw > cx &&
				py < cy+ch && py+ph > cy {
				p.Collide(c)
			}
		}
		results <- 1
	}
}
