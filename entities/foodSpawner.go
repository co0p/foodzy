package entities

import (
	"github.com/co0p/foodzy/components"
)

func NewFoodSpawner(spawnRate int) *Entity {

	entity := Entity{Active: true}
	velocity := struct {
		X float64
		Y float64
	}{X: 0, Y: 2.5}

	entity.AddComponent(&components.FoodSpawner{
		Rate:     spawnRate,
		CoolDown: spawnRate,
		Velocity: velocity,
	})
	return &entity
}
