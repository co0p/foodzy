package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
)

const StartSceneName string = "start"

type StartScene struct {
	scene.GameScene

	entityManager *ecs.EntityManager
}

func (s *StartScene) Name() string {
	return StartSceneName
}

func NewStartScene(startAction ActionType, quitAction ActionType) *StartScene {

	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewFoodSpawner(40))
	entityManager.AddEntity(NewTitle())
	entityManager.AddEntity(NewMenuItem("start", startAction, 200))
	entityManager.AddEntity(NewMenuItem("quit", quitAction, 250))

	s := &StartScene{}
	s.Systems = append(s.Systems,
		NewMovementSystem(&entityManager),
		NewFoodSpawningSystem(&entityManager),
		NewSpriteRenderSystem(&entityManager),
		NewInteractionSystem(&entityManager),
		NewTextRenderSystem(&entityManager),
		NewCleanupSystem(&entityManager, 100),
	)

	return s
}
