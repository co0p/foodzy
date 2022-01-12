package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

func NewTitle() *ecs.Entity {
	txt := &component.Text{
		Value: "Foodzy",
		Color: color.RGBA{R: 0, G: 20, B: 27, A: 255},
		Font:  &FontHuge,
	}
	posX, _ := txt.RelativeCenter(ScreenWidth, ScreenHeight)
	transform := &component.Transform{X: posX, Y: 50, Z: 1, Scale: 1}
	velocity := &component.Velocity{Y: 5}

	entity := ecs.NewEntity("title", true)
	entity.AddComponents(transform, velocity, txt)
	return entity
}
