package box

import tl "github.com/badele/termloop"

// A type representing a Text with Area dimension
type TextArea struct {
	*tl.Text
	x int
	y int
	width int
	height int
	align Align
}

// NewTextArea creates a new TextArea
func NewTextArea(x, y, w, h int, text string, bgcolor, fgcolor  tl.Attr, align Align) *TextArea {
	t := TextArea{x:x, y:y, width:w, height:h, Text: tl.NewText(0, 0, "", fgcolor, bgcolor), }
	t.SetText(text,align)
	return &t
}

// Draws the TextArea r onto Screen s.
func (t *TextArea) Draw(s *tl.Screen) {
	// Draw TextArea
	t.Text.Draw(s)
}

func (t *TextArea) Tick(ev tl.Event) {}

// Size returns the width and height in characters of the TextArea.
func (t *TextArea) Size() (int, int) {
	return t.Text.Size()
}

// Position returns the x and y coordinates of the TextArea.
func (t *TextArea) Position() (int, int) {
	return t.Text.Position()
}

// Set Text
func (t *TextArea) SetText(text string, align Align) {

	t.SetAlign(align)
	t.Text.SetText(text)

	// Get TextArea informations
	twidth,_ := t.Text.Size()

	// Default value for Align = AlignNone | AlignLeft | AlignTop
	posx := t.x
	posy := t.y


	if t.align&AlignRight == AlignRight{
		posx = t.x + (t.width - twidth)
	}

	if t.align&AlignHCenter == AlignHCenter{
		posx = t.x + ((t.width - twidth) / 2)
	}

	if t.align&AlignBottom == AlignBottom{
		posy = t.y + t.height
	}

	if t.align&AlignVCenter == AlignVCenter{
		posy = t.y + t.height /2
	}

	t.Text.SetPosition(posx,posy)
}


// SetPosition sets the coordinates of the TextArea to be x and y.
func (t *TextArea) SetAlign(align Align) {
	t.align = align
}

// SetPosition sets the coordinates of the TextArea to be x and y.
func (t *TextArea) SetPosition(x, y int) {
	t.Text.SetPosition(x,y)
}

// Color returns the color of the TextArea.
func (t *TextArea) Color() (tl.Attr, tl.Attr)  {
	return t.Text.Color()
}

func (t *TextArea) SetColor(fg, bg tl.Attr) {
	t.Text.SetColor(fg,bg)
}
