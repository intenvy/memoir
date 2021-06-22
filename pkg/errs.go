package pkg

import "fmt"

// error that is returned from the repository
// when a key is required and is not found
type ErrKeyNotFound struct {
	key string
}

var _ error = (*ErrKeyNotFound)(nil)

func NewErrKeyNotFound(key string) ErrKeyNotFound {
	return ErrKeyNotFound{key: key}
}

func (e *ErrKeyNotFound) Error() string {
	return fmt.Sprintf(`key: "%s" not found`, e.key)
}

// error that is returned from the repository
// when a key already exists in the repository
type ErrKeyAlreadyExists struct {
	key string
}

var _ error = (*ErrKeyAlreadyExists)(nil)

func NewErrKeyAlreadyExist(key string) ErrKeyAlreadyExists {
	return ErrKeyAlreadyExists{key: key}
}

func (e *ErrKeyAlreadyExists) Error() string {
	return fmt.Sprintf(`key: "%s" already exists`, e.key)
}
