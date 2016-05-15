package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

type EventInfo struct {
	text *tl.Text
}

func NewEventInfo(x, y int) *EventInfo {
	info := &EventInfo{}
	info.text = tl.NewText(x, y, "Click somewhere", tl.ColorWhite, tl.ColorBlack)
	return info
}

func (info *EventInfo) Draw(screen *tl.Screen) {
	info.text.Draw(screen)
}

func (info *EventInfo) Tick(ev tl.Event) {
	if ev.Type != tl.EventMouse {
		return
	}
	var name string
	switch ev.Key {
	case tl.MouseLeft:
		name = "Mouse Left"
	case tl.MouseMiddle:
		name = "Mouse Middle"
	case tl.MouseRight:
		name = "Mouse Right"
	case tl.MouseWheelUp:
		name = "Mouse Wheel Up"
	case tl.MouseWheelDown:
		name = "Mouse Wheel Down"
	case tl.MouseRelease:
		name = "Mouse Release"
	default:
		name = fmt.Sprintf("Unknown Key (%#x)", ev.Key)
	}
	info.text.SetText(fmt.Sprintf("%s @ [%d, %d]", name, ev.MouseX, ev.MouseY))
}

type Clickable struct {
	r *tl.Rectangle
}

func NewClickable(x, y, w, h int, col tl.Attr) *Clickable {
	return &Clickable{
		r: tl.NewRectangle(x, y, w, h, col),
	}
}

func (c *Clickable) Draw(s *tl.Screen) {
	c.r.Draw(s)
}

func (c *Clickable) Tick(ev tl.Event) {
	x, y := c.r.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		if c.r.Color() == tl.ColorWhite {
			c.r.SetColor(tl.ColorBlack)
		} else {
			c.r.SetColor(tl.ColorWhite)
		}
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().AddEntity(NewEventInfo(0, 0))
	for i := 0; i < 40; i++ {
		for j := 1; j < 20; j++ {
			g.Screen().AddEntity(NewClickable(i, j, 1, 1, tl.ColorWhite))
		}
	}

	g.Start()
}
