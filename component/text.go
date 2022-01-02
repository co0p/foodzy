package component

import (
	"github.com/co0p/foodzy/internal/ecs"
	"golang.org/x/image/font"
	"image/color"
)

type Text struct {
	Value string
	Color color.Color
	Font  *font.Face
}

const TextType ecs.ComponentType = "Text"

func (t Text) Type() ecs.ComponentType {
	return TextType
}
