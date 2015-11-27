package termloop

import "github.com/nsf/termbox-go"

type input struct {
	endKey termbox.Key
	eventQ chan termbox.Event
	ctrl   chan bool
}

func newInput() *input {
	i := input{eventQ: make(chan termbox.Event),
		ctrl:   make(chan bool, 2),
		endKey: termbox.KeyCtrlC}
	return &i
}

func (i *input) start() {
	go poll(i)
}

func (i *input) stop() {
	i.ctrl <- true
}

func poll(i *input) {
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
