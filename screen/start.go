package screen

import (
	"github.com/co0p/foodzy/entity"
	"github.com/co0p/foodzy/system"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const StartScreenName Name = "start"

type StartScreen struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *entity.Manager
	systems       []system.System
}

func NewStartScreen(ScreenWidth int, ScreenHeight int) *StartScreen {
	return &StartScreen{

		ScreenWidth:   ScreenWidth,
		ScreenHeight:  ScreenHeight,
		entityManager: &entity.Manager{},
		systems:       []system.System{},
	}
}

func (s *StartScreen) Name() Name {
	return StartScreenName
}

func (s *StartScreen) Init() {
	log.Printf("[screen:%s] initializing\n", s.Name())

	s.entityManager.AddEntity(entity.NewBackground())
	s.entityManager.AddEntity(entity.NewFoodSpawner(40))
	s.entityManager.AddEntity(entity.NewTitle(s.ScreenWidth, s.ScreenHeight))

	s.systems = append(s.systems,
		system.NewMenuSystem(s.entityManager, s.ScreenWidth, s.ScreenHeight),
		system.NewMovementSystem(s.entityManager),
		system.NewSoundSystem(),
		system.NewFoodSpawningSystem(s.entityManager, s.ScreenWidth),
		system.NewSpriteRenderSystem(s.entityManager),
		system.NewCleanupSystem(s.entityManager, 100, s.ScreenHeight),
	)
}

func (s *StartScreen) Exit() {
	log.Printf("[screen:%s] exit\n", s.Name())
}

func (s StartScreen) Update() error {

	for _, sys := range s.systems {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (s StartScreen) Draw(screen *ebiten.Image) {
	for _, sys := range s.systems {
		sys.Draw(screen)
	}
}
