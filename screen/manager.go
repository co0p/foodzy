package screen

import "log"

type Manager struct {
	Current Screen
	screens map[string]Screen
}

func NewManager() *Manager {
	return &Manager{
		Current: nil,
		screens: make(map[string]Screen),
	}
}

func (m *Manager) AddScreen(s Screen) {
	log.Printf("[screen.Manager] activating screen:%s\n", s.Name())
	m.screens[s.Name()] = s

}

func (m *Manager) ActiveScreen(name string) error {
	l, ok := m.screens[name]

	if !ok {
		panic("could not find screen: " + name)
	}

	m.Current = l
	return l.Init()
}
