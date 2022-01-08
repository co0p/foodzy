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
	exitAction    func(*ecs.EntityManager)
}

func NewGameoverSystem(manager *ecs.EntityManager, soundManager *sound.SoundManager, exitAction func(*ecs.EntityManager)) *GameoverSystem {
	return &GameoverSystem{
		entityManager: manager,
		soundManager:  soundManager,
		exitAction:    exitAction,
	}
}

func (s *GameoverSystem) Update() error {
	// decrease sound level, until exit
	if s.count%ebiten.MaxTPS() == 0 {
		s.soundManager.Volume(SoundBackground, -0.1)
	}
	vol := s.soundManager.Volume(SoundBackground, 0)
	if vol < 0 {
		s.exitAction(s.entityManager)
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

func (s *GameoverSystem) Draw(image *ebiten.Image) { /* nothing to do */ }
