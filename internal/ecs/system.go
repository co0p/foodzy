package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Draw(image *ebiten.Image)
	Update() error
}
