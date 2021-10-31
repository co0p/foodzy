package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
)

func NewBackground() *Entity {
	bg := NewEntity("background", true)
	sprite := components.NewSprite("background", assets.Background)

	bg.AddComponent(&components.Position{X: 0, Y: 0})
	bg.AddComponent(sprite)
	return bg
}
