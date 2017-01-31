package box

import (
	tl "github.com/badele/termloop"
)

type BorderDefinitions struct {
	hc rune
	vc rune
	luc rune
	ruc rune
	lbc rune
	rbc rune
}

// A type representing a 2D rectangle, with position, size and colof.
type Frame struct {
	x,y,w,h int
	*tl.Rectangle
	fgcolor   tl.Attr
	frametype LineType
	*BorderDefinitions
	signmode  bool
	*TextArea
	level *tl.BaseLevel
}

// NewFrame creates a new Rectangle at position (x, y), with size
// (width, height) and color colof.
// Returns a pointer to the new NewFrame.
func NewFrame(x, y, w, h int, bgcolor tl.Attr, fgcolor tl.Attr, frametype LineType, signmode bool) *Frame {
	f := Frame{x: x, y: y, w: w, h:h,
		fgcolor: fgcolor,
		frametype: frametype,
		BorderDefinitions: BorderTheme[uint(frametype)],
		signmode: signmode,
		Rectangle: tl.NewRectangle(x, y, w, h, bgcolor),
		TextArea: NewTextArea(x, y, w, h,"",bgcolor, fgcolor, AlignCenter),
		level: nil,
	}
	return &f
}

// Draws the Frame r onto Screen s.
func (f *Frame) Draw(s *tl.Screen) {

	posx, posy := f.x, f.y
	sw,sh := f.w, f.h

	// If attached into level, no move box
	if f.level != nil {
		offSetX, offSetY := f.level.Offset()
		posx += -offSetX
		posy += -offSetY
	}

	// Draw Rectangle
	f.Rectangle.SetPosition(posx, posy)
	f.Rectangle.Draw(s)

	// Draw corners
	s.RenderCell(posx, posy, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.luc})
	s.RenderCell(posx +sw-1, posy, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.ruc})
	s.RenderCell(posx, posy +sh-1, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.lbc})
	s.RenderCell(posx +sw-1, posy +sh-1, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.rbc})

	// Draw Horizontal line
	if !f.signmode {
		hline := NewHLine(f.x + 1, f.y, f.w - 2, f.BgColor(), f.FgColor(), f.frametype)
		hline.LevelFollow(f.level)
		hline.Draw(s)

		hline = NewHLine(f.x + 1, f.y + sh - 1, f.w - 2, f.BgColor(), f.FgColor(), f.frametype)
		hline.LevelFollow(f.level)
		hline.Draw(s)

		// Draw Horizontal line
		vhline := NewVLine(f.x, f.y + 1, f.h - 2, f.BgColor(), f.FgColor(), f.frametype)
		vhline.LevelFollow(f.level)
		vhline.Draw(s)

		vhline = NewVLine(f.x + f.w - 1, f.y + 1, f.h - 2, f.BgColor(), f.FgColor(), f.frametype)
		vhline.LevelFollow(f.level)
		vhline.Draw(s)
	}

	// Draw Text Area
	f.TextArea.Draw(s)


}

func (f *Frame) Tick(ev tl.Event) {}

// Size returns the width and height in characters of the Rectangle.
func (f *Frame) Size() (int, int) {
	return f.w, f.h
}

// Position returns the x and y coordinates of the Rectangle.
func (f *Frame) Position() (int, int) {
	return f.x, f.y
}

// Level follow
func (f *Frame) LevelFollow(level *tl.BaseLevel) {
	f.level = level
	f.TextArea.level = level
}

// Set Title
func (f *Frame) SetTitle(text string, align Align) {

	f.TextArea.SetText(text, align)
}


// SetPosition sets the coordinates of the Rectangle to be x and y.
func (f *Frame) SetPosition(x, y int) {
	f.x, f.y = x,y
}

// SetSize sets the width and height of the Rectangle to be w and h.
func (f *Frame) SetSize(w, h int) {
	f.w, f.h = w, h
}

// Color returns the color of the Rectangle.
func (f *Frame) BgColor() tl.Attr {
	return f.Rectangle.Color()
}

func (f *Frame) FgColor() tl.Attr {
	return f.fgcolor
}


// SetColor sets the color of the Rectangle.
func (f *Frame) SetBgColor(color tl.Attr) {
	f.Rectangle.SetColor(color)
}

func (f *Frame) SetFgColor(color tl.Attr) {
	f.fgcolor = color
}
