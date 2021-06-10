package foodzy

import (
	"bytes"
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
)

const (
	ScreenWidth     int    = 800
	ScreenHeight    int    = 600
	GameName        string = "Foodzy"
	AudioSampleRate int    = 44100
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
	player       *components.Element
	audioContext *audio.Context
	audioPlayer  *audio.Player
}

// NewGame returns a new and initialized game ready to play
func NewGame() *Game {
	game := &Game{
		Speed:      0.12,
		Food:       []Entity{},
		player:     NewPlayerElement(),
		background: NewBackground(),
	}

	game.loadFood()
	game.initializeAudio()

	return game
}

// Update gets called by ebiten and is meant to update any game logic
func (g *Game) Update() error {

	g.player.OnUpdate()

	for _, v := range g.Food {
		v.Update(g)
	}
	return nil
}

// Draw is called by ebiten and is meant to be used to draw stuff on the screen
func (g *Game) Draw(screen *ebiten.Image) {

	g.background.Draw(screen)
	g.player.OnDraw(screen)

	for _, entity := range g.Food {
		entity.Draw(screen)
	}
}

// Layout returns the same screen width and height; could be used to react on window change
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) initializeAudio() {
	log.Print("initializing audio ...")
	g.audioContext = audio.NewContext(AudioSampleRate)
	src, err := mp3.Decode(g.audioContext, bytes.NewReader(assets.Soundtrack))

	if err != nil {
		log.Fatal("failed loading soundtrack")
	}
	s := audio.NewInfiniteLoop(src, src.Length())
	g.audioPlayer, err = audio.NewPlayer(g.audioContext, s)

	if err != nil {
		log.Fatal("failed initializing audio player")
	}

	g.audioPlayer.Play()
	log.Print("initializing audio ... done")
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
