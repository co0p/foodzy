package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const GameOverSceneName = "gameover"

type GameOverScene struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *ecs.EntityManager
	soundManager  *sound.SoundManager
	systems       []ecs.System
	actions       map[scene.Action]func()
	initialized   bool
}

func NewGameOverScene(manager *sound.SoundManager) *GameOverScene {
	return &GameOverScene{
		initialized:   false,
		actions:       make(map[scene.Action]func()),
		entityManager: &ecs.EntityManager{},
		soundManager:  manager,
		systems:       []ecs.System{},
	}
}

func (s *GameOverScene) Name() string {
	return GameOverSceneName
}

func (s *GameOverScene) Init() {

	if s.initialized {
		log.Printf("[screen:%s] resuming\n", s.Name())
		return
	}

	log.Printf("[screen:%s] initializing\n", s.Name())

	s.entityManager.AddEntity(NewBackground())
	s.entityManager.AddEntity(NewFoodSpawner(40))
	s.entityManager.AddEntity(NewGameOverTitle())

	s.systems = append(s.systems,
		NewMovementSystem(s.entityManager),
		NewFoodSpawningSystem(s.entityManager),
		NewSpriteRenderSystem(s.entityManager),
		NewTextRenderSystem(s.entityManager),
		NewGameoverSystem(s.entityManager, s.soundManager, s.actions[ActionQuit]),
		NewCleanupSystem(s.entityManager, 10),
	)

	s.initialized = true
}

func (s *GameOverScene) Exit() {
	log.Printf("[screen:%s] exit\n", s.Name())
}

func (s *GameOverScene) Update() error {
	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	for _, sys := range s.systems {
		sys.Draw(screen)
	}
}

func (s *GameOverScene) AddAction(action scene.Action, callback func()) {
	s.actions[action] = callback
}
