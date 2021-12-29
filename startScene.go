package foodzy

import (
	ecs2 "github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const StartScreenName string = "start"

type StartScreen struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *ecs2.Manager
	systems       []ecs2.System
	actions       map[scene.Action]func()
	initialized   bool
}

func NewStartScreen() *StartScreen {
	return &StartScreen{
		initialized:   false,
		actions:       make(map[scene.Action]func()),
		entityManager: &ecs2.Manager{},
		systems:       []ecs2.System{},
	}
}

func (s *StartScreen) Name() string {
	return StartScreenName
}

func (s *StartScreen) Init() {

	if s.initialized {
		log.Printf("[screen:%s] resuming\n", s.Name())
		return
	}

	log.Printf("[screen:%s] initializing\n", s.Name())

	s.entityManager.AddEntity(NewBackground())
	s.entityManager.AddEntity(NewFoodSpawner(40))
	s.entityManager.AddEntity(NewTitle())

	startMenuItem := NewMenuStartItem(s.actions[scene.ActionActivateGameScreen])
	quitMenuItem := NewMenuQuitItem(s.actions[scene.ActionQuit])
	items := []*ecs2.Entity{startMenuItem, quitMenuItem}

	s.systems = append(s.systems,
		NewMenuSystem(s.entityManager, items),
		NewMovementSystem(s.entityManager),
		NewFoodSpawningSystem(s.entityManager),
		NewSpriteRenderSystem(s.entityManager),
		NewCleanupSystem(s.entityManager, 100),
	)

	s.initialized = true
}

func (s *StartScreen) Exit() {
	log.Printf("[screen:%s] exit\n", s.Name())
}

func (s *StartScreen) Update() error {

	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (s *StartScreen) Draw(screen *ebiten.Image) {
	for _, sys := range s.systems {
		sys.Draw(screen)
	}
}

func (s *StartScreen) AddAction(action scene.Action, callback func()) {
	s.actions[action] = callback
}
