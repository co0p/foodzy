package sound

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
)

const AudioSampleRate int = 44100

type SoundManager struct {
	context *audio.Context
	players map[string]*audio.Player
}

func NewSoundManager() *SoundManager {
	audioContext := audio.NewContext(AudioSampleRate)

	return &SoundManager{
		context: audioContext,
		players: make(map[string]*audio.Player),
	}
}

func (a *SoundManager) Add(key string, data []byte) {

	src, err := mp3.DecodeWithSampleRate(AudioSampleRate, bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed loading clip:" + err.Error())
	}

	player, err := a.context.NewPlayer(src)
	if err != nil {
		log.Fatal("failed creating player:" + err.Error())
	}
	a.players[key] = player
}

func (a *SoundManager) AddLoop(key string, data []byte) {

	src, err := mp3.DecodeWithSampleRate(AudioSampleRate, bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed loading clip:" + err.Error())
	}

	stream := audio.NewInfiniteLoop(src, src.Length())
	player, err := a.context.NewPlayer(stream)
	if err != nil {
		log.Fatal("failed creating player:" + err.Error())
	}
	a.players[key] = player
}

func (a *SoundManager) Play(key string) {
	player, ok := a.players[key]
	if !ok {
		log.Printf("could not find clip:%s\n", key)
		return
	}

	if !player.IsPlaying() {
		player.Rewind()
		player.Play()
	}
}

func (a *SoundManager) Stop(key string) {
	player, ok := a.players[key]
	if !ok {
		log.Printf("could not find clip:%s\n", key)
		return
	}
	player.Pause()
	player.Rewind()
}
