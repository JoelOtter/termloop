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
	LineEmpty LineType = 0
	LineSingleBorder = 1 << iota
	LineDoubleBorder
)

var BorderTheme = map[uint]*BorderDefinitions{
  uint(LineEmpty): &BorderDefinitions{hc: ' ', vc: ' ', luc: ' ', ruc: ' ', lbc: ' ', rbc: ' '},
  uint(LineSingleBorder): &BorderDefinitions{hc: '─', vc: '│', luc: '┌', ruc: '┐', lbc: '└', rbc: '┘'},
  uint(LineDoubleBorder): &BorderDefinitions{hc: '═', vc: '║', luc: '╔', ruc: '╗', lbc: '╚', rbc: '╝'},
}
