package scene

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Init()
	Exit()
	Update() error
	Draw(screen *ebiten.Image)
	Name() string
}
