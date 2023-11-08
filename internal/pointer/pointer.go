package pointer

func To[T any](a T) *T {
	return &a
}
