package entity

import (
	"github.com/co0p/foodzy/component"
	"math/rand"
)

func NewFood(nutrient *component.Nutrient, sprite *component.Sprite, velocity *component.Velocity, transform *component.Transform) *Entity {

	entity := NewEntity("", true)
	entity.AddComponent(nutrient)
	entity.AddComponent(sprite)
	w, h := sprite.Image.Size()

	transform.Scale = randomScale()
	velocity.Y += randomSpeedBoost()

	entity.AddComponent(transform)
	entity.AddComponent(velocity)
	entity.AddComponent(&component.Collision{Width: float64(w), Height: float64(h)})

	return entity
}

func randomScale() float64 {
	scale := 0.9
	chance := rand.Float64()
	if chance < 0.8 {
		scale = 0.8
	}
	if chance < 0.4 {
		scale = 0.7
	}
	return scale
}

func randomSpeedBoost() float64 {
	speed := 0.0
	chance := rand.Float64()

	if chance > 0.6 {
		speed = 0.7
	}
	if chance > 0.7 {
		speed = 0.8
	}
	if chance > 0.8 {
		speed = 1.0
	}
	if chance > 0.9 {
		speed = 2.0
	}

	return speed

}
