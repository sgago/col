package slice

func IndexOf[T comparable](slice []T, value T) (int, error) {
	index, _, e := First(slice, func(i int, v T) bool {
		return value == v
	})

	return index, e
}
