package entities

import (
	"github.com/co0p/foodzy/components"
)

type Manager struct {
	entities []*Entity
}

// AddEntity adds an entity to the manager
func (m *Manager) AddEntity(e *Entity) {
	m.entities = append(m.entities, e)
}

// QueryByComponents returns the set of entities having all provided components
func (m *Manager) QueryByComponents(components ...components.Component) []*Entity {

	candidates := []*Entity{}

	for _, entity := range m.entities {

		matchingComponentCount := 0
		for _, c := range components {
			if entity.GetComponent(c) != nil {
				matchingComponentCount++
			}
		}

		if matchingComponentCount == len(components) {
			candidates = append(candidates, entity)
		}
	}

	return candidates
}

func (m *Manager) QueryByTag(tag string) []*Entity {
	candidates := []*Entity{}
	for _, entity := range m.entities {
		if entity.Tag == tag {
			candidates = append(candidates, entity)
		}
	}

	return candidates
}

func (m *Manager) RemoveInactive() {
	cleaned := m.entities[:0]

	for _, entity := range m.entities {
		if entity.Active {
			cleaned = append(cleaned, entity)
		}
	}

	m.entities = cleaned
}
