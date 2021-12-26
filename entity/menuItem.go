package entity

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/utils"
)

func NewMenuStartItem(ScreenWidth int, ScreenHeight int) *Entity {

	menuStartAsset, _ := utils.LoadImage(assets.MenuStart)
	menuStartActiveAsset, _ := utils.LoadImage(assets.MenuStartActive)

	sprite := component.NewSprite("menu_start", assets.MenuStart)
	width, height := sprite.Image.Size()

	mouseCollider := component.MouseCollider{Width: float64(width), Height: float64(height)}
	menuItem := component.MenuItem{
		DefaultSprite: menuStartAsset,
		ActiveSprite:  menuStartActiveAsset,
	}

	xPos := float64((ScreenWidth - width) / 2)
	yPos := float64((ScreenHeight - height) - 200)
	transform := component.Transform{X: xPos, Y: yPos, Z: 1, Scale: 1}

	entity := NewEntity("menu_start", true)
	entity.AddComponent(sprite)
	entity.AddComponent(&mouseCollider)
	entity.AddComponent(&menuItem)
	entity.AddComponent(&transform)

	return entity
}

func NewMenuQuitItem(ScreenWidth int, ScreenHeight int) *Entity {

	menuQuitAsset, _ := utils.LoadImage(assets.MenuQuit)
	menuQuitActiveAsset, _ := utils.LoadImage(assets.MenuQuitActive)

	sprite := component.NewSprite("menu_quit", assets.MenuQuit)
	width, height := sprite.Image.Size()

	mouseCollider := component.MouseCollider{Width: float64(width), Height: float64(height)}
	menuItem := component.MenuItem{
		DefaultSprite: menuQuitAsset,
		ActiveSprite:  menuQuitActiveAsset,
	}

	xPos := float64((ScreenWidth - width) / 2)
	yPos := float64((ScreenHeight - height) - 100)
	transform := component.Transform{X: xPos, Y: yPos, Z: 1, Scale: 1}

	entity := NewEntity("menu_quit", true)
	entity.AddComponent(sprite)
	entity.AddComponent(&mouseCollider)
	entity.AddComponent(&menuItem)
	entity.AddComponent(&transform)

	return entity
}
