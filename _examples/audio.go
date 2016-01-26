package main

// A very simple music player. Takes an audio file as its argument.

import (
	tl "github.com/JoelOtter/termloop"
	tlx "github.com/JoelOtter/termloop/extra"
	"os"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

type SoundCtrl struct {
	audio *tlx.Audio
	text  []*tl.Text
	track *tlx.Track
}

func CreateSoundCtrl(filename string, loop bool) *SoundCtrl {
	a, err := tlx.InitAudio()
	chk(err)
	t1 := tl.NewText(1, 1, "Push the right arrow to play", tl.ColorWhite, 0)
	t2 := tl.NewText(1, 3, "Push the up arrow to pause", tl.ColorWhite, 0)
	t3 := tl.NewText(1, 5, "Push the left arrow to restart", tl.ColorWhite, 0)
	t4 := tl.NewText(1, 7, "Push the down arrow to stop", tl.ColorWhite, 0)
	text := []*tl.Text{t1, t2, t3, t4}
	track, err := a.LoadTrack(filename, loop)
	chk(err)
	return &SoundCtrl{
		audio: a,
		track: track,
		text:  text,
	}
}

func (sc *SoundCtrl) Draw(s *tl.Screen) {
	for _, t := range sc.text {
		t.Draw(s)
	}
}

func (sc *SoundCtrl) Tick(ev tl.Event) {
	if ev.Type == tl.EventKey {
		switch ev.Key {
		case tl.KeyArrowRight:
			sc.track.Play()
		case tl.KeyArrowLeft:
			sc.track.Restart()
		case tl.KeyArrowUp:
			sc.track.Pause()
		case tl.KeyArrowDown:
			sc.track.Stop()
		}
	}
}

func main() {
	g := tl.NewGame()
	sound := CreateSoundCtrl(os.Args[1], true)
	defer sound.audio.Stop()
	g.Screen().AddEntity(sound)
	g.Start()
}
