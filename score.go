package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

func NewHealth() *ecs.Entity {

	e := ecs.NewEntity("score", true)

	transform := &component.Transform{
		X: float64(ScreenWidth - 100),
		Y: 50,
	}
	text := &component.Text{Value: "", Color: color.White, Font: &FontMedium}

	e.AddComponents(text, transform)
	return e
}
