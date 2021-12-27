package foodzy

import (
	"bytes"
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/screen"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	ScreenWidth     int    = 800
	ScreenHeight    int    = 600
	GameName        string = "Foodzy"
	AudioSampleRate int    = 44100
)

type Game struct {
	screenManager *screen.Manager
}

func NewGame() *Game {

	rand.Seed(time.Now().UnixNano())

	screenManager := screen.NewManager()

	// actions
	closeGame := func() { os.Exit(0) }
	startGame := func() { screenManager.ActivateScreen(screen.GameScreenName) }
	showStartScreen := func() { screenManager.ActivateScreen(screen.StartScreenName) }

	// screens
	startScreen := screen.NewStartScreen(ScreenWidth, ScreenHeight)
	startScreen.AddAction(screen.ActionActivateGameScreen, startGame)
	startScreen.AddAction(screen.ActionQuit, closeGame)

	gameScreen := screen.NewGameScreen(ScreenWidth, ScreenHeight)
	gameScreen.AddAction(screen.ActionQuit, closeGame)
	gameScreen.AddAction(screen.ActionActivateStartScreen, showStartScreen)

	screenManager.AddScreen(startScreen)
	screenManager.AddScreen(gameScreen)

	screenManager.ActivateScreen(screen.StartScreenName)

	// audio
	startAudio()

	return &Game{
		screenManager: screenManager,
	}
}

func (g *Game) Update() error {

	if ebiten.IsWindowBeingClosed() {
		// TODO: raise WINDOW_CLOSE EVENT
		log.Printf("[game] window close received\n")
	}

	return g.screenManager.Current.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screenManager.Current.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func startAudio() {
	log.Print("[sound] start initializing")

	audioContext := audio.NewContext(AudioSampleRate)
	src, err := mp3.DecodeWithSampleRate(AudioSampleRate, bytes.NewReader(assets.Soundtrack))

	if err != nil {
		log.Fatal("failed loading soundtrack")
	}
	s := audio.NewInfiniteLoop(src, src.Length())
	audioPlayer, err := audio.NewPlayer(audioContext, s)

	if err != nil {
		log.Fatal("failed initializing audio player")
	}

	audioPlayer.Play()
	log.Print("[sound] initializing done")
}
