package radix

import "github.com/intenvy/memoir/pkg"

type Repository struct {
	// TODO
}

func NewRepository() *Repository {
	// TODO
	return &Repository{}
}

var _ pkg.KeyValueRepository = (*Repository)(nil)

func (r Repository) Insert(key string, value int) error {
	// TODO
	panic("implement me")
}

func (r Repository) Make(key string) error {
	// TODO
	panic("implement me")
}

func (r Repository) GetMap(pattern string) map[string]int {
	// TODO
	panic("implement me")
}

func (r Repository) GetValue(pattern string) int {
	// TODO
	panic("implement me")
}

func (r Repository) Inc(pattern string) error {
	// TODO
	panic("implement me")
}

func (r Repository) ContainsKey(key string) bool {
	// TODO
	panic("implement me")
}

func (r Repository) ContainsPattern(pattern string) bool {
	// TODO
	panic("implement me")
}

func (r Repository) Evict(key string) error {
	// TODO
	panic("implement me")
}

func (r Repository) SetKeyConverter(converter *pkg.KeyConverter) {
	// TODO
	panic("implement me")
}
