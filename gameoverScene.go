package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"time"
)

const GameOverSceneName = "gameover"

type GameOverScene struct {
	scene.GameScene
	startTime     time.Time
	actionStart   ActionType
	entityManager *ecs.EntityManager
}

func NewGameOverScene(soundManager *sound.SoundManager, actionStart ActionType) *GameOverScene {

	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewFoodSpawner(40))
	entityManager.AddEntity(NewGameOverTitle())

	s := &GameOverScene{
		startTime:     time.Now(),
		actionStart:   actionStart,
		entityManager: &entityManager,
	}

	s.Systems = append(s.Systems,
		NewMovementSystem(&entityManager),
		NewFoodSpawningSystem(&entityManager),
		NewSpriteRenderSystem(&entityManager),
		NewGameoverSystem(&entityManager, soundManager, actionStart),
		NewTextRenderSystem(&entityManager),
		NewDebugRendererSystem(&entityManager),
		NewCleanupSystem(&entityManager, 10),
	)

	return s
}

func (s *GameOverScene) Name() string {
	return GameOverSceneName
}
