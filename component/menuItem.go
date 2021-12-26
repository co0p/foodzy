package component

import "github.com/hajimehoshi/ebiten/v2"

type MenuItem struct {
	DefaultSprite *ebiten.Image
	ActiveSprite  *ebiten.Image
}

const MenuItemType ComponentType = "MenuItem"

func (c *MenuItem) Type() ComponentType {
	return MenuItemType
}