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
	for _, e := range l.entities {
		e.Tick(ev)
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
