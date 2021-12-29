package component

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType ecs.ComponentType = "Sprite"

type Sprite struct {
	Image *ebiten.Image
	name  string
}

func (s *Sprite) Type() ecs.ComponentType {
	return SpriteType
}

func NewSprite(tag string, img []byte) *Sprite {
	sprite, _ := asset.LoadImage(img)
	return &Sprite{name: tag, Image: sprite}
}
