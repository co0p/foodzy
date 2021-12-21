package component

type Position struct {
	X float64
	Y float64
}

const PositionType ComponentType = "Position"

func (p Position) Type() ComponentType {
	return PositionType
}
