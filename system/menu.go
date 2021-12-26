package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type MenuSystem struct {
	manager      *entity.Manager
	screenWidth  int
	screenHeight int
}

func NewMenuSystem(manager *entity.Manager, width, height int) *MenuSystem {
	return &MenuSystem{
		manager:      manager,
		screenWidth:  width,
		screenHeight: height,
	}
}

func (s *MenuSystem) Update() error {

	// title animation
	maxTitleHeight := float64(s.screenHeight / 4)
	title := s.manager.QueryFirstByTag("title")

	titleVelocity := title.GetComponent(component.VelocityType).(*component.Velocity)
	titleTransform := title.GetComponent(component.TransformType).(*component.Transform)

	// stop animation of title, display menu
	if titleVelocity.Y != 0 && titleTransform.Y > maxTitleHeight {
		titleVelocity.Y = 0.0

		s.manager.AddEntity(entity.NewMenuStartItem(s.screenWidth, s.screenHeight))
		s.manager.AddEntity(entity.NewMenuQuitItem(s.screenWidth, s.screenHeight))
	}

	// highlight menu item
	mx, my := ebiten.CursorPosition()
	menuItems := s.manager.QueryByComponents(component.MouseColliderType, component.TransformType, component.SpriteType, component.MenuItemType)
	for _, v := range menuItems {
		pos := v.GetComponent(component.TransformType).(*component.Transform)
		boundingBox := v.GetComponent(component.MouseColliderType).(*component.MouseCollider)
		sprite := v.GetComponent(component.SpriteType).(*component.Sprite)
		mc := v.GetComponent(component.MenuItemType).(*component.MenuItem)

		if mouseIsInBoundingbox(pos.X, pos.Y, boundingBox.Width, boundingBox.Height, float64(mx), float64(my)) {
			sprite.Image = mc.ActiveSprite
		} else {
			sprite.Image = mc.DefaultSprite
		}
	}

	return nil
}

func (s *MenuSystem) Draw(image *ebiten.Image) { /* nothing to do */ }

func mouseIsInBoundingbox(px, py, w, h, mx, my float64) bool {
	return mx >= px && mx <= (px+w) && my >= py && my <= (py+h)
}
