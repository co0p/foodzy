package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderSystem struct {
	manager *entities.Manager
}

func NewSpriteRenderSystem(manager *entities.Manager) *SpriteRenderSystem {
	return &SpriteRenderSystem{manager: manager}
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {
	candidates := s.manager.QueryByComponents(&components.Position{}, &components.Dimension{}, &components.Sprite{})

	for _, e := range candidates {

		if !e.Active {
			continue
		}

		position := e.GetComponent(&components.Position{}).(*components.Position)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)

		sprite := e.GetComponent(&components.Sprite{}).(*components.Sprite)
		screen.DrawImage(sprite.Image, op)
	}
}

func (s *SpriteRenderSystem) Update() error {
	return nil
}
