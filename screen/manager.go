package screen

import (
	"log"
)

type Action string

const (
	ActionQuit                Action = "quit"
	ActionActivateStartScreen Action = "startscreen"
	ActionActivateGameScreen  Action = "gamescreen"
)

type Manager struct {
	Current Screen
	screens map[string]Screen
}

func NewManager() *Manager {

	manager := &Manager{
		Current: nil,
		screens: make(map[string]Screen),
	}

	return manager
}

func (m *Manager) AddScreen(s Screen) {
	log.Printf("[screen.Manager] adding screen:%s\n", s.Name())
	m.screens[s.Name()] = s

}

func (m *Manager) quit() {
	log.Printf("[screen.Manager] quit\n")
}

func (m *Manager) ActivateScreen(name string) {
	log.Printf("[screen.Manager] activate screen :%s\n", name)

	l, ok := m.screens[name]

	if !ok {
		panic("could not find screen: " + name)
	}

	if m.Current != nil {
		m.Current.Exit()
	}
	m.Current = l
	l.Init()
}
