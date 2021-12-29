package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type Collision struct {
	Width  float64
	Height float64
}

const CollisionType ecs.ComponentType = "Collision"

func (c *Collision) Type() ecs.ComponentType {
	return CollisionType
}
