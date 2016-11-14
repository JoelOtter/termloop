package termloop

import "strconv"

// Utility types and associated methods

// FpsText provides a Text which updates with the current 'framerate'
// at specified intervals, to be used for testing performance.
// Please note that the framerate displayed is a measure of Termloop's
// processing speed - visible framerate is largely dependent on your terminal!
type FpsText struct {
	*Text
	time   float64
	update float64
}

// NewFpsText creates a new FpsText at position (x, y) and with background
// and foreground colors fg and bg respectively. It will refresh every
// 'update' seconds.
// Returns a pointer to the new FpsText.
func NewFpsText(x, y int, fg, bg Attr, update float64) *FpsText {
	return &FpsText{
		Text:   NewText(x, y, "", fg, bg),
		time:   0,
		update: update,
	}
}

// Draw updates the framerate on the FpsText and draws it to the Screen s.
func (f *FpsText) Draw(s *Screen) {
	f.time += s.TimeDelta()
	if f.time > f.update {
		fps := strconv.FormatFloat(1.0/s.TimeDelta(), 'f', 10, 64)
		f.SetText(fps)
		f.time -= f.update
	}
	f.Text.Draw(s)
}

func cubeIndex(x int, points [5]int) int {
	n := 0
	for _, p := range points {
		if x <= p {
			break
		} else {
			n++
		}
	}
	return n
}

// RgbTo256Color takes RGB values and returns the closest color for
// a 256-color terminal, as an Attr type.
func RgbTo256Color(r, g, b int) Attr {
	cubepoints := [5]int{47, 115, 155, 195, 235}
	r256 := cubeIndex(r, cubepoints)
	g256 := cubeIndex(g, cubepoints)
	b256 := cubeIndex(b, cubepoints)
	return Attr(r256*36 + g256*6 + b256 + 17)
}
