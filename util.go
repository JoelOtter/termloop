package termloop

import "strconv"

// Utility types and associated methods

// FpsText provides a Text which updates with the current 'framerate'
// at specified intervals, to be used for testing performance.
// Please note that the framerate displayed is a measure of Termloop's
// processing speed - visible framerate is largely dependent on your terminal!
type FpsText struct {
	text   *Text
	time   float64
	update float64
}

// NewFpsText creates a new FpsText at position (x, y) and with background
// and foreground colors fg and bg respectively. It will refresh every
// 'update' seconds.
// Returns a pointer to the new FpsText.
func NewFpsText(x, y int, fg, bg Attr, update float64) *FpsText {
	return &FpsText{
		text:   NewText(x, y, "", fg, bg),
		time:   0,
		update: update,
	}
}

func (f *FpsText) Tick(ev Event) {}

// Draw updates the framerate on the FpsText and draws it to the Screen s.
func (f *FpsText) Draw(s *Screen) {
	f.time += s.TimeDelta()
	if f.time > f.update {
		fps := strconv.FormatFloat(1.0/s.TimeDelta(), 'f', 10, 64)
		f.text.SetText(fps)
		f.time -= f.update
	}
	f.text.Draw(s)
}
