package entity

import (
	"github.com/co0p/foodzy/component"
)

func NewFood(nutrient *component.Nutrient, sprite *component.Sprite, velocity *component.Velocity, position *component.Position) *entity {

	entity := NewEntity("", true)
	entity.AddComponent(nutrient)
	entity.AddComponent(sprite)
	w, h := sprite.Image.Size()

	entity.AddComponent(position)
	entity.AddComponent(velocity)
	entity.AddComponent(&component.Collision{Width: float64(w), Height: float64(h)})

	return entity
}
