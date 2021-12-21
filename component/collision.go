package component

type Collision struct {
	Width  float64
	Height float64
}

const CollisionType ComponentType = "Collision"

func (c *Collision) Type() ComponentType {
	return CollisionType
}
