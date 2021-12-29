package component

import (
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

type Text struct {
	Value string
	Color color.Color
}

const TextType ecs.ComponentType = "Text"

func (t Text) Type() ecs.ComponentType {
	return TextType
}
