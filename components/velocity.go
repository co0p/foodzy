package components

type Velocity struct {
	X float64
	Y float64
}

func (p Velocity) ID() string {
	return "velocity"
}
