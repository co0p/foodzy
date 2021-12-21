package component

type KeyboardMover struct {
	Speed float64
}

const KeyboardMoverType ComponentType = "KeyboardMover"

func (k *KeyboardMover) Type() ComponentType {
	return KeyboardMoverType
}
