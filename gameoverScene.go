package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
)

const GameOverSceneName = "gameover"

type GameOverScene struct {
	scene.GameScene
}

func NewGameOverScene(soundManager *sound.SoundManager, actionQuit ActionType) *GameOverScene {

	entityManager := ecs.EntityManager{}

	s := &GameOverScene{}

	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewFoodSpawner(40))
	entityManager.AddEntity(NewGameOverTitle())

	s.Systems = append(s.Systems,
		NewMovementSystem(&entityManager),
		NewFoodSpawningSystem(&entityManager),
		NewSpriteRenderSystem(&entityManager),
		NewGameoverSystem(&entityManager, soundManager, actionQuit),
		NewTextRenderSystem(&entityManager),
		NewDebugRendererSystem(&entityManager),
		NewCleanupSystem(&entityManager, 10),
	)

	return s
}

func (s *GameOverScene) Name() string {
	return GameOverSceneName
}
