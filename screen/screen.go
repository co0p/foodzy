package screen

import "github.com/hajimehoshi/ebiten/v2"

type Name string

type Screen interface {
	Init()
	Exit()
	Update() error
	Draw(screen *ebiten.Image)
	Name() Name
}
