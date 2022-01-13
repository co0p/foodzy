package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameoverSystem struct {
	entityManager *ecs.EntityManager
	soundManager  *sound.SoundManager
	count         int
	startAction   ActionType
	original      *ebiten.Image
}

func NewGameoverSystem(manager *ecs.EntityManager, soundManager *sound.SoundManager, startAction ActionType) *GameoverSystem {
	return &GameoverSystem{
		entityManager: manager,
		soundManager:  soundManager,
		startAction:   startAction,
	}
}

func (s *GameoverSystem) Update() error {
	// decrease sound level, until exit
	if s.count%ebiten.MaxTPS() == 0 {
		s.soundManager.Volume(SoundBackground, -0.1)
	}
	vol := s.soundManager.Volume(SoundBackground, 0)
	if vol < 0 {
		s.startAction(s.entityManager)
	}
	s.count++

	// title animation
	title := s.entityManager.QueryFirstByTag("gameovertitle")
	titleVelocity := title.GetComponent(component.VelocityType).(*component.Velocity)
	titleTransform := title.GetComponent(component.TransformType).(*component.Transform)

	// stop animation of title
	if titleVelocity.Y != 0 && titleTransform.Y > float64(ScreenHeight/2) {
		titleVelocity.Y = 0.0
	}

	return nil
}

func (s *GameoverSystem) Draw(screen *ebiten.Image) {
	if s.original == nil {
		s.original = ebiten.NewImageFromImage(screen)
	}

	for j := -3; j <= 3; j++ {
		for i := -3; i <= 3; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i), float64(j))
			op.ColorM.Scale(1, 1, 1, 1.0/25.0)
			screen.DrawImage(s.original, op)
		}
	}
}
