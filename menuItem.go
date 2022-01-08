package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

func NewMenuItem(s string, f func(*ecs.EntityManager), offset float64) *ecs.Entity {
	txt := &component.Text{Value: s, Font: &FontBig, Color: color.RGBA{R: 0, G: 20, B: 27, A: 200}}
	w, h := txt.Dimensions()
	collider := &component.MouseCollider{Width: w, Height: h}
	transform := &component.Transform{X: 200, Y: offset, Z: 1, Scale: 1}
	interaction := &component.Interaction{Action: f}

	entity := ecs.NewEntity(s, true)
	entity.AddComponents(interaction, transform, collider, txt)
	return entity
}
