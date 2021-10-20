package entities

import (
	"fmt"
	"github.com/co0p/foodzy/components"
	"reflect"
)

// Entity is the basic build block of the Component based architecture holding the components
type Entity struct {
	Active     bool
	components []components.Component
	Tag        string
}

func NewEntity(tag string, active bool) *Entity {
	return &Entity{
		Tag:    tag,
		Active: active,
	}
}

// AddComponent adds a Component to the Entity; panics if there is a Component of same type already
func (elem *Entity) AddComponent(new components.Component) {
	newType := reflect.TypeOf(new)
	for _, existing := range elem.components {
		if newType == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("attempted to add a Component of type %v that already has been added", new))
		}
	}

	elem.components = append(elem.components, new)
}

// GetComponent gets a Component from the Entity; panics if Component does not exist
func (elem *Entity) GetComponent(withType components.Component) components.Component {
	askingType := reflect.TypeOf(withType)
	for _, existing := range elem.components {
		if askingType == reflect.TypeOf(existing) {
			return existing
		}
	}
	return nil
}

func (e *Entity) String() string {
	return fmt.Sprintf("active: %v, components: %d", e.Active, len(e.components))
}
