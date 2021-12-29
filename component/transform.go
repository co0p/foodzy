package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type Transform struct {
	X        float64
	Y        float64
	Z        float64
	Scale    float64
	Rotation float64
}

const TransformType ecs.ComponentType = "Transform"

func (p Transform) Type() ecs.ComponentType {
	return TransformType
}
