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

func (e *CleanupSystem) Draw(image *ebiten.Image) {}

func (e *CleanupSystem) Update() error {

	candidates := e.manager.QueryByComponents(component.PositionType)
	for _, entity := range candidates {
		position := entity.GetComponent(component.PositionType).(*component.Position)

		if int(position.Y) > e.screenHeight {
			entity.Active = false
		}
	}

	if e.coolDown > 0 {
		e.coolDown--
		return nil
	}

	e.manager.RemoveInactive()
	e.coolDown = e.runFrequency
	return nil
}
