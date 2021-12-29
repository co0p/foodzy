package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	ScreenWidth  int    = 800
	ScreenHeight int    = 600
	GameName     string = "Foodzy"
)

const (
	SoundEat        string = "eating"
	SoundBackground        = "soundtrack"
)

type Game struct {
	screenManager *scene.SceneManager
	soundManager  *sound.SoundManager
}

func NewGame() *Game {

	rand.Seed(time.Now().UnixNano())

	sceneManager := scene.NewSceneManager()
	soundManager := sound.NewSoundManager()

	// actions
	closeGame := func() { os.Exit(0) }
	startGame := func() { sceneManager.ActivateScreen(GameScreenName) }
	showStartScreen := func() { sceneManager.ActivateScreen(StartScreenName) }

	// screens
	startScreen := NewStartScreen()
	startScreen.AddAction(scene.ActionActivateGameScreen, startGame)
	startScreen.AddAction(scene.ActionQuit, closeGame)

	gameScreen := NewGameScreen(soundManager)
	gameScreen.AddAction(scene.ActionQuit, closeGame)
	gameScreen.AddAction(scene.ActionActivateStartScreen, showStartScreen)

	sceneManager.AddScreen(startScreen)
	sceneManager.AddScreen(gameScreen)
	sceneManager.ActivateScreen(StartScreenName)

	// sound
	soundManager.Add(SoundEat, asset.Eating)
	soundManager.AddLoop(SoundBackground, asset.Soundtrack)
	soundManager.Play(SoundBackground)

	return &Game{
		soundManager:  soundManager,
		screenManager: sceneManager,
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
