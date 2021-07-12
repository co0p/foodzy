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
	entity.AddComponent(components.NewSprite("player", assets.Plate))
	entity.AddComponent(&components.KeyboardMover{Speed: 5.0})

	return entity
}
