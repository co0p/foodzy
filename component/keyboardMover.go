package component

import (
	"github.com/co0p/foodzy/internal/ecs"
)

type KeyboardMover struct {
	Speed float64
}

const KeyboardMoverType ecs.ComponentType = "KeyboardMover"

func (k *KeyboardMover) Type() ecs.ComponentType {
	return KeyboardMoverType
}
