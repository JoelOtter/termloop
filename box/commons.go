package box

type Align uint
const (
	AlignNone Align = 0
	AlignLeft Align = 1 << iota
	AlignRight
	AlignHCenter
	AlignTop
	AlignBottom
	AlignVCenter
	AlignCenter = AlignHCenter | AlignVCenter
)

type LineType uint
const (
	FrameEmpty LineType = 0
	FrameSingleBorder = 1 << iota
	FrameDoubleBorder
)

var BorderTheme = map[uint]*BorderDefinitions{
  uint(FrameEmpty): &BorderDefinitions{hc: ' ', vc: ' ', luc: ' ', ruc: ' ', lbc: ' ', rbc: ' '},
  uint(FrameSingleBorder): &BorderDefinitions{hc: '─', vc: '│', luc: '┌', ruc: '┐', lbc: '└', rbc: '┘'},
  uint(FrameDoubleBorder): &BorderDefinitions{hc: '═', vc: '║', luc: '╔', ruc: '╗', lbc: '╚', rbc: '╝'},
}
