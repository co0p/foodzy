package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type CleanupSystem struct {
	manager      *ecs.Manager
	runFrequency int
	coolDown     int
}

func NewCleanupSystem(manager *ecs.Manager, frequency int) *CleanupSystem {

	return &CleanupSystem{
		manager:      manager,
		runFrequency: frequency,
		coolDown:     frequency,
	}
}

func (s *CleanupSystem) Draw(image *ebiten.Image) {}

func (s *CleanupSystem) Update() error {

	candidates := s.manager.QueryByComponents(component.TransformType)
	for _, e := range candidates {
		position := e.GetComponent(component.TransformType).(*component.Transform)

		if int(position.Y) > ScreenHeight {
			e.Active = false
		}
	}

	if s.coolDown > 0 {
		s.coolDown--
		return nil
	}

	s.manager.RemoveInactive()
	s.coolDown = s.runFrequency
	return nil
}
