package main

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/co0p/foodzy/systems"
	"github.com/co0p/foodzy/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const ScreenWidth = 500
const ScreenHeight = 500

func main() {
	sprites, _ := utils.LoadSpriteSheet(SpriteSheet, 200, 80, 8)

	animation := components.NewAnimation(0.2, sprites)
	sprite := components.Sprite{Image: animation.GetCurrentFrame()}
	w, h := sprite.Image.Size()

	position := components.Position{X: (ScreenWidth - float64(w)) / 2, Y: (ScreenHeight - float64(h)) / 2}

	entity := entities.NewEntity("animation", true)
	entity.AddComponent(animation)
	entity.AddComponent(&sprite)
	entity.AddComponent(&position)

	entityManager := entities.Manager{}
	entityManager.AddEntity(entity)

	spriteRenderSystem := systems.NewSpriteRenderSystem(&entityManager)
	animationSystem := systems.NewAnimationSystem(&entityManager)

	example := AnimationExample{
		ScreenHeight:    ScreenHeight,
		ScreenWidth:     ScreenHeight,
		Title:           "Animation example",
		entityManager:   &entityManager,
		animationSystem: animationSystem,
		spriteRenderer:  spriteRenderSystem,
	}

	ebiten.SetWindowSize(example.ScreenWidth, example.ScreenHeight)
	ebiten.SetWindowTitle(example.Title)
	_ = ebiten.RunGame(&example)
}
