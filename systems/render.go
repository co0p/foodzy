package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderSystem struct {
	manager *entities.Manager
}

func NewRenderSystem(manager *entities.Manager) *RenderSystem {
	return &RenderSystem{manager: manager}
}

func (s *RenderSystem) Draw(screen *ebiten.Image) {
	entities := s.manager.QueryByComponents(&components.Position{}, &components.Sprite{})

	for _, entity := range entities {
		position := entity.GetComponent(&components.Position{}).(*components.Position)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.X, position.Y)

		sprite := entity.GetComponent(&components.Sprite{}).(*components.Sprite)
		screen.DrawImage(sprite.Image, op)

	}
}

func (s *RenderSystem) Update() error {
	return nil
}
