package entity

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
)

const playerSize = 50

func NewPlayer(ScreenWidth int, ScreenHeight int) *entity {
	entity := NewEntity("player", true)

	entity.AddComponent(&component.Transform{
		X:     float64(ScreenWidth / 2.0),
		Y:     float64(ScreenHeight - playerSize*2.0),
		Scale: 1,
	})
	sprite := component.NewSprite("player", assets.Plate)
	width, height := sprite.Image.Size()
	entity.AddComponent(sprite)
	entity.AddComponent(&component.KeyboardMover{Speed: 5.0})

	entity.AddComponent(&component.Collision{Width: float64(width), Height: float64(height / 10)})
	entity.AddComponent(&component.Nutrient{})
	entity.AddComponent(&component.Consumption{
		Corn:      component.ConsumptionDefaultRate,
		Dairy:     component.ConsumptionDefaultRate,
		Drink:     component.ConsumptionDefaultRate,
		Fish:      component.ConsumptionDefaultRate,
		Meat:      component.ConsumptionDefaultRate,
		Treat:     component.ConsumptionDefaultRate,
		Fruit:     component.ConsumptionDefaultRate,
		Vegetable: component.ConsumptionDefaultRate,
		KCal:      component.ConsumptionDefaultRate,
	})

	return entity
}
