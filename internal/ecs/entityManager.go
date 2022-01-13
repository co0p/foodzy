package ecs

type EntityManager struct {
	entities []*Entity
}

// AddEntity adds the given Entity to the internal list of entities
func (m *EntityManager) AddEntity(e *Entity) {
	m.entities = append(m.entities, e)
}

// QueryByComponents returns the set of Entity having all provided components
func (m *EntityManager) QueryByComponents(types ...ComponentType) []*Entity {

	var candidates []*Entity

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

func (m *EntityManager) QueryByTag(tag string) []*Entity {
	var candidates []*Entity
	for _, e := range m.entities {
		if e.Tag == tag {
			candidates = append(candidates, e)
		}
	}

	return candidates
}

func (m *EntityManager) QueryFirstByTag(tag string) *Entity {
	for _, e := range m.entities {
		if e.Tag == tag {
			return e
		}
	}

	panic("expected to find at least one entity with tag:" + tag)
	return nil
}

func (m *EntityManager) RemoveInactive() {
	cleaned := m.entities[:0]

	for _, entity := range m.entities {
		if entity.Active {
			cleaned = append(cleaned, entity)
		}
	}

	m.entities = cleaned
}

func (m *EntityManager) Clear() {
	m.entities = []*Entity{}
}
