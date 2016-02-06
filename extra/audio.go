package extra

// extras/audio.go

// This file provides types and functions for loading and playing audio.
// Please note that this is in termloop/extras because it has external
// dependencies - it won't work without PortAudio and libsndfile being
// installed on your system.

import (
	"github.com/gordonklaus/portaudio"
	"github.com/mkb218/gosndfile/sndfile"
)

// The Audio type represents the audio controller.
// It should be used as a singleton, as multiple streams in PortAudio
// may not be supported on some hardware.
type Audio struct {
	tracks []*Track
	stream *portaudio.Stream
}

// The Track type represents an audio track which can be played.
type Track struct {
	loop     bool
	playing  bool
	playhead int
	buffer   []float32
	volume   float32
}

// InitAudio starts up PortAudio, creates a stream and
// returns a pointer to an Audio struct, or an error.
func InitAudio() (*Audio, error) {
	a := Audio{
		tracks: make([]*Track, 0),
	}
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	stream, err := portaudio.OpenDefaultStream(
		0, 2, float64(44100), 0, a.playCallback,
	)
	if err != nil {
		return nil, err
	}
	a.stream = stream
	a.stream.Start()
	return &a, nil
}

// Stop shuts down PortAudio
func (a *Audio) Stop() {
	portaudio.Terminate()
}

func (a *Audio) playCallback(out []float32) {
	for i := range out {
		var data float32

		for _, t := range a.tracks {
			if !t.playing {
				continue
			}
			if t.loop {
				data += (t.buffer[t.playhead%len(t.buffer)]) * t.volume
				t.playhead = (t.playhead + 1) % len(t.buffer)
			} else if t.playhead < len(t.buffer) {
				data += t.buffer[t.playhead] * t.volume
				t.playhead++
			} else {
				t.playing = false
				t.playhead = 0
			}
		}
		out[i] = data
	}
}

// LoadTrack reads an audio track from a file, and returns a pointer to
// a Track struct, or an error. The boolean parameter 'loop' determines
// whether or not a Track should loop when it is finished playing.
//
// Supported filetypes are whatever libsndfile supports, e.g. WAV or OGG.
func (a *Audio) LoadTrack(filename string, loop bool) (*Track, error) {
	// Load file
	var info sndfile.Info
	soundFile, err := sndfile.Open(filename, sndfile.Read, &info)
	if err != nil {
		return nil, err
	}
	buffer := make([]float32, info.Frames*int64(info.Channels))
	numRead, err := soundFile.ReadItems(buffer)
	if err != nil {
		return nil, err
	}
	defer soundFile.Close()

	// Create track
	track := Track{
		loop:   loop,
		buffer: buffer[:numRead],
		volume: 1,
	}

	a.tracks = append(a.tracks, &track)

	return &track, nil
}

// Play triggers a Track to start playing.
func (t *Track) Play() {
	t.playing = true
}

// Stop stops a Track playing, and resets it to the beginning.
func (t *Track) Stop() {
	t.playing = false
	t.playhead = 0
}

// Pause stops a Track playing, but does not reset its position.
// The track can be resumed by calling Play().
func (t *Track) Pause() {
	t.playing = false
}

// Restart resets the Track to the beginning but does not stop playback.
func (t *Track) Restart() {
	t.playhead = 0
}

// Volume returns the current volume of the Track. Default is 1.0.
func (t *Track) Volume() float32 {
	return t.volume
}

// SetVolume sets the Track's volume to v.
func (t *Track) SetVolume(v float32) {
	t.volume = v
}

// Looping returns whether or not the Track is set to loop.
func (t *Track) Looping() bool {
	return t.loop
}

// SetLooping sets whether or not a Track should loop.
func (t *Track) SetLooping(looping bool) {
	t.loop = looping
}
