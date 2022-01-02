package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

const (
	MaxHealth   float64 = 100
	Consumption float64 = 0.02
)

const HealthType ecs.ComponentType = "Health"

type Health struct {
	Value float64
}

func (h *Health) Type() ecs.ComponentType {
	return HealthType
}

func (h *Health) CurrentHealth() int {

	// out of bounds -> dead
	if h.Value == 0 || h.Value == 100 {
		return 0
	}

	return int(h.Value)
}

func (h *Health) EatFood(food *Food) {
	h.Value = normalize(h.Value + food.Value)
}

func (h *Health) Consume() {
	h.Value = normalize(h.Value - Consumption)
}

func normalize(val float64) float64 {
	if val < 0 {
		return 0
	}
	if val > MaxHealth {
		return MaxHealth
	}
	return val
}
