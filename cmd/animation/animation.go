package main

import (
	"github.com/co0p/foodzy/entities"
	"github.com/co0p/foodzy/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

// Asset from https://luizmelo.itch.io/martial-hero/ ... run animation
//go:generate file2byteslice -package=main -input=spriteSheet.png -output=spriteSheet.go -var=SpriteSheet

// AnimationExample is a showcase of how to create an animation by creating an entity
// with the animation component attached and running the animation system in conjunction with the sprite render system
type AnimationExample struct {
	ScreenWidth, ScreenHeight int
	Title                     string
	entityManager             *entities.Manager
	animationSystem           *systems.AnimationSystem
	spriteRenderer            *systems.SpriteRenderSystem
}

func (a *AnimationExample) Update() error {
	return a.animationSystem.Update()
}

func (a *AnimationExample) Draw(screen *ebiten.Image) {
	a.spriteRenderer.Draw(screen)
}

func (a *AnimationExample) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return a.ScreenWidth, a.ScreenHeight
}
