package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type InteractionSystem struct {
	manager *ecs.EntityManager
}

func NewInteractionSystem(manager *ecs.EntityManager) *InteractionSystem {
	return &InteractionSystem{manager: manager}
}

func (s *InteractionSystem) Draw(screen *ebiten.Image) {}

func (s *InteractionSystem) Update() error {

	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	mx, my := ebiten.CursorPosition()
	entities := s.manager.QueryByComponents(component.MouseColliderType, component.InteractionType, component.TransformType)

	for _, v := range entities {
		boundingBox := v.GetComponent(component.MouseColliderType).(*component.MouseCollider)
		pos := v.GetComponent(component.TransformType).(*component.Transform)

		if mouseIsInBoundingbox(pos.X, pos.Y, boundingBox.Width, boundingBox.Height, float64(mx), float64(my)) {
			interaction := v.GetComponent(component.InteractionType).(*component.Interaction)
			interaction.Action(s.manager)
		}
	}

	return nil

}

func mouseIsInBoundingbox(px, py, w, h, mx, my float64) bool {
	return mx >= px && mx <= (px+w) && my >= py && my <= (py+h)
}
