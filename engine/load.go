package engine

type Loader[T any] interface {
	Load() (T, error)
}

func MustLoad[T any](l Loader[T]) T {
	value, err := l.Load()
	if err != nil {
		panic(err)
	}
	return value
}
