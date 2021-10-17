package entities

import (
	"github.com/co0p/foodzy/components"
	"image/color"
)

func NewScore(text string, nutrient components.Nutrient, posX float64, posY float64) *Entity {
	bg := NewEntity("score", true)

	bg.AddComponent(&components.Position{X: posX, Y: posY})
	bg.AddComponent(&components.Text{Value: text, Color: color.White})
	bg.AddComponent(&nutrient)

	return bg
}
