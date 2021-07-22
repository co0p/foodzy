package components

type ObjectType string

const Food ObjectType = "food"
const Plate ObjectType = "plate"

type Collision struct {
	ObjectType   ObjectType
	Width        int
	Height       int
	CollidedWith ObjectType
}

func (c *Collision) ID() string {
	return "collider"
}
