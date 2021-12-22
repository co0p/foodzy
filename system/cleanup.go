package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type CleanupSystem struct {
	manager      *entity.Manager
	runFrequency int
	coolDown     int
	screenHeight int
}

func NewCleanupSystem(manager *entity.Manager, frequency int, screenHeight int) *CleanupSystem {

	return &CleanupSystem{
		manager:      manager,
		runFrequency: frequency,
		coolDown:     frequency,
		screenHeight: screenHeight,
	}
}

func (s *CleanupSystem) Draw(image *ebiten.Image) {}

func (s *CleanupSystem) Update() error {

	candidates := s.manager.QueryByComponents(component.TransformType)
	for _, e := range candidates {
		position := e.GetComponent(component.TransformType).(*component.Transform)

		if int(position.Y) > s.screenHeight {
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
