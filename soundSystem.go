package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
)

type SoundSystem struct {
	entityManager *ecs.Manager
	soundManager  *sound.SoundManager
}

func NewSoundSystem(entityManager *ecs.Manager, soundManager *sound.SoundManager) *SoundSystem {
	return &SoundSystem{entityManager: entityManager, soundManager: soundManager}
}

func (s *SoundSystem) Update() error {
	entities := s.entityManager.QueryByComponents(component.SoundType)

	for _, e := range entities {
		soundEffect := e.GetComponent(component.SoundType).(*component.Sound)

		s.soundManager.Play(soundEffect.Clip)
		e.RemoveComponent(soundEffect)
	}

	return nil
}

func (s *SoundSystem) Draw(image *ebiten.Image) { /* nothing to do */ }
