package entity

import (
	"fmt"
	"github.com/co0p/foodzy/component"
)

type Entity struct {
	Tag        string
	Active     bool
	components map[component.ComponentType]component.ComponentTyper
}

func NewEntity(tag string, active bool) *Entity {
	return &Entity{
		Tag:        tag,
		Active:     active,
		components: make(map[component.ComponentType]component.ComponentTyper),
	}
}

// HasComponent returns true if the Entity has the component of type cType associated
func (e *Entity) HasComponent(cType component.ComponentType) bool {
	_, ok := e.components[cType]
	return ok
}

// GetComponent returns the component, panics if not found
func (e *Entity) GetComponent(cType component.ComponentType) component.ComponentTyper {
	if c, ok := e.components[cType]; !ok {
		panic(fmt.Sprintf("expected Entity to have component of type %s attached", cType))
	} else {
		return c
	}
}

// AddComponent adds a component to the Entity, panics if a component of the same type already exists
func (e *Entity) AddComponent(c component.ComponentTyper) {
	if e.HasComponent(c.Type()) {
		panic(fmt.Sprintf("Entity already has component of type %s attached", c))
	}
	e.components[c.Type()] = c
}

// RemoveComponent removes the associated component from the Entity
func (e *Entity) RemoveComponent(c component.ComponentTyper) {
	delete(e.components, c.Type())
}

func (e *Entity) String() string {
	return fmt.Sprintf("active: %v, components: %d", e.Active, len(e.components))
}
