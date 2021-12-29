package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type Velocity struct {
	X float64
	Y float64
	Z float64
}

const VelocityType ecs.ComponentType = "Velocity"

func (v Velocity) Type() ecs.ComponentType {
	return VelocityType
}
