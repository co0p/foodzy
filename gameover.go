package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

func NewGameOverTitle() *ecs.Entity {
	entity := ecs.NewEntity("gameovertitle", true)
	text := &component.Text{
		Value: "GAME OVER",
		Color: color.RGBA{R: 0, G: 20, B: 27, A: 200},
		Font:  &FontHuge,
	}
	velocity := &component.Velocity{X: 0, Y: 1}
	transform := &component.Transform{X: 150, Y: -150.0, Z: 1, Scale: 1}
	entity.AddComponents(text, velocity, transform)
	return entity
}
