package main

import (
	"github.com/co0p/foodzy"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {

	game := foodzy.NewGame()

	ebiten.SetWindowSize(foodzy.ScreenWidth, foodzy.ScreenHeight)
	ebiten.SetWindowTitle(foodzy.GameName)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
