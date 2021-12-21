package main

import (
	"github.com/co0p/foodzy/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const ScreenWidth = 500
const ScreenHeight = 500

type ComponentType string

const (
	Transform  ComponentType = "TRANSFORM"
	SpriteType               = "SPRITE"
	Animation                = "ANIMATION"
)

type AnimationComponent struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	CurrentFrame      *ebiten.Image
	Count             float64
	AnimationSpeed    float64
}

func (a AnimationComponent) Type() ComponentType { return Animation }

func main() {
	sprites, _ := utils.LoadSpriteSheet(SpriteSheet, 200, 80, 8)

	animation := component.NewAnimation(0.2, sprites)
	sprite := component.Sprite{Image: animation.GetCurrentFrame()}
	w, h := sprite.Image.Size()

	position := component.Position{X: (ScreenWidth - float64(w)) / 2, Y: (ScreenHeight - float64(h)) / 2}

	entity := entity.NewEntity("animation", true)
	entity.AddComponent(animation)
	entity.AddComponent(&sprite)
	entity.AddComponent(&position)

	entityManager := entity.Manager{}
	entityManager.AddEntity(entity)

	spriteRenderSystem := system.NewSpriteRenderSystem(&entityManager)
	animationSystem := system.NewAnimationSystem(&entityManager)

	example := AnimationExample{
		// AnimationExample satisfies the ebiten.Game interface
		entityManager:   &entityManager,
		animationSystem: animationSystem,
		spriteRenderer:  spriteRenderSystem,
	}

	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Animation example")
	if err := ebiten.RunGame(&example); err != nil {
		panic(err)
	}
}
