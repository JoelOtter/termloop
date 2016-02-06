package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	entity *tl.Entity
	prevX  int
	prevY  int
	level  *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.entity.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.entity.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.entity.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.entity.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.entity.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.entity.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Size() (int, int)     { return player.entity.Size() }
func (player *Player) Position() (int, int) { return player.entity.Position() }

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.entity.SetPosition(player.prevX, player.prevY)
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	player := Player{
		entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)
	game.Screen().SetLevel(level)
	game.Start()
}
