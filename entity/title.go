package entity

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
)

func NewTitle(screenWidth int, screenHeight int) *Entity {
	sprite := component.NewSprite("title", assets.Title)

	entity := NewEntity("title", true)
	entity.AddComponent(&component.Velocity{X: 0, Y: 1})
	entity.AddComponent(sprite)

	w, _ := sprite.Image.Size()
	width := float64(screenWidth-w) / 2

	entity.AddComponent(&component.Transform{X: width, Y: -150.0, Z: 1, Scale: 1})
	return entity
}
