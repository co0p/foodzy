package entity

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
)

func NewBackground() *entity {
	bg := NewEntity("background", true)
	sprite := component.NewSprite("background", assets.Background)

	bg.AddComponent(&component.Transform{X: 0, Y: 0, Scale: 1})
	bg.AddComponent(sprite)
	return bg
}
