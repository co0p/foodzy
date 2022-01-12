package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"time"
)

const StartSceneName string = "start"

type StartSceneStates = int

const (
	stateInitial StartSceneStates = iota
	stateAnimateTitle
	stateAnimateMenu
	stateMovePlateLeft
	stateMovePlateRight
	stateMovePlateUp
	stateMovePlateDown
	stateMovePlateStop
)

type StartScene struct {
	scene.GameScene
	entityManager *ecs.EntityManager

	quitAction  ActionType
	startAction ActionType
	transitions []StateTransition
}

func (s *StartScene) Name() string {
	return StartSceneName
}

func NewStartScene(startAction ActionType, quitAction ActionType) *StartScene {
	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewFoodSpawner(40))

	s := &StartScene{
		entityManager: &entityManager,
		startAction:   startAction,
		quitAction:    quitAction,
	}

	now := time.Now()
	transitions := []StateTransition{
		{now.Add(1 * time.Second), func() {
			s.animateTitle()
		}}, {now.Add(3 * time.Second), func() {
			s.addPlayer()
		}}, {now.Add(4 * time.Second), func() {
			s.movePlateLeft()
		}}, {now.Add(5 * time.Second), func() {
			s.movePlateRight()
		}}, {now.Add(6 * time.Second), func() {
			s.movePlateUp()
			s.stopTitle()
		}}, {now.Add(7 * time.Second), func() {
			s.movePlateDown()
		}}, {now.Add(8 * time.Second), func() {
			s.removePlate()
		}}, {now.Add(9 * time.Second), func() {
			s.displayMenu()
		}},
	}
	s.transitions = transitions

	s.Systems = append(s.Systems,
		NewMovementSystem(&entityManager),
		NewFoodSpawningSystem(&entityManager),
		NewSpriteRenderSystem(&entityManager),
		NewInteractionSystem(&entityManager),
		NewTextRenderSystem(&entityManager),
		NewDebugRendererSystem(&entityManager),
		NewCleanupSystem(&entityManager, 100),
	)

	return s
}

func (s *StartScene) Update() error {
	s.processStateTransition()
	return s.GameScene.Update()
}

type StateTransition struct {
	delta      time.Time
	transition func()
}

func (s *StartScene) processStateTransition() {
	if len(s.transitions) == 0 {
		return
	}

	t := s.transitions[0]
	if time.Now().After(t.delta) {
		t.transition()
		s.transitions = s.transitions[1:]
	}
}

func (s *StartScene) animateTitle() {
	title := NewTitle()

	pos := title.GetComponent(component.TransformType).(*component.Transform)
	pos.Y = -200
	velocity := title.GetComponent(component.VelocityType).(*component.Velocity)
	velocity.Y = 1
	s.entityManager.AddEntity(title)
}

func (s *StartScene) addPlayer() {
	player := NewPlayer()
	player.AddComponent(&component.Velocity{})
	s.entityManager.AddEntity(player)
}

func (s *StartScene) stopTitle() {
	title := s.entityManager.QueryFirstByTag("title")
	title.RemoveComponent(component.VelocityType)
}

func (s *StartScene) displayMenu() {
	s.entityManager.AddEntity(NewMenuItem("start", s.startAction, 200))
	s.entityManager.AddEntity(NewMenuItem("quit", s.quitAction, 250))
}

func (s *StartScene) movePlateLeft() {
	plate := s.entityManager.QueryFirstByTag("player")
	mover := plate.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
	velocity := plate.GetComponent(component.VelocityType).(*component.Velocity)
	velocity.X = -0.5 * mover.Speed
	velocity.Y = 0
}

func (s *StartScene) movePlateRight() {
	plate := s.entityManager.QueryFirstByTag("player")
	mover := plate.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
	velocity := plate.GetComponent(component.VelocityType).(*component.Velocity)
	velocity.X = mover.Speed * 0.5
	velocity.Y = 0
}

func (s *StartScene) movePlateUp() {
	plate := s.entityManager.QueryFirstByTag("player")
	mover := plate.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
	velocity := plate.GetComponent(component.VelocityType).(*component.Velocity)
	velocity.X = 0
	velocity.Y = -0.5 * mover.Speed
}

func (s *StartScene) movePlateDown() {
	plate := s.entityManager.QueryFirstByTag("player")
	mover := plate.GetComponent(component.KeyboardMoverType).(*component.KeyboardMover)
	velocity := plate.GetComponent(component.VelocityType).(*component.Velocity)
	velocity.X = 0
	velocity.Y = mover.Speed * 0.5
}

func (s *StartScene) removePlate() {
	plate := s.entityManager.QueryFirstByTag("player")
	plate.Active = false
}
