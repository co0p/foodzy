package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	manager *entities.Manager
}

func NewMovementSystem(manager *entities.Manager) *MovementSystem {
	return &MovementSystem{
		manager: manager,
	}
}

func (s *MovementSystem) Update() error {

	entities := s.manager.QueryByComponents(&components.Velocity{}, &components.Position{})

	for _, entity := range entities {

		velocity := entity.GetComponent(&components.Velocity{}).(*components.Velocity)
		position := entity.GetComponent(&components.Position{}).(*components.Position)

		position.X = position.X + velocity.X
		position.Y = position.Y + velocity.Y
	}
	return nil
}

func (s *MovementSystem) Draw(image *ebiten.Image) { /* nothing to do */ }
