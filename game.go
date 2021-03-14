package foodzy

import (
	"github.com/co0p/foodzy/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"log"
)

const (
	ScreenWidth  int    = 800
	ScreenHeight int    = 600
	GameName     string = "Foodzy"
)

type Entity interface {
	Draw(*ebiten.Image)
	Update(*Game)
}

type Game struct {
	tick         float64
	Speed        float64
	Food         []Entity
	background   Entity
	player       Entity
	audioContext *audio.Context
}

func (g *Game) loadFood() {
	g.Food = append(g.Food, NewFood("beer", assets.Beer))
	g.Food = append(g.Food, NewFood("bread", assets.Bread))
	g.Food = append(g.Food, NewFood("carrot", assets.Carrot))
	g.Food = append(g.Food, NewFood("cheese", assets.Cheese))
	g.Food = append(g.Food, NewFood("fish", assets.Fish))
	g.Food = append(g.Food, NewFood("lemon", assets.Lemon))
	g.Food = append(g.Food, NewFood("meat", assets.Meat))
	g.Food = append(g.Food, NewFood("strawberry", assets.Strawberry))
	g.Food = append(g.Food, NewFood("tomato", assets.Tomato))
}

func (g *Game) Update() error {

	g.player.Update(g)

	for _, v := range g.Food {
		v.Update(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)
	g.player.Draw(screen)

	for _, v := range g.Food {
		v.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) initializeAudio() {

	audioContext := audio.NewContext(44100)
	if audioContext == nil {
		log.Fatal("could not initialize Audio")
	}

	g.audioContext = audioContext
}

func NewGame() *Game {

	game := &Game{
		Speed: 0.12,
		Food:  []Entity{},
	}

	game.player = NewPlayer()
	game.background = NewBackground()

	game.loadFood()
	return game
}
