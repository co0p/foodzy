package foodzy

import (
	"github.com/co0p/foodzy/screen"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

const (
	ScreenWidth  int    = 800
	ScreenHeight int    = 600
	GameName     string = "Foodzy"
)

type Game struct {
	screenManager *screen.Manager
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	startScreen := screen.NewStartScreen(ScreenWidth, ScreenHeight)
	gameScreen := screen.NewGameScreen(ScreenWidth, ScreenHeight)

	screenManager := screen.NewManager()
	screenManager.AddScreen(startScreen)
	screenManager.AddScreen(gameScreen)

	screenManager.ActiveScreen(startScreen.Name())

	return &Game{
		screenManager: screenManager,
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
