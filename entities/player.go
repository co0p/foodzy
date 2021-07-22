package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
)

const playerSize = 50

func NewPlayer(ScreenWidth int, ScreenHeight int) *Entity {
	entity := NewEntity("player", true)

	entity.AddComponent(&components.Position{
		X: float64(ScreenWidth / 2.0),
		Y: float64(ScreenHeight - playerSize*2.0),
	})
	sprite := components.NewSprite("player", assets.Plate)
	width, height := sprite.Image.Size()
	entity.AddComponent(sprite)
	entity.AddComponent(&components.Dimension{Width: float64(width), Height: float64(height)})
	entity.AddComponent(&components.KeyboardMover{Speed: 5.0})
	entity.AddComponent(&components.Collision{})

	return entity
}
