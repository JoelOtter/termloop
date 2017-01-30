package box

import tl "github.com/badele/termloop"

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
	*tl.Rectangle
	fgcolor   tl.Attr
	frametype LineType
	*BorderDefinitions
	signmode  bool
	*TextArea
}

// NewFrame creates a new Rectangle at position (x, y), with size
// (width, height) and color colof.
// Returns a pointer to the new NewFrame.
func NewFrame(x, y, w, h int, bgcolor tl.Attr, fgcolor tl.Attr, frametype LineType, signmode bool) *Frame {
	f := Frame{Rectangle: tl.NewRectangle(x, y, w, h, bgcolor),
		fgcolor: fgcolor,
		frametype: FrameEmpty,
		BorderDefinitions: BorderTheme[uint(frametype)],
		signmode: signmode,
		TextArea: NewTextArea(x, y, w, h,"",bgcolor, fgcolor, AlignCenter),
	}
	return &f
}

// Draws the Frame r onto Screen s.
func (f *Frame) Draw(s *tl.Screen) {

	// Draw Rectangle
	f.Rectangle.Draw(s)

	// Get Rectangle informations
	rx,ry := f.Rectangle.Position()
	sw,sh := f.Rectangle.Size()

	// Draw corners
	s.RenderCell(rx, ry, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.luc})
	s.RenderCell(rx+sw-1, ry, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.ruc})
	s.RenderCell(rx, ry+sh-1, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.lbc})
	s.RenderCell(rx+sw-1, ry+sh-1, &tl.Cell{Bg: f.BgColor(), Fg: f.FgColor(), Ch: f.BorderDefinitions.rbc})

	// Draw Horizontal line
	if !f.signmode {
		NewHLine(rx + 1, ry, f.width - 2, f.BgColor(), f.FgColor(), f.frametype).Draw(s)
		NewHLine(rx + 1, ry + sh - 1, f.width - 2, f.BgColor(), f.FgColor(), f.frametype).Draw(s)

		// Draw Horizontal line
		NewVLine(rx, ry + 1, f.height - 2, f.BgColor(), f.FgColor(), f.frametype).Draw(s)
		NewVLine(rx + f.width - 1, ry + 1, f.height - 2, f.BgColor(), f.FgColor(), f.frametype).Draw(s)
	}

	// Draw Text Area
	f.TextArea.Draw(s)


}

func (f *Frame) Tick(ev tl.Event) {}

// Size returns the width and height in characters of the Rectangle.
func (f *Frame) Size() (int, int) {
	return f.Rectangle.Size()
}

// Position returns the x and y coordinates of the Rectangle.
func (f *Frame) Position() (int, int) {
	return f.Rectangle.Position()
}

// Set Title
func (f *Frame) SetTitle(text string, align Align) {

	f.TextArea.SetText(text, align)
}


// SetPosition sets the coordinates of the Rectangle to be x and y.
func (f *Frame) SetPosition(x, y int) {
	f.Rectangle.SetPosition(x,y)
}

// SetSize sets the width and height of the Rectangle to be w and h.
func (f *Frame) SetSize(w, h int) {
	f.Rectangle.SetSize(w,h)
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
