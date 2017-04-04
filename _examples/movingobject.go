package main

import tl "github.com/JoelOtter/termloop"

// Simple demo of a MovableEntity

// A block moving around
type movingBlock struct {
	*tl.MovableEntity
	steps int
}

// We code a simple pattern that repeats every 160 frames
func (mb *movingBlock) Draw(s *tl.Screen) {
	if mb.steps == 160 {
		mb.SetDirection(tl.DirDown)
	}
	if mb.steps == 120 {
		mb.SetDirection(tl.DirRight | tl.DirUp)
	}
	if mb.steps == 80 {
		mb.SetDirection(tl.DirDown)
	}
	if mb.steps == 40 {
		mb.SetDirection(tl.DirLeft | tl.DirUp)
	}
	mb.steps--
	if mb.steps <= 0 {
		mb.steps = 160
	}
	mb.MovableEntity.Draw(s)
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})

	// Make a moving block with a speed of 4 units per second
	block := new(movingBlock)
	block.MovableEntity = tl.NewMovableEntity(20, 10, 1, 1, tl.DirDown, 4.0)
	block.MovableEntity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '\u2B1C'})

	level.AddEntity(block)

	game.Screen().SetLevel(level)
	// Don't forget to set a reasonable fps.
	game.Screen().SetFps(30)
	game.Start()
}
