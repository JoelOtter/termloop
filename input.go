package termloop

import "github.com/nsf/termbox-go"

type Input struct {
	endKey termbox.Key
	eventQ chan termbox.Event
	ctrl   chan bool
}

func NewInput() *Input {
	i := Input{eventQ: make(chan termbox.Event),
		ctrl:   make(chan bool, 2),
		endKey: termbox.KeyEsc}
	return &i
}

func (i *Input) Start() {
	go Poll(i)
}

func (i *Input) Stop() {
	i.ctrl <- true
}

func Poll(i *Input) {
loop:
	for {
		select {
		case <-i.ctrl:
			break loop
		default:
			i.eventQ <- termbox.PollEvent()
		}
	}
}
