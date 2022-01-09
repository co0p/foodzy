package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewGameOverTitle() *ecs.Entity {
	entity := ecs.NewEntity("gameovertitle", true)
	velocity := &component.Velocity{X: 0, Y: 1}
	txt := &component.Text{
		Value: "GAME OVER",
		Color: PrimaryColor,
		Font:  &FontHuge,
	}

	posX, _ := txt.RelativeCenter(ScreenWidth, ScreenHeight)
	transform := &component.Transform{X: posX, Y: -150.0, Z: 1, Scale: 1}
	entity.AddComponents(txt, velocity, transform)
	return entity
}
