package components

type Collision struct {
	Width  int
	Height int
}

func (c *Collision) ID() string {
	return "collider"
}
