package main

import (
	"fmt"
	tl "github.com/JoelOtter/termloop"
)

type EventInfo struct {
	*tl.Text
}

func NewEventInfo(x, y int) *EventInfo {
	return &EventInfo{tl.NewText(x, y, "Click somewhere", tl.ColorWhite, tl.ColorBlack)}
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
	info.SetText(fmt.Sprintf("%s @ [%d, %d]", name, ev.MouseX, ev.MouseY))
}

type Clickable struct {
	*tl.Rectangle
}

func NewClickable(x, y, w, h int, col tl.Attr) *Clickable {
	return &Clickable{tl.NewRectangle(x, y, w, h, col)}
}

func (c *Clickable) Tick(ev tl.Event) {
	x, y := c.Position()
	if ev.Type == tl.EventMouse && ev.MouseX == x && ev.MouseY == y {
		if c.Color() == tl.ColorWhite {
			c.SetColor(tl.ColorBlack)
		} else {
			c.SetColor(tl.ColorWhite)
		}
	}
}

func main() {
	g := tl.NewGame()
	g.Screen().SetFps(60)
	g.Screen().AddEntity(NewEventInfo(0, 0))
	for i := 0; i < 40; i++ {
		for j := 1; j < 20; j++ {
			g.Screen().AddEntity(NewClickable(i, j, 1, 1, tl.ColorWhite))
		}
	}

	g.Start()
}
