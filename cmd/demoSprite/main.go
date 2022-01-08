package main

import (
	"github.com/co0p/foodzy"
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteDemo struct {
	renderSystem *foodzy.SpriteRenderSystem
}

func (g *SpriteDemo) Update() error {
	return g.renderSystem.Update()
}

func (g *SpriteDemo) Draw(screen *ebiten.Image) {
	g.renderSystem.Draw(screen)
}

func (g *SpriteDemo) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 400
}

func main() {

	// E - Entity
	entityManager := ecs.EntityManager{}
	food := ecs.NewEntity("food", true)
	foodBig := ecs.NewEntity("foodBig", true)
	entityManager.AddEntity(food)
	entityManager.AddEntity(foodBig)

	// C - components
	sprite := component.NewSprite("food", asset.Fruit_strawberry)
	transform := component.Transform{X: 100, Y: 100, Z: 1, Scale: 1}
	food.AddComponents(&transform, sprite)

	sprite2 := component.NewSprite("food", asset.Fruit_strawberry)
	transform2 := component.Transform{X: 250, Y: 100, Z: 1, Scale: 2}
	foodBig.AddComponents(&transform2, sprite2)

	// S - System
	spriteRenderSystem := foodzy.NewSpriteRenderSystem(&entityManager)

	game := SpriteDemo{renderSystem: spriteRenderSystem}

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
