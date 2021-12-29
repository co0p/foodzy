package main

import (
	"github.com/co0p/foodzy"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	ebiten.SetWindowSize(foodzy.ScreenWidth, foodzy.ScreenHeight)
	ebiten.SetWindowTitle(foodzy.GameName)

	game := foodzy.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
