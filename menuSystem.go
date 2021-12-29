package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuSystem struct {
	entityManager *ecs.Manager
	items         []*ecs.Entity
}

func NewMenuSystem(manager *ecs.Manager, items []*ecs.Entity) *MenuSystem {
	return &MenuSystem{
		entityManager: manager,
		items:         items,
	}
}

func (s *MenuSystem) Update() error {

	// title animation
	maxTitleHeight := 100.0
	title := s.entityManager.QueryFirstByTag("title")

	titleVelocity := title.GetComponent(component.VelocityType).(*component.Velocity)
	titleTransform := title.GetComponent(component.TransformType).(*component.Transform)

	// stop animation of title, display menu
	if titleVelocity.Y != 0 && titleTransform.Y > maxTitleHeight {
		titleVelocity.Y = 0.0

		s.entityManager.AddEntity(s.items[0])
		s.entityManager.AddEntity(s.items[1])
	}

	// highlight menu item
	mx, my := ebiten.CursorPosition()
	menuItems := s.entityManager.QueryByComponents(component.MouseColliderType, component.TransformType, component.SpriteType, component.MenuItemType)
	for _, v := range menuItems {
		pos := v.GetComponent(component.TransformType).(*component.Transform)
		boundingBox := v.GetComponent(component.MouseColliderType).(*component.MouseCollider)
		itemsSprite := v.GetComponent(component.SpriteType).(*component.Sprite)
		menuItem := v.GetComponent(component.MenuItemType).(*component.MenuItem)

		if mouseIsInBoundingbox(pos.X, pos.Y, boundingBox.Width, boundingBox.Height, float64(mx), float64(my)) {
			itemsSprite.Image = menuItem.ActiveSprite
		} else {
			itemsSprite.Image = menuItem.DefaultSprite
		}

		// handle click
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if mouseIsInBoundingbox(pos.X, pos.Y, boundingBox.Width, boundingBox.Height, float64(mx), float64(my)) {
				menuItem.Action()
			}
		}
	}

	return nil
}

func (s *MenuSystem) Draw(image *ebiten.Image) { /* nothing to do */ }

func mouseIsInBoundingbox(px, py, w, h, mx, my float64) bool {
	return mx >= px && mx <= (px+w) && my >= py && my <= (py+h)
}
