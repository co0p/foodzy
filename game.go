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

const (
	ActionQuit                   scene.Action = "quit"
	ActionActivateStartScreen    scene.Action = "startscene"
	ActionActivateGameScreen     scene.Action = "gamescene"
	ActionActivateGameOverScreen scene.Action = "gameover"
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
	startGame := func() { sceneManager.ActivateScreen(GameScreenName, false) }
	showStartScreen := func() { sceneManager.ActivateScreen(StartScreenName, false) }
	showGameOverScene := func() { sceneManager.ActivateScreen(GameOverSceneName, false) }

	// screens
	startScreen := NewStartScreen()
	startScreen.AddAction(ActionActivateGameScreen, startGame)
	startScreen.AddAction(ActionQuit, closeGame)

	gameOverScene := NewGameOverScene(soundManager)
	gameOverScene.AddAction(ActionActivateGameScreen, startGame)
	gameOverScene.AddAction(ActionQuit, closeGame)

	gameScreen := NewGameScreen(soundManager)
	gameScreen.AddAction(ActionQuit, closeGame)
	gameScreen.AddAction(ActionActivateStartScreen, showStartScreen)
	gameScreen.AddAction(ActionActivateGameOverScreen, showGameOverScene)

	sceneManager.AddScreen(startScreen)
	sceneManager.AddScreen(gameScreen)
	sceneManager.AddScreen(gameOverScene)
	sceneManager.ActivateScreen(StartScreenName, false)

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
