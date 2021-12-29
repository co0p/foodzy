package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewFoodSpawner(spawnRate int) *ecs.Entity {

	entity := ecs.NewEntity("spawner", true)
	velocity := component.Velocity{X: 0, Y: 2.5, Z: 0}

	entity.AddComponent(&component.FoodSpawner{
		Rate:     spawnRate,
		CoolDown: spawnRate,
		Velocity: velocity,
	})
	return entity
}
