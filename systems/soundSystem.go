package systems

import (
	"bytes"
	"github.com/co0p/foodzy/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
)

const AudioSampleRate int = 44100

type SoundSystem struct {
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

func (s SoundSystem) Draw(image *ebiten.Image) {}

func (s SoundSystem) Update() {}

func NewSoundSystem() *SoundSystem {

	system := SoundSystem{}

	log.Print("initializing audio ...")
	system.audioContext = audio.NewContext(AudioSampleRate)
	src, err := mp3.Decode(system.audioContext, bytes.NewReader(assets.Soundtrack))

	if err != nil {
		log.Fatal("failed loading soundtrack")
	}
	s := audio.NewInfiniteLoop(src, src.Length())
	system.audioPlayer, err = audio.NewPlayer(system.audioContext, s)

	if err != nil {
		log.Fatal("failed initializing audio player")
	}

	system.audioPlayer.Play()
	log.Print("initializing audio ... done")

	return &system
}
