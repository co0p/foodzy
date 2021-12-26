package screen

import "log"

type Manager struct {
	activateChannel chan Name
	Current         Screen
	screens         map[Name]Screen
}

func NewManager() *Manager {

	manager := &Manager{
		activateChannel: make(chan Name),
		Current:         nil,
		screens:         make(map[Name]Screen),
	}

	go manager.ListenOnActive()

	return manager
}

func (m *Manager) AddScreen(s Screen) {
	log.Printf("[screen.Manager] activating screen:%s\n", s.Name())
	m.screens[s.Name()] = s

}

func (m *Manager) ActiveScreen(name Name) {
	m.activateChannel <- name
}

func (m *Manager) ListenOnActive() {

	name := <-m.activateChannel
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
