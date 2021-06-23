package pkg

// Interface of a Key(String) to Value(int) repository
type KeyValueRepository interface {

	// inserts "pattern" directly and sets the value
	// if pattern already exists, returns error
	Insert(pattern string, value int) error

	// returns a map from all keys that match "pattern" to all values
	GetMap(pattern string) map[string]int

	// returns the total sum of values of keys
	// that match the given "pattern"
	GetValue(pattern string) int

	// increments all keys
	// that match the given "pattern"
	Inc(pattern string) error

	// checks if repository contains any keys
	// that match given "pattern"
	Contains(pattern string) bool

	// returns the number of keys
	// that are present in the trie
	Size() int
}
