package component

type Transform struct {
	X        float64
	Y        float64
	Z        float64
	Scale    float64
	Rotation float64
}

const TransformType ComponentType = "Transform"

func (p Transform) Type() ComponentType {
	return TransformType
}
