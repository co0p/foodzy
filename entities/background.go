package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
)

func NewBackground() *Entity {
	bg := NewEntity("background", true)
	bg.AddComponent(&components.Position{X: 0, Y: 0})
	bg.AddComponent(components.NewSprite("background", assets.Background))
	return bg
}
