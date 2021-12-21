package entity

import "github.com/co0p/foodzy/component"

type Manager struct {
	entities []*entity
}

// NewEntity creates a new entity and adds it to the internal list of entity
func (m *Manager) AddEntity(e *entity) {
	m.entities = append(m.entities, e)
}

// QueryByComponents returns the set of entity having all provided components
func (m *Manager) QueryByComponents(types ...component.ComponentType) []*entity {

	var candidates []*entity

	for _, e := range m.entities {
		matchCount := 0
		for _, c := range types {
			if e.HasComponent(c) {
				matchCount++
			}
		}

		if matchCount == len(types) {
			candidates = append(candidates, e)
		}
	}
	return candidates
}

func (m *Manager) QueryByTag(tag string) []*entity {
	var candidates []*entity
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