package box

import tl "github.com/badele/termloop"


////////////////////////////////////////////
// Horizontal Line
////////////////////////////////////////////

// A type representing a 2D horizontal Line
type HLine struct {
	x       int
	y       int
	size    int
	bgcolor tl.Attr
	fgcolor tl.Attr
	linetype LineType
	*BorderDefinitions
	level *tl.BaseLevel
}

// NewHLine creates a new HLine at position (x, y)
func NewHLine(x, y, s int, bgcolor, fgcolor tl.Attr, linetype LineType) *HLine {
	return &HLine{
		x: x, y: y, size: s,
		bgcolor: bgcolor, fgcolor: fgcolor, linetype: linetype, BorderDefinitions: BorderTheme[uint(linetype)],
		level: nil,
	}
}

// Draws the Horizontal line
func (l *HLine) Draw(s *tl.Screen) {
	posx, posy := l.x, l.y
	// If attached into level, no move text
	if l.level != nil {
		offSetX, offSetY := l.level.Offset()
		posx += -offSetX
		posy += -offSetY
	}

	for i := 0; i < l.size; i++ {
		s.RenderCell(posx+i, posy, &tl.Cell{Bg: l.bgcolor, Fg: l.fgcolor, Ch: l.BorderDefinitions.hc})
	}
}

func (l *HLine) Tick(ev tl.Event) {}

// Return size of horizontal line
func (l *HLine) Size() int {
	return l.size
}

// Position returns the x and y coordinates of the horizontal line.
func (l *HLine) Position() (int, int) {
	return l.x, l.y
}

// Level Follow
func (f *HLine) LevelFollow(level *tl.BaseLevel) {
	f.level = level
}


// SetPosition sets the coordinates of the horizontal line to be x and y.
func (l *HLine) SetPosition(x, y int) {
	l.x = x
	l.y = y
}

// SetSize sets the width and height of the horizontal line to be w and h.
func (l *HLine) SetSize(w, h int) {
	l.size = w
}

// Color returns the color of the horizontal line .
func (l *HLine) BgColor() tl.Attr {
	return l.bgcolor
}

// Color returns the color of the horizontal line .
func (l *HLine) FgColor() tl.Attr {
	return l.fgcolor
}


// SetBGColor set the background color of the horizontal line .
func (l *HLine) SetBGColor(color tl.Attr) {
	l.bgcolor = color
}

// SetFGColor set the background color of the horizontal line .
func (l *HLine) SetFGColor(color tl.Attr) {
	l.fgcolor = color
}

////////////////////////////////////////////
// Vertical Line
////////////////////////////////////////////

type VLine struct {
	x       int
	y       int
	size    int
	bgcolor tl.Attr
	fgcolor tl.Attr
	linetype LineType
	*BorderDefinitions
	level *tl.BaseLevel
}

// NewVLine creates a new VLine at position (x, y)
func NewVLine(x, y, s int, bgcolor, fgcolor tl.Attr, linetype LineType) *VLine {
	return &VLine{
		x: x, y: y, size: s,
		bgcolor: bgcolor, fgcolor: fgcolor, linetype: linetype, BorderDefinitions: BorderTheme[uint(linetype)],
		level: nil,
	}
}

// Draws the Horizontal line
func (l *VLine) Draw(s *tl.Screen) {
	posx, posy := l.x, l.y
	// If attached into level, no move text
	if l.level != nil {
		offSetX, offSetY := l.level.Offset()
		posx += -offSetX
		posy += -offSetY
	}


	for i := 0; i < l.size; i++ {
		s.RenderCell(posx, posy+i, &tl.Cell{Bg: l.bgcolor, Fg: l.fgcolor, Ch: l.BorderDefinitions.vc})
	}
}

func (l *VLine) Tick(ev tl.Event) {}

// Return size of vertical line
func (l *VLine) Size() int {
	return l.size
}

// Position returns the x and y coordinates of the horizontal line.
func (l *VLine) Position() (int, int) {
	return l.x, l.y
}

// Level Follow
func (f *VLine) LevelFollow(level *tl.BaseLevel) {
	f.level = level
}


// SetPosition sets the coordinates of the horizontal line to be x and y.
func (l *VLine) SetPosition(x, y int) {
	l.x = x
	l.y = y
}

// SetSize sets the height and height of the horizontal line to be w and h.
func (l *VLine) SetWidth(w, h int) {
	l.size = w
}

// Color returns the color of the horizontal line .
func (l *VLine) BgColor() tl.Attr {
	return l.bgcolor
}

// Color returns the color of the horizontal line .
func (l *VLine) FgColor() tl.Attr {
	return l.fgcolor
}


// SetBGColor set the background color of the horizontal line .
func (l *VLine) SetBGColor(color tl.Attr) {
	l.bgcolor = color
}

// SetFGColor set the background color of the horizontal line .
func (l *VLine) SetFGColor(color tl.Attr) {
	l.fgcolor = color
}
