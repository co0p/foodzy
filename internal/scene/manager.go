package scene

type Action string

type SceneManager struct {
	Current Scene
	scenes  map[string]Scene
}

func NewSceneManager() *SceneManager {

	manager := &SceneManager{
		Current: nil,
		scenes:  make(map[string]Scene),
	}

	return manager
}

func (m *SceneManager) AddScene(s Scene) {
	m.scenes[s.Name()] = s
}

func (m *SceneManager) Activate(name string) {

	scene, ok := m.scenes[name]

	if !ok {
		panic("could not find scene: " + name)
	}

	if m.Current != nil {
		m.Current.Stop()
	}
	m.Current = scene
	scene.Start()
}
