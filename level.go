package termloop

type Level interface {
	DrawBackground(*Screen)
	Draw(*Screen)
	Tick(Event)
}

type BaseLevel struct {
	entities []Drawable
	bg       Cell
}

func NewBaseLevel(bg Cell) *BaseLevel {
	level := BaseLevel{entities: make([]Drawable, 0), bg: bg}
	return &level
}

func (l *BaseLevel) Tick(ev Event) {
	// Handle collisions
	colls := make([]Physical, 0)
	for _, e := range l.entities {
		if p, ok := interface{}(e).(Physical); ok {
			colls = append(colls, p)
		}
	}
	jobs := make(chan Physical, len(colls))
	results := make(chan int, len(colls))
	for w := 0; w <= len(colls)/3; w++ {
		go checkCollisionsWorker(colls, jobs, results)
	}
	for _, p := range colls {
		jobs <- p
	}
	close(jobs)
	for r := 0; r < len(colls); r++ {
		<-results
	}

	// Handle input
	if ev.Type != EventNone {
		for _, e := range l.entities {
			e.Tick(ev)
		}
	}
}

func (l *BaseLevel) DrawBackground(s *Screen) {
	for i, row := range s.canvas {
		for j, _ := range row {
			s.canvas[i][j] = l.bg
		}
	}
}

func (l *BaseLevel) Draw(s *Screen) {
	for _, e := range l.entities {
		e.Draw(s)
	}
}

func (l *BaseLevel) AddEntity(d Drawable) {
	l.entities = append(l.entities, d)
}

func (l *BaseLevel) RemoveEntity(d Drawable) {
	for i, elem := range l.entities {
		if elem == d {
			l.entities = append(l.entities[:i], l.entities[i+1:]...)
			return
		}
	}
}

func checkCollisionsWorker(ps []Physical, jobs <-chan Physical, results chan<- int) {
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
				c.Collide(p)
				p.Collide(c)
			}
		}
		results <- 1
	}
}
