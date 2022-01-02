package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

// Food represents the kcal and nutrient value relative to 100%
type Food struct {
	Value float64
}

const FoodType ecs.ComponentType = "Food"

func (n *Food) Type() ecs.ComponentType {
	return FoodType
}
