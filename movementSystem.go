package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	manager *ecs.EntityManager
}

func NewMovementSystem(manager *ecs.EntityManager) *MovementSystem {
	return &MovementSystem{
		manager: manager,
	}
}

func (s *MovementSystem) Update() error {

	entities := s.manager.QueryByComponents(component.VelocityType, component.TransformType)

	for _, e := range entities {

		velocity := e.GetComponent(component.VelocityType).(*component.Velocity)
		position := e.GetComponent(component.TransformType).(*component.Transform)

		position.X = position.X + velocity.X
		position.Y = position.Y + velocity.Y
		position.Z = position.Z + velocity.Z
	}
	return nil
}

func (s *MovementSystem) Draw(image *ebiten.Image) { /* nothing to do */ }
