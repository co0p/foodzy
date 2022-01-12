package ecs

import (
	"fmt"
	"log"
)

type Entity struct {
	Tag        string
	Active     bool
	components map[ComponentType]ComponentTyper
}

func NewEntity(tag string, active bool) *Entity {
	return &Entity{
		Tag:        tag,
		Active:     active,
		components: make(map[ComponentType]ComponentTyper),
	}
}

// HasComponent returns true if the Entity has the gameplay of type cType associated
func (e *Entity) HasComponent(cType ComponentType) bool {
	_, ok := e.components[cType]
	return ok
}

// GetComponent returns the gameplay, panics if not found
func (e *Entity) GetComponent(cType ComponentType) ComponentTyper {
	if c, ok := e.components[cType]; !ok {
		panic(fmt.Sprintf("expected Entity to have component of type %s attached", cType))
	} else {
		return c
	}
}

// AddComponent adds a gameplay to the Entity
func (e *Entity) AddComponent(c ComponentTyper) {
	if e.HasComponent(c.Type()) {
		log.Printf("Entity already has component of type %s attached", c)
	}
	e.components[c.Type()] = c
}

func (e *Entity) AddComponents(c ...ComponentTyper) {
	for _, v := range c {
		e.AddComponent(v)
	}
}

// RemoveComponent removes the associated gameplay from the Entity
func (e *Entity) RemoveComponent(c ComponentType) {
	delete(e.components, c)
}

func (e *Entity) String() string {
	return fmt.Sprintf("active: %v, components: %d", e.Active, len(e.components))
}
