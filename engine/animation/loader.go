package animation

type Loader interface {
	Load() (*Animation, error)
}

func Load(loader Loader) (*Animation, error) {
	return loader.Load()
}

func MustLoad(loader Loader) (animation *Animation) {
	var err error
	if animation, err = Load(loader); err != nil {
		panic(err)
	}
	return
}
