package box

import tl "github.com/badele/termloop"

// A type representing a Text with Area dimension
type TextArea struct {
	x,y,w,h     int
	align Align
	*tl.Text
	level *tl.BaseLevel

}

// NewTextArea creates a new TextArea
func NewTextArea(x, y, w, h int, text string, bgcolor, fgcolor  tl.Attr, align Align) *TextArea {
	t := TextArea{x:x, y:y, w:w, h:h, Text: tl.NewText(0, 0, "", fgcolor, bgcolor),level: nil}
	t.SetText(text,align)
	return &t
}

// Draws the TextArea r onto Screen s.
func (t *TextArea) Draw(s *tl.Screen) {
	posx, posy := t.x, t.y
	twidth,_ := t.Text.Size()

	if t.align&AlignRight == AlignRight{
		posx = t.x + (t.w - twidth)
	}

	if t.align&AlignHCenter == AlignHCenter{
		posx = t.x + ((t.w - twidth) / 2)
	}

	if t.align&AlignBottom == AlignBottom{
		posy = t.y + t.h
	}

	if t.align&AlignVCenter == AlignVCenter{
		posy = t.y + t.h /2
	}

	// If attached into level, no move text
	if t.level != nil {
		offSetX, offSetY := t.level.Offset()
		posx += -offSetX
		posy += -offSetY
	}
	t.Text.SetPosition(posx,posy)
	t.Text.Draw(s)
}

func (t *TextArea) Tick(ev tl.Event) {}

// Size returns the width and height in characters of the TextArea.
func (t *TextArea) Size() (int, int) {
	return t.w, t.h
}

// Position returns the x and y coordinates of the TextArea.
func (t *TextArea) Position() (int, int) {
	return t.x, t.y
}

// Set Text
func (t *TextArea) SetText(text string, align Align) {

	t.SetAlign(align)
	t.Text.SetText(text)
}


// SetPosition sets the coordinates of the TextArea to be x and y.
func (t *TextArea) SetAlign(align Align) {
	t.align = align
}

// SetPosition sets the coordinates of the TextArea to be x and y.
func (t *TextArea) SetPosition(x, y int) {
	t.x, t.y = x,y
}

// Level Follow
func (f *TextArea) LevelFollow(level *tl.BaseLevel) {
	f.level = level
}


// Color returns the color of the TextArea.
func (t *TextArea) Color() (tl.Attr, tl.Attr)  {
	return t.Text.Color()
}

func (t *TextArea) SetColor(fg, bg tl.Attr) {
	t.Text.SetColor(fg,bg)
}
