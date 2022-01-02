package scene

import (
	"log"
)

type Action string

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

func (m *SceneManager) ActivateScreen(name string, pause bool) {
	log.Printf("[SceneManager] activate scene:%s\n", name)

	l, ok := m.scenes[name]

	if !ok {
		panic("could not find scene: " + name)
	}

	if m.Current != nil && !pause {
		m.Current.Exit()
	}
	m.Current = l
	l.Init()
}
