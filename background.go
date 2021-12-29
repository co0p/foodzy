package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewBackground() *ecs.Entity {
	bg := ecs.NewEntity("background", true)
	sprite := component.NewSprite("background", asset.Background)

	bg.AddComponent(&component.Transform{X: 0, Y: 0, Scale: 1})
	bg.AddComponent(sprite)
	return bg
}
