package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
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

type ActionType func(*ecs.EntityManager)

type Game struct {
	screenManager *scene.SceneManager
	soundManager  *sound.SoundManager
}

func NewGame() *Game {

	rand.Seed(time.Now().UnixNano())

	sceneManager := scene.NewSceneManager()
	soundManager := sound.NewSoundManager()

	// actions
	exit := func(m *ecs.EntityManager) { os.Exit(0) }
	startGame := func(m *ecs.EntityManager) { sceneManager.Activate(GameSceneName) }
	showPauseScene := func(m *ecs.EntityManager) { sceneManager.Activate(PauseSceneName) }
	showGameOverScene := func(*ecs.EntityManager) { sceneManager.Activate(GameOverSceneName) }

	// scenes
	sceneManager.AddScene(NewStartScene(startGame, exit))
	sceneManager.AddScene(NewPauseScene(startGame))
	sceneManager.AddScene(NewGameScene(soundManager, showGameOverScene, showPauseScene))
	sceneManager.AddScene(NewGameOverScene(soundManager, exit))

	sceneManager.Activate(StartSceneName)

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
	return g.screenManager.Current.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.screenManager.Current.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
