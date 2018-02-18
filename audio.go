package numino

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var (
	isSpeakerInitialized = false
	sounds               = make(map[Sound]beep.StreamSeeker)
	loopingSoundCtrl     = make(map[int]*beep.Ctrl)
	loopingSoundCount    = 0
	loaded               = false
)

type Sound int

const (
	ShiftSound Sound = iota
	SlamSound
	MergeSound
	DieSound
	BackgroundMusic
)

// LoadSounds loads all sounds.
func LoadSounds() {
	if loaded {
		return
	}
	sounds[DieSound] = load("die")
	sounds[SlamSound] = load("slam")
	sounds[MergeSound] = load("merge")
	sounds[ShiftSound] = load("shift")
	sounds[BackgroundMusic] = load("background-slow")
	loaded = true
}

// PlaySound plays the specified sound.
func PlaySound(sound Sound) {
	speaker.Play(sounds[sound])
	// move seeker back to start of recording.
	sounds[sound].Seek(0)
}

// LoopSound loops a sound.
//
// A number that can be used to stop the sound via StopSound() is returned.
func LoopSound(sound Sound) int {
	speaker.Lock()
	loopingSoundCount++
	loopingSoundCtrl[loopingSoundCount] = &beep.Ctrl{
		Streamer: beep.Loop(-1, sounds[sound]),
	}

	speaker.Unlock()
	speaker.Play(loopingSoundCtrl[loopingSoundCount])
	return loopingSoundCount
}

// StopSound stops a looping sound.
//
// If the given ref does not identify a looping sound, an error is logged.
func StopSound(ref int) {
	if _, ok := loopingSoundCtrl[ref]; !ok {
		log.Println("invalid sound ref:", ref)
	}
	speaker.Lock()
	loopingSoundCtrl[ref].Paused = true
	loopingSoundCtrl[ref].Streamer = nil
	delete(loopingSoundCtrl, ref)
	speaker.Unlock()
}

func load(name string) beep.StreamSeeker {
	f, err := os.Open("./assets/audio/" + name + ".wav")
	if err != nil {
		log.Println(err)
	}
	streamSeeker, format, err := wav.Decode(f)
	if err != nil {
		log.Println(err)
	}

	if !isSpeakerInitialized {
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/35))
		isSpeakerInitialized = true
	}

	return streamSeeker
}
