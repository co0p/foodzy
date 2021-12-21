package component

type Velocity struct {
	X float64
	Y float64
}

const VelocityType ComponentType = "Velocity"

func (v Velocity) Type() ComponentType {
	return VelocityType
}
