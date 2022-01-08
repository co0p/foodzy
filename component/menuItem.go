package component

import (
	"github.com/co0p/foodzy/internal/ecs"
	"image/color"
)

type MenuItem struct {
	Action         func()
	Text           *Text
	DefaultColor   color.Color
	HighlightColor color.Color
}

const MenuItemType ecs.ComponentType = "MenuItem"

func (c *MenuItem) Type() ecs.ComponentType {
	return MenuItemType
}
