package components

import (
	"github.com/co0p/foodzy/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image *ebiten.Image
	name  string
}

func (s Sprite) ID() string {
	return s.name
}

func NewSprite(name string, img []byte) *Sprite {
	sprite, _ := utils.LoadImage(img)
	return &Sprite{name: name, Image: sprite}
}
