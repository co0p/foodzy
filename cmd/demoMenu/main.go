package main

import "github.com/hajimehoshi/ebiten/v2"

type MenuDemo struct {
}

func (m *MenuDemo) Update() error {
	return nil
}

func (m *MenuDemo) Draw(screen *ebiten.Image) {
}

func (m *MenuDemo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
}

func main() {

	game := MenuDemo{}

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}

}
