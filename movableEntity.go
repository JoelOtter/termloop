package termloop

// MovableEntity extends Entity to a moving object.
type MovableEntity struct {
	*Entity
	x         float64 // need higher precision for frame speed calculation
	y         float64
	direction Direction
	speed     float64
}

// NewMovableEntity creates a new MovableEntity, with position (x, y), size
// (width, height), direction and speed. Speed is defines as units per second.
// Returns a pointer to the new MovableEntity.
func NewMovableEntity(x, y, width, height int, direction Direction, speed float64) *MovableEntity {
	e := NewEntity(x, y, width, height)

	me := new(MovableEntity)
	me.Entity = e
	me.x = float64(x)
	me.y = float64(y)
	me.SetDirection(direction)
	me.SetSpeed(speed)

	return me
}

// NewMovableEntityFromCanvas returns a pointer to a new MovableEntity, with
// position (x, y), Canvas c, direction and speed. Width and height are calculated
// using the Canvas. Speed is defined as units per second.
func NewMovableEntityFromCanvas(x, y int, c Canvas, direction Direction, speed float64) *MovableEntity {
	e := NewEntityFromCanvas(x, y, c)
	me := new(MovableEntity)
	me.Entity = e
	me.x = float64(x)
	me.y = float64(y)
	me.SetDirection(direction)
	me.SetSpeed(speed)

	return me
}

// Draw updates x and y positions according to the speed and direction of the MovableEntity.
func (e *MovableEntity) Draw(s *Screen) {
	frameSpeed := e.speed / s.Fps()

	if e.direction & DirUp == DirUp {
		e.y -= frameSpeed
	}
	if e.direction & DirDown == DirDown {
		e.y += frameSpeed
	}
	if e.direction & DirLeft == DirLeft {
		e.x -= frameSpeed
	}
	if e.direction & DirRight == DirRight {
		e.x += frameSpeed
	}
	e.Entity.SetPosition(int(e.x), int(e.y))
	e.Entity.Draw(s)
}

// Speed returns the current speed of the MovableEntity.
func (e *MovableEntity) Speed() float64 {
	return e.speed
}

// SetSpeed sets the new speed of the MovableEntity.
func (e *MovableEntity) SetSpeed(speed float64) {
	e.speed = speed
}

// Direction returns the current direction of the MovableEntity
func (e *MovableEntity) Direction() Direction {
	return e.direction
}

// SetDirection sets the new direction of the MovableEntity
func (e *MovableEntity) SetDirection(direction Direction) {
	e.direction = direction
}
