package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type ControllerSystem struct {
	manager *entities.Manager
	height  int
	width   int
}

func NewControllerSystem(manager *entities.Manager, width int, height int) *ControllerSystem {
	return &ControllerSystem{manager: manager, width: width, height: height}
}

func (s *ControllerSystem) Draw(image *ebiten.Image) { /* nothing to do */ }

func (s *ControllerSystem) Update() error {

	entities := s.manager.Query(&components.KeyboardMover{}, &components.Sprite{}, &components.Position{})

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

		for _, entity := range entities {

			position := entity.GetComponent(&components.Position{}).(*components.Position)
			move := entity.GetComponent(&components.KeyboardMover{}).(*components.KeyboardMover)

			newLeft := position.X - move.Speed
			if newLeft > 0 {
				position.X = newLeft
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		for _, entity := range entities {

			position := entity.GetComponent(&components.Position{}).(*components.Position)
			move := entity.GetComponent(&components.KeyboardMover{}).(*components.KeyboardMover)
			sprite := entity.GetComponent(&components.Sprite{}).(*components.Sprite)
			width, _ := sprite.Image.Size()

			newRight := position.X + move.Speed
			if newRight < (float64(s.width) - float64(width)) {
				position.X = newRight
			}
		}

	}
	return nil
}
