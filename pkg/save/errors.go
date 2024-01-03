package save

import "fmt"

type ErrMissingKey struct {
	Map, Key string
}

func NewErrMissingKey(m, key string) error {
	return &ErrMissingKey{m, key}
}

func (e *ErrMissingKey) Error() string {
	return fmt.Sprintf("The map \"%s\" has no key \"%s\"", e.Map, e.Key)
}
