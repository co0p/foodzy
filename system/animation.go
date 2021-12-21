package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationSystem struct {
	manager *entity.Manager
}

func NewAnimationSystem(manager *entity.Manager) *AnimationSystem {
	return &AnimationSystem{manager: manager}
}

func (a *AnimationSystem) Draw(image *ebiten.Image) {
	// see sprite renderer
}

func (a *AnimationSystem) Update() error {
	candidates := a.manager.QueryByComponents(component.AnimationType, component.SpriteType)

	for _, e := range candidates {
		a := e.GetComponent(component.AnimationType).(*component.Animation)
		a.Step()

		s := e.GetComponent(component.SpriteType).(*component.Sprite)
		s.Image = a.GetCurrentFrame()
	}
	return nil
}
