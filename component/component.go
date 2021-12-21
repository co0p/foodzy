package component

type ComponentType string

type ComponentTyper interface {
	Type() ComponentType
}
