package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type CleanupSystem struct {
	manager      *entities.Manager
	runFrequency int
	coolDown     int
	screenHeight int
}

func NewCleanupSystem(manager *entities.Manager, frequency int, screenHeight int) *CleanupSystem {

	return &CleanupSystem{
		manager:      manager,
		runFrequency: frequency,
		coolDown:     frequency,
		screenHeight: screenHeight,
	}
}

func (e *CleanupSystem) Draw(image *ebiten.Image) {}

func (e *CleanupSystem) Update() error {

	candidates := e.manager.QueryByComponents(&components.Position{})
	for _, entity := range candidates {
		position := entity.GetComponent(&components.Position{}).(*components.Position)

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
