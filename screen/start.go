package screen

import (
	"github.com/co0p/foodzy/entity"
	"github.com/co0p/foodzy/system"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const StartScreenName string = "start"

type StartScreen struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *entity.Manager
	systems       []system.System
	actions       map[Action]func()
	initialized   bool
}

func NewStartScreen(ScreenWidth int, ScreenHeight int) *StartScreen {
	return &StartScreen{
		initialized:   false,
		actions:       make(map[Action]func()),
		ScreenWidth:   ScreenWidth,
		ScreenHeight:  ScreenHeight,
		entityManager: &entity.Manager{},
		systems:       []system.System{},
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

	s.entityManager.AddEntity(entity.NewBackground())
	s.entityManager.AddEntity(entity.NewFoodSpawner(40))
	s.entityManager.AddEntity(entity.NewTitle(s.ScreenWidth, s.ScreenHeight))

	startMenuItem := entity.NewMenuStartItem(s.ScreenWidth, s.ScreenHeight, s.actions[ActionActivateGameScreen])
	quitMenuItem := entity.NewMenuQuitItem(s.ScreenWidth, s.ScreenHeight, s.actions[ActionQuit])
	items := []*entity.Entity{startMenuItem, quitMenuItem}

	s.systems = append(s.systems,
		system.NewMenuSystem(s.entityManager, items),
		system.NewMovementSystem(s.entityManager),
		system.NewFoodSpawningSystem(s.entityManager, s.ScreenWidth),
		system.NewSpriteRenderSystem(s.entityManager),
		system.NewCleanupSystem(s.entityManager, 100, s.ScreenHeight),
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

func (s *StartScreen) AddAction(action Action, callback func()) {
	s.actions[action] = callback
}
