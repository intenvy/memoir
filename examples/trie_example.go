package examples

import (
	"fmt"
	"github.com/intenvy/memoir/pkg/key"
	"github.com/intenvy/memoir/pkg/trie"
)

func buildRawTrie() *trie.Repository{
	return trie.New().
		AddValidator(key.NewValidatorPipeline()).
		AddConverter(key.NewConverterPipeline())
}

func InsertionExample() {
	repo := buildRawTrie()
	_ = repo.Insert("home/", 1)
	_ = repo.Insert("home/user/", 1)
	_ = repo.Insert("home/bin/", 1)
	_ = repo.Insert("home/bin/tar", 1)
	_ = repo.Insert("root/etc", 1)

	repo.Print()
}

func IncExample() {
	repo := buildRawTrie()
	_ = repo.Insert("home/", 1)
	_ = repo.Insert("home/user/", 1)
	_ = repo.Insert("home/bin/", 1)
	_ = repo.Insert("home/bin/tar", 1)
	_ = repo.Insert("root/etc", 1)
	_ = repo.Insert("root/*", 1)
	_ = repo.Insert("home/bin/*", 1)

	repo.Print()
}

func GetValueExample() {
	repo := buildRawTrie()
	_ = repo.Insert("home/", 1)
	_ = repo.Insert("home/user/", 1)
	_ = repo.Insert("home/bin/", 1)
	_ = repo.Insert("home/bin/tar", 1)
	_ = repo.Insert("root/etc", 1)
	_ = repo.Inc("root/*")
	_ = repo.Inc("home/bin/*")

	for _, pattern := range []string{"*", "home/*", "root/*", "zzz"} {
		fmt.Println("pattern:", pattern, "value:", repo.GetValue(pattern), "map:", repo.GetMap(pattern))
	}

	repo.Print()
}