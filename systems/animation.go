package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationSystem struct {
	manager *entities.Manager
}

func NewAnimationSystem(manager *entities.Manager) *AnimationSystem {
	return &AnimationSystem{manager: manager}
}

func (a *AnimationSystem) Draw(image *ebiten.Image) {
	// see sprite renderer
}

func (a *AnimationSystem) Update() error {
	candidates := a.manager.QueryByComponents(&components.Animation{}, &components.Sprite{})

	for _, e := range candidates {
		a := e.GetComponent(&components.Animation{}).(*components.Animation)
		a.Step()

		s := e.GetComponent(&components.Sprite{}).(*components.Sprite)
		s.Image = a.GetCurrentFrame()
	}
	return nil
}
