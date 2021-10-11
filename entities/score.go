package entities

import (
	"github.com/co0p/foodzy/components"
)

func NewScore(text string, nutrient components.Nutrient) *Entity {
	bg := NewEntity("score", true)

	bg.AddComponent(&components.Text{Text: text})
	bg.AddComponent(&nutrient)

	return bg
}
