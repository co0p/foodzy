package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

const playerSize = 50

func NewPlayer(ScreenWidth int, ScreenHeight int) *ecs.Entity {

	keyboardMover := &component.KeyboardMover{Speed: 4.5}
	transform := &component.Transform{
		X:     float64(ScreenWidth / 2.0),
		Y:     float64(ScreenHeight - playerSize*2.0),
		Scale: 1,
	}
	sprite := component.NewSprite("player", asset.Plate)
	width, height := sprite.Image.Size()
	collision := &component.Collision{Width: float64(width), Height: float64(height)}

	entity := ecs.NewEntity("player", true)
	entity.AddComponents(collision, sprite, keyboardMover, transform)
	return entity
}
