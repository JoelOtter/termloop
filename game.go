package termloop

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

// Represents a top-level Termloop application.
type Game struct {
	screen *Screen
	input  *input
	debug  bool
	logs   []string
}

// NewGame creates a new Game, along with a Screen and input handler.
// Returns a pointer to the new Game.
func NewGame() *Game {
	g := Game{
		screen: NewScreen(),
		input:  newInput(),
		logs:   make([]string, 0),
	}
	return &g
}

// SetScreen sets the current Screen of a Game.
func (g *Game) SetScreen(s *Screen) {
	g.screen = s
}

// CreateLevel creates a new BaseLevel with background color bg,
// and sets it as the Game's level.
func (g *Game) CreateLevel(bg Attr) {
	g.screen.level = NewBaseLevel(Cell{Bg: bg})
}

// SetLevel sets the game's current level to be l.
func (g *Game) SetLevel(l Level) {
	g.screen.level = l
}

// Level returns the game's current level.
func (g *Game) Level() Level {
	return g.screen.level
}

// AddEntity adds a Drawable to the Game's current Screen, to be rendered.
func (g *Game) AddEntity(d Drawable) {
	g.screen.entities = append(g.screen.entities, d)
}

// DebugOn returns a bool showing whether or not debug mode is on.
func (g *Game) DebugOn() bool {
	return g.debug
}

// SetDebugOn sets debug mode's on status to be debugOn.
func (g *Game) SetDebugOn(debugOn bool) {
	g.debug = debugOn
}

// Log takes a log string and additional parameters, which can be substituted
// into the string using standard fmt.Printf rules.
// The formatted string is added to Game g's logs. If debug mode is on, the log will
// be printed to the terminal when Termloop exits.
func (g *Game) Log(log string, items ...interface{}) {
	toLog := "[" + time.Now().Format(time.StampMilli) + "] " +
		fmt.Sprintf(log, items...)
	g.logs = append(g.logs, toLog)
}

func (g *Game) dumpLogs() {
	if g.debug {
		fmt.Println("=== Logs: ===")
		for _, l := range g.logs {
			fmt.Println(l)
		}
		fmt.Println("=============")
	}
}

// Start starts a Game running. This should be the last thing called in your
// main function. By default, the escape key exits.
func (g *Game) Start() {
	// Init Termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer g.dumpLogs()
	defer termbox.Close()
	g.screen.resize(termbox.Size())

	// Init input
	g.input.start()
	defer g.input.stop()
	clock := time.Now()

mainloop:
	for {
		select {
		case ev := <-g.input.eventQ:
			if ev.Key == g.input.endKey {
				break mainloop
			} else if EventType(ev.Type) == EventResize {
				g.screen.resize(ev.Width, ev.Height)
			} else if EventType(ev.Type) == EventError {
				g.Log(ev.Err.Error())
			}
			g.screen.Tick(convertEvent(ev))
		default:
			g.screen.Tick(Event{Type: EventNone})
		}
		update := time.Now()
		g.screen.delta = update.Sub(clock).Seconds()
		clock = update
		g.screen.Draw()
	}
}
