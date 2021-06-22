package pkg

// Interface of a Key(String) to Value(int) repository
type KeyValueRepository interface {

	// inserts "key" directly and sets the value
	// if key already exists, returns error
	Insert(key string, value int) error

	// inserts "key" directly and sets value to zero
	// if key already exists, returns error
	Make(key string) error

	// returns a map from all keys that match "pattern" to all values
	GetMap(pattern string) map[string]int

	// returns the total sum of values of keys
	// that match the given "pattern"
	GetValue(pattern string) int

	// Increments all keys that match the given "pattern"
	Inc(pattern string) error

	// checks if repository contains the exact "key"
	ContainsKey(key string) bool

	// checks if repository contains any keys
	// that match given "pattern"
	ContainsPattern(pattern string) bool

	// will remove the given "key" and its value
	// if no such key exists, return error
	Evict(key string) error

	// sets the key converter
	// any pattern or key will be converted using it,
	// and then will be processed
	SetKeyConverter(converter *KeyConverter)
}
