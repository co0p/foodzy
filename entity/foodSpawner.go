package entity

import "github.com/co0p/foodzy/component"

func NewFoodSpawner(spawnRate int) *entity {

	entity := NewEntity("spawner", true)
	velocity := component.Velocity{X: 0, Y: 2.5}

	entity.AddComponent(&component.FoodSpawner{
		Rate:     spawnRate,
		CoolDown: spawnRate,
		Velocity: velocity,
	})
	return entity
}
