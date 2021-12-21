package component

import (
	"github.com/co0p/foodzy/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType ComponentType = "Sprite"

type Sprite struct {
	Image *ebiten.Image
	name  string
}

func (s *Sprite) Type() ComponentType {
	return SpriteType
}

func NewSprite(tag string, img []byte) *Sprite {
	sprite, _ := utils.LoadImage(img)
	return &Sprite{name: tag, Image: sprite}
}
