package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
)

func NewMenuStartItem(action func()) *ecs.Entity {

	menuStartAsset, _ := asset.LoadImage(asset.MenuStart)
	menuStartActiveAsset, _ := asset.LoadImage(asset.MenuStartActive)

	sprite := component.NewSprite("menu_start", asset.MenuStart)
	width, height := sprite.Image.Size()

	mouseCollider := component.MouseCollider{Width: float64(width), Height: float64(height)}
	menuItem := component.MenuItem{
		DefaultSprite: menuStartAsset,
		ActiveSprite:  menuStartActiveAsset,
		Action:        action,
	}

	xPos := float64((ScreenWidth - width) / 2)
	yPos := float64((ScreenHeight - height) - 200)
	transform := component.Transform{X: xPos, Y: yPos, Z: 1, Scale: 1}

	entity := ecs.NewEntity("menu_start", true)
	entity.AddComponent(sprite)
	entity.AddComponent(&mouseCollider)
	entity.AddComponent(&menuItem)
	entity.AddComponent(&transform)

	return entity
}

func NewMenuQuitItem(action func()) *ecs.Entity {

	menuQuitAsset, _ := asset.LoadImage(asset.MenuQuit)
	menuQuitActiveAsset, _ := asset.LoadImage(asset.MenuQuitActive)

	sprite := component.NewSprite("menu_quit", asset.MenuQuit)
	width, height := sprite.Image.Size()

	mouseCollider := component.MouseCollider{Width: float64(width), Height: float64(height)}
	menuItem := component.MenuItem{
		DefaultSprite: menuQuitAsset,
		ActiveSprite:  menuQuitActiveAsset,
		Action:        action,
	}

	xPos := float64((ScreenWidth - width) / 2)
	yPos := float64((ScreenHeight - height) - 100)
	transform := component.Transform{X: xPos, Y: yPos, Z: 1, Scale: 1}

	entity := ecs.NewEntity("menu_quit", true)
	entity.AddComponent(sprite)
	entity.AddComponent(&mouseCollider)
	entity.AddComponent(&menuItem)
	entity.AddComponent(&transform)

	return entity
}
