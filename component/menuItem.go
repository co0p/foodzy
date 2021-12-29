package component

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuItem struct {
	DefaultSprite *ebiten.Image
	ActiveSprite  *ebiten.Image
	Action        func()
}

const MenuItemType ecs.ComponentType = "MenuItem"

func (c *MenuItem) Type() ecs.ComponentType {
	return MenuItemType
}
