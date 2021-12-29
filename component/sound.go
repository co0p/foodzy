package component

import "github.com/co0p/foodzy/internal/ecs"

type Sound struct {
	Clip string
}

const SoundType ecs.ComponentType = "Sound"

func (s *Sound) Type() ecs.ComponentType {
	return SoundType
}
