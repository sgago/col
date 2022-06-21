package err

import "fmt"

type KeyNotFound struct {
	Key int
}

func (e *KeyNotFound) Error() string {
	return fmt.Sprintf("Key %d not found.", e.Key)
}

type NotFound struct{}

func (e *NotFound) Error() string {
	return "Value not found."
}
