package box

import (
	tl "github.com/badele/termloop"
	"time"
)

// A type representing a Text with Area dimension
type TextArea struct {
	x,y,w,h     int
	align Align
	bgcolor tl.Attr
	fgcolor tl.Attr
	text string
	level *tl.BaseLevel
	previoustimewriter time.Time
	typewriterduration int64

}

// NewTextArea creates a new TextArea
func NewTextArea(x, y, w, h int, text string, bgcolor, fgcolor  tl.Attr, align Align) *TextArea {
	t := TextArea{x:x, y:y, w:w, h:h, fgcolor:fgcolor,bgcolor:bgcolor,level: nil,typewriterduration:0, previoustimewriter:time.Now()}
	return &t
}

// Draws the TextArea r onto Screen s.
func (t *TextArea) Draw(s *tl.Screen) {
	posx, posy := t.x, t.y

	text := tl.NewText(0, 0, t.text, t.fgcolor, t.bgcolor)
	if t.typewriterduration != 0 {
		now := time.Now()
		elapsed := int64(now.Sub(t.previoustimewriter)/time.Millisecond)
		nbchar := elapsed / t.typewriterduration
		if int(nbchar) < len(t.text) {
			text = tl.NewText(0, 0, t.text[0:nbchar], t.fgcolor, t.bgcolor)
		}
	}

	twidth,_ := text.Size()

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
	text.SetPosition(posx,posy)
	text.Draw(s)
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

// Set typewriter speed
func (t *TextArea) SetTypewriterDuration(typewriterduration int64) {
	t.previoustimewriter = time.Now()
	t.typewriterduration = typewriterduration
}

// Set Text
func (t *TextArea) SetText(text string, align Align) {
	t.SetAlign(align)
	t.text = text
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
	return t.bgcolor, t.fgcolor
}

func (t *TextArea) SetColor(bg, fg tl.Attr) {
	t.bgcolor, t.fgcolor = bg, fg
}
