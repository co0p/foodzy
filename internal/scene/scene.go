package scene

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Start()
	Stop()
	Update() error
	Draw(screen *ebiten.Image)
	Name() string
}

type GameScene struct {
	Systems []ecs.System
	running bool
}

func (g *GameScene) Start() {
	g.running = true
}

func (g *GameScene) Stop() {
	g.running = false
}

func (g *GameScene) Update() error {
	if !g.running {
		return nil
	}

	var err error
	for _, v := range g.Systems {
		err = v.Update()
	}
	return err
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	if !g.running {
		return
	}

	for _, v := range g.Systems {
		v.Draw(screen)
	}
}
