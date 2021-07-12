package entities

import "github.com/co0p/foodzy/components"

type Manager struct {
	entities []*Entity
}

// AddEntity adds an entity to the manager
func (m *Manager) AddEntity(e *Entity) {
	m.entities = append(m.entities, e)
}

// Query returns the set of entities having all provided components
func (m *Manager) Query(components ...components.Component) []*Entity {

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
