package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type MouseCollider struct {
	Width  float64
	Height float64
}

const MouseColliderType ecs.ComponentType = "MouseCollider"

func (c MouseCollider) Type() ecs.ComponentType {
	return MouseColliderType
}
