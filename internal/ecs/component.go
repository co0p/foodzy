package ecs

type ComponentType string

type ComponentTyper interface {
	Type() ComponentType
}
