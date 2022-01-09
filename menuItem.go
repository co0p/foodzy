package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewMenuItem(s string, f func(*ecs.EntityManager), offset float64) *ecs.Entity {
	txt := &component.Text{Value: s, Font: &FontBig, Color: SecondaryColor}
	w, h := txt.Dimensions()
	collider := &component.MouseCollider{Width: w, Height: h}
	transform := &component.Transform{X: 200, Y: offset, Z: 1, Scale: 1}
	interaction := &component.Interaction{Action: f}

	entity := ecs.NewEntity(s, true)
	entity.AddComponents(interaction, transform, collider, txt)
	return entity
}
