package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

func NewTitle() *ecs.Entity {
	text := &component.Text{
		Value: "Foodzy",
		Color: color.RGBA{R: 0, G: 20, B: 27, A: 255},
		Font:  &FontHuge,
	}
	transform := &component.Transform{X: 200, Y: 150, Z: 1, Scale: 1}

	entity := ecs.NewEntity("title", true)
	entity.AddComponents(transform, text)
	return entity
}
