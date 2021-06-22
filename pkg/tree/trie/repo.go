package trie

import "github.com/intenvy/memoir/pkg"

type Repository struct {
	// TODO
}

func New() *Repository {
	// TODO
	return &Repository{}
}

func (t Repository) Insert(key string, value int) error {
	// TODO
	panic("implement me")
}

func (t Repository) Make(key string) error {
	// TODO
	panic("implement me")
}

func (t Repository) GetMap(pattern string) map[string]int {
	// TODO
	panic("implement me")
}

func (t Repository) GetValue(pattern string) int {
	// TODO
	panic("implement me")
}

func (t Repository) Inc(pattern string) error {
	// TODO
	panic("implement me")
}

func (t Repository) ContainsKey(key string) bool {
	// TODO
	panic("implement me")
}

func (t Repository) ContainsPattern(pattern string) bool {
	// TODO
	panic("implement me")
}

func (t Repository) Evict(key string) error {
	// TODO
	panic("implement me")
}

func (t Repository) SetKeyConverter(converter *pkg.KeyConverter) {
	// TODO
	panic("implement me")
}
