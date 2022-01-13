package foodzy

import (
	"fmt"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"image/color"
)

var mouseDebugColor = color.RGBA{0, 255, 255, 150}
var collisionDebugColor = color.RGBA{178, 34, 34, 150}

type DebugRendererSystem struct {
	manager *ecs.EntityManager
	enabled bool
}

func NewDebugRendererSystem(manager *ecs.EntityManager) *DebugRendererSystem {
	return &DebugRendererSystem{manager: manager}
}

func (s *DebugRendererSystem) Draw(screen *ebiten.Image) {

	if !s.enabled {
		return
	}

	mouseEntities := s.manager.QueryByComponents(component.MouseColliderType, component.TransformType)
	for _, v := range mouseEntities {
		boundingBox := v.GetComponent(component.MouseColliderType).(*component.MouseCollider)
		transform := v.GetComponent(component.TransformType).(*component.Transform)

		img := ebiten.NewImage(int(boundingBox.Width), int(boundingBox.Height))
		img.Fill(mouseDebugColor)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(transform.Scale, transform.Scale)
		op.GeoM.Translate(transform.X, transform.Y)

		screen.DrawImage(img, op)
	}

	collisionEntities := s.manager.QueryByComponents(component.CollisionType, component.TransformType)
	for _, v := range collisionEntities {
		boundingBox := v.GetComponent(component.CollisionType).(*component.Collision)
		transform := v.GetComponent(component.TransformType).(*component.Transform)

		img := ebiten.NewImage(int(boundingBox.Width), int(boundingBox.Height))
		img.Fill(collisionDebugColor)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(transform.Scale, transform.Scale)
		op.GeoM.Translate(transform.X, transform.Y)

		screen.DrawImage(img, op)
	}

	fps := fmt.Sprintf("FPS: %.1f", ebiten.CurrentFPS())
	text.Draw(screen, fps, FontSmall, 10, 50, colornames.Firebrick)
	tps := fmt.Sprintf("FPS: %.1f", ebiten.CurrentFPS())
	text.Draw(screen, tps, FontSmall, 100, 50, colornames.Firebrick)
}

func (s *DebugRendererSystem) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		s.enabled = !s.enabled
	}

	return nil
}
