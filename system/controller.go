package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type ControllerSystem struct {
	manager *entity.Manager
	height  int
	width   int
}

func NewControllerSystem(manager *entity.Manager, width int, height int) *ControllerSystem {
	return &ControllerSystem{manager: manager, width: width, height: height}
}

func (s *ControllerSystem) Draw(image *ebiten.Image) { /* nothing to do */ }

func (s *ControllerSystem) Update() error {

	entities := s.manager.QueryByComponents(component.KeyboardMoverType, component.SpriteType, component.TransformType)

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

		for _, e := range entities {

			position := e.GetComponent(component.TransformType).(*component.Transform)
			move := e.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)

			newLeft := position.X - move.Speed
			if newLeft > 0 {
				position.X = newLeft
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		for _, e := range entities {

			position := e.GetComponent(component.TransformType).(*component.Transform)
			move := e.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
			sprite := e.GetComponent(component.SpriteType).(*component.Sprite)
			width, _ := sprite.Image.Size()

			newRight := position.X + move.Speed
			if newRight < (float64(s.width) - float64(width)) {
				position.X = newRight
			}
		}

	}
	return nil
}
