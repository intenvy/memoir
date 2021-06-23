package key

import "fmt"

// error that is returned from the repository
// when a key is required and is not found
type ErrKeyNotFound struct {
	key Key
}

var _ error = (*ErrKeyNotFound)(nil)

func NewErrKeyNotFound(key Key) *ErrKeyNotFound {
	return &ErrKeyNotFound{key: key}
}

func (e *ErrKeyNotFound) Error() string {
	return fmt.Sprintf(`key: "%s" not found`, e.key)
}

// error that is returned from the repository
// when a key already exists in the repository
type ErrKeyAlreadyExists struct {
	key Key
}

var _ error = (*ErrKeyAlreadyExists)(nil)

func NewErrKeyAlreadyExist(key Key) *ErrKeyAlreadyExists {
	return &ErrKeyAlreadyExists{key: key}
}

func (e *ErrKeyAlreadyExists) Error() string {
	return fmt.Sprintf(`key: "%s" already exists`, e.key)
}

// error that is returned from the repository
// when a selector key cannot be accepted as input
type ErrSelectorKeyNotAllowed struct {
	key Key
}

var _ error = (*ErrSelectorKeyNotAllowed)(nil)

func NewErrSelectorKeyNotAllowed(key Key) *ErrSelectorKeyNotAllowed {
	return &ErrSelectorKeyNotAllowed{key: key}
}

func (e *ErrSelectorKeyNotAllowed) Error() string {
	return fmt.Sprintf(`key: "%s" is a selector and is not allowed`, e.key)
}
