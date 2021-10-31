package components

type Collision struct {
	Width  float64
	Height float64
}

func (c *Collision) ID() string {
	return "collider"
}
