package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewTitle() *ecs.Entity {
	sprite := component.NewSprite("title", asset.Title)

	entity := ecs.NewEntity("title", true)
	entity.AddComponent(&component.Velocity{X: 0, Y: 1})
	entity.AddComponent(sprite)

	w, _ := sprite.Image.Size()
	width := float64(ScreenWidth-w) / 2

	entity.AddComponent(&component.Transform{X: width, Y: -150.0, Z: 1, Scale: 1})
	return entity
}
