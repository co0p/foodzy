package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type Interaction struct {
	Action func(manager *ecs.EntityManager)
}

const InteractionType ecs.ComponentType = "Interaction"

func (c *Interaction) Type() ecs.ComponentType {
	return InteractionType
}
