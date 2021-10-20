package entities

import (
	"github.com/co0p/foodzy/components"
)

func NewFoodSpawner(spawnRate int) *Entity {

	entity := Entity{Active: true}
	entity.AddComponent(&components.FoodSpawner{
		Rate:     spawnRate,
		CoolDown: spawnRate,
		Velocity: struct {
			X float64
			Y float64
		}{X: 0, Y: 2.5},
		Types: []components.FoodType{components.FoodTreat, components.FoodFish, components.FoodFruit, components.FoodDrink},
	})
	return &entity
}
