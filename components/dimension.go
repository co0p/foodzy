package components

type Dimension struct {
	Width  float64
	Height float64
}

func (d Dimension) ID() string {
	return ""
}
