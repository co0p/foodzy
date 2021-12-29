package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type ControllerSystem struct {
	manager *ecs.Manager
}

func NewControllerSystem(manager *ecs.Manager) *ControllerSystem {
	return &ControllerSystem{manager: manager}
}

func (s *ControllerSystem) Draw(image *ebiten.Image) {}

func (s *ControllerSystem) Update() error {

	entities := s.manager.QueryByComponents(component.KeyboardMoverType, component.SpriteType, component.TransformType)

	if ebiten.IsKeyPressed(ebiten.KeyUp) {

		for _, e := range entities {

			position := e.GetComponent(component.TransformType).(*component.Transform)
			move := e.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)

			newUp := position.Y - move.Speed
			if newUp > 0 {
				position.Y = newUp
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {

		for _, e := range entities {

			position := e.GetComponent(component.TransformType).(*component.Transform)
			move := e.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
			sprite := e.GetComponent(component.SpriteType).(*component.Sprite)

			_, height := sprite.Image.Size()

			newDown := position.Y + move.Speed
			if newDown < (float64(ScreenHeight) - float64(height)) {
				position.Y = newDown
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
			if newRight < (float64(ScreenWidth) - float64(width)) {
				position.X = newRight
			}
		}

	}

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

	return nil
}
