package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderSystem struct {
	manager *entity.Manager
}

func NewSpriteRenderSystem(manager *entity.Manager) *SpriteRenderSystem {
	return &SpriteRenderSystem{manager: manager}
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {
	candidates := s.manager.QueryByComponents(component.PositionType, component.SpriteType)

	for _, e := range candidates {

		if !e.Active {
			continue
		}

		position := e.GetComponent(component.PositionType).(*component.Position)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)

		sprite := e.GetComponent(component.SpriteType).(*component.Sprite)
		screen.DrawImage(sprite.Image, op)
	}
}

func (s *SpriteRenderSystem) Update() error {
	return nil
}
