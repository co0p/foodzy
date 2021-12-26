package component

type MouseCollider struct {
	Width  float64
	Height float64
}

const MouseColliderType ComponentType = "MouseCollider"

func (c MouseCollider) Type() ComponentType {
	return MouseColliderType
}
