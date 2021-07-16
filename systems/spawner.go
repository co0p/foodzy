package systems

import (
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

// FoodManager is responsible for spawning random food based on the spawning frequency
type FoodManager struct {
	manager        *entities.Manager
	spawnFrequency int
	coolDown       int
	windowWidth    int
}

func NewFoodSystem(manager *entities.Manager, frequency int, windowWidth int) *FoodManager {

	return &FoodManager{
		manager:        manager,
		spawnFrequency: frequency,
		windowWidth:    windowWidth,
	}
}

func (s *FoodManager) Draw(image *ebiten.Image) {}

func (s *FoodManager) Update() error {
	if s.coolDown > 0 {
		s.coolDown--
		return nil
	}

	food := entities.NewRandomFood(s.windowWidth)
	s.manager.AddEntity(&food)

	s.coolDown = s.spawnFrequency

	return nil
}
