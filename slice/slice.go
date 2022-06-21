package slice

import "collections/err"

func First[T any](slice []T) T {
	return slice[0]
}

func Last[T any](slice []T) T {
	return slice[len(slice)-1]
}

func Clear[T any](slice []T) []T {
	return make([]T, 0, cap(slice))
}

func RemoveFirst[T any](slice []T) []T {
	return slice[1:]
}

func RemoveLast[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

func Swap[T any](slice []T, indexA int, indexB int) []T {
	temp := slice[indexA]
	slice[indexA] = slice[indexB]
	slice[indexB] = temp

	return slice
}

func Contains[T comparable](slice []T, value T) bool {
	for _, val := range slice {
		if value == val {
			return true
		}
	}

	return false
}

func IndexOf[T comparable](slice []T, value T) (int, error) {
	for index, val := range slice {
		if value == val {
			return index, nil
		}
	}

	return -1, &err.NotFound{}
}
