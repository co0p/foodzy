package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	manager *entity.Manager
}

func NewMovementSystem(manager *entity.Manager) *MovementSystem {
	return &MovementSystem{
		manager: manager,
	}
}

func (s *MovementSystem) Update() error {

	entities := s.manager.QueryByComponents(component.VelocityType, component.PositionType)

	for _, e := range entities {

		velocity := e.GetComponent(component.VelocityType).(*component.Velocity)
		position := e.GetComponent(component.PositionType).(*component.Position)

		position.X = position.X + velocity.X
		position.Y = position.Y + velocity.Y
	}
	return nil
}

func (s *MovementSystem) Draw(image *ebiten.Image) { /* nothing to do */ }
