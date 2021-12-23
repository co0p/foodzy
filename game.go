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
	layerManager := screen.NewManager()

	gameScreen := screen.NewGameScreen(ScreenWidth, ScreenHeight)

	layerManager.AddScreen(gameScreen)
	if err := layerManager.ActiveScreen(gameScreen.Name()); err != nil {
		panic("failed to activate screen: " + err.Error())
	}

	return &Game{
		screenManager: layerManager,
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
