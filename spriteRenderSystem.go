package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderSystem struct {
	manager *ecs.EntityManager
}

func NewSpriteRenderSystem(manager *ecs.EntityManager) *SpriteRenderSystem {
	return &SpriteRenderSystem{manager: manager}
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {
	candidates := s.manager.QueryByComponents(component.TransformType, component.SpriteType)

	var foreground []*ecs.Entity
	var background []*ecs.Entity
	for _, e := range candidates {

		if !e.Active {
			continue
		}

		transform := e.GetComponent(component.TransformType).(*component.Transform)
		if transform.Z > 0 {
			foreground = append(foreground, e)
		} else {
			background = append(background, e)
		}
	}

	orderedEntities := append(background, foreground...)
	for _, e := range orderedEntities {

		transform := e.GetComponent(component.TransformType).(*component.Transform)
		sprite := e.GetComponent(component.SpriteType).(*component.Sprite)
		if sprite.Image != nil {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(transform.Scale, transform.Scale)
			op.GeoM.Translate(transform.X, transform.Y)

			screen.DrawImage(sprite.Image, op)
		}
	}
}

func (s *SpriteRenderSystem) Update() error {
	return nil
}
