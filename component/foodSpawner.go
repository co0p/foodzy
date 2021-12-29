package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type FoodSpawner struct {
	CoolDown int
	Rate     int
	Variance float64
	Velocity Velocity
}

const FoodSpawnerType ecs.ComponentType = "FoodSpawner"

func (f *FoodSpawner) Type() ecs.ComponentType {
	return FoodSpawnerType
}
