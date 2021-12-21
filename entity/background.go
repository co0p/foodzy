package entity

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
)

func NewBackground() *entity {
	bg := NewEntity("background", true)
	sprite := component.NewSprite("background", assets.Background)

	bg.AddComponent(&component.Position{X: 0, Y: 0})
	bg.AddComponent(sprite)
	return bg
}
