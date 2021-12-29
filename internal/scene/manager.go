package scene

import (
	"log"
)

type Action string

const (
	ActionQuit                Action = "quit"
	ActionActivateStartScreen Action = "startscene"
	ActionActivateGameScreen  Action = "gamescene"
)

type SceneManager struct {
	Current Scene
	scenes  map[string]Scene
}

func NewSceneManager() *SceneManager {

	manager := &SceneManager{
		Current: nil,
		scenes:  make(map[string]Scene),
	}

	return manager
}

func (m *SceneManager) AddScreen(s Scene) {
	log.Printf("[SceneManager] adding scene:%s\n", s.Name())
	m.scenes[s.Name()] = s

}

func (m *SceneManager) quit() {
	log.Printf("[SceneManager] quit\n")
}

func (m *SceneManager) ActivateScreen(name string) {
	log.Printf("[SceneManager] activate scene:%s\n", name)

	l, ok := m.scenes[name]

	if !ok {
		panic("could not find scene: " + name)
	}

	if m.Current != nil {
		m.Current.Exit()
	}
	m.Current = l
	l.Init()
}
