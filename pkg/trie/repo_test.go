package trie

import (
	"errors"
	"github.com/intenvy/memoir/pkg/key"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func buildDefaultTrie() *Repository {
	converter := key.NewConverterPipeline().AddHook(
		func(str string) string {
			return strings.ToLower(str)
		},
	)
	validator := key.NewValidatorPipeline().AddHook(
		func(str string) error {
			if strings.Contains(str, " ") {
				return errors.New("spacing is not allowed")
			}
			return nil
		},
	)
	repo := New().AddConverter(converter).AddValidator(validator)
	return repo
}

func buildTrieFromTokens(value int, tokens ...string) *Repository {
	repo := buildDefaultTrie()
	for _, token := range tokens {
		_ = repo.Insert(token, value)
	}
	return repo
}

func TestRepository_AddConverter(t *testing.T) {
	repo := New()
	expected := repo
	actual := repo.AddConverter(key.NewConverterPipeline())
	assert.Equal(t, expected, actual)
}

func TestRepository_AddValidator(t *testing.T) {
	repo := New()
	expected := repo
	actual := repo.AddValidator(key.NewValidatorPipeline())
	assert.Equal(t, expected, actual)
}

func TestRepository_Contains_StrictKey(t *testing.T) {
	notInRepo := []string{"Ab", "abc", "dab", "zop", "abcdeb"}
	inRepo := []string{"abcDaa", "abCdab", "aBcdee", "AbcdeA"}
	repo := buildTrieFromTokens(100, inRepo...)
	for _, token := range inRepo {
		assert.True(t, repo.Contains(token))
	}
	for _, token := range notInRepo {
		assert.False(t, repo.Contains(token))
	}
}

func TestRepository_GetMap_AllSelector(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("t1", 1)
	_ = repo.Insert("t2", 2)
	_ = repo.Insert("t3", 3)
	_ = repo.Insert("t11", 11)
	_ = repo.Insert("t12", 12)
	_ = repo.Inc("t*")
	expected := map[string]int{
		"t1":  1,
		"t2":  2,
		"t3":  3,
		"t11": 11,
		"t12": 12,
		"t*":  1,
	}
	actual := repo.GetMap("*")
	assert.Equal(t, expected, actual)
}

func TestRepository_GetMap_SubTrieSelector(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("t1", 1)
	_ = repo.Insert("t2", 2)
	_ = repo.Insert("t3", 3)
	_ = repo.Insert("t11", 11)
	_ = repo.Insert("t12", 12)
	_ = repo.Inc("t1*")
	expected := map[string]int{
		"t1":  1,
		"t1*": 1,
		"t11": 11,
		"t12": 12,
	}
	actual := repo.GetMap("t1*")
	assert.Equal(t, expected, actual)
}

func TestRepository_GetMap_ExistingStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("t1", 1)
	_ = repo.Insert("t2", 2)
	_ = repo.Insert("t3", 3)
	_ = repo.Insert("t11", 11)
	_ = repo.Insert("t12", 12)
	_ = repo.Inc("t*")
	expected := map[string]int{"t1": 1}
	actual := repo.GetMap("t1")
	assert.Equal(t, expected, actual)
}

func TestRepository_GetMap_NonExistingStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("t1", 1)
	_ = repo.Insert("t2", 2)
	_ = repo.Insert("t3", 3)
	_ = repo.Insert("t11", 11)
	_ = repo.Insert("t12", 12)
	_ = repo.Inc("t*")
	expected := map[string]int{}
	actual := repo.GetMap("t")
	assert.Equal(t, expected, actual)
}

func TestRepository_GetValue_AllSelector(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("a", 10)
	_ = repo.Insert("aa", 10)
	_ = repo.Insert("ab", 10)
	_ = repo.Insert("ac", 10)
	_ = repo.Insert("aaa", 10)
	_ = repo.Insert("aab", 10)
	_ = repo.Insert("aac", 10)
	_ = repo.Inc("a*")
	_ = repo.Inc("aa*")
	expected := 81
	actual := repo.GetValue("*")
	assert.Equal(t, expected, actual)
}

func TestRepository_Inc_ExistingStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	keys := []string{"abc", "def", "abd"}
	insertedValue := 1
	expectedAfterInc := 2
	for _, k := range keys {
		_ = repo.Insert(k, insertedValue)
		_ = repo.Inc(k)
		assert.Equal(t, expectedAfterInc, repo.GetValue(k))
	}
}

func TestRepository_Inc_NonExistingStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	keys := []string{"abc", "def", "abd"}
	for _, k := range keys {
		assert.Error(t, repo.Inc(k))
	}
}

func TestRepository_Inc_Selector(t *testing.T) {
	repo := buildDefaultTrie()
	keys := []string{"abc", "def", "abd"}
	insertedValue := 1
	expectedAfterInc := 2
	for _, k := range keys {
		_ = repo.Insert(k, insertedValue)
		_ = repo.Inc(k)
		assert.Equal(t, expectedAfterInc, repo.GetValue(k))
	}
}

func TestRepository_Insert_SingleStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	err := repo.Insert("ali", 3)
	expected := 3
	actual := repo.GetValue("ali")
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestRepository_Insert_MultiStrictKey(t *testing.T) {
	repo := buildDefaultTrie()
	insertedValue := 3
	for _, pattern := range []string{"aaa", "aa", "ab", "abc", "baa", "bba"} {
		err := repo.Insert(pattern, insertedValue)
		expected := insertedValue
		actual := repo.GetValue(pattern)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestRepository_Insert_RepeatedKey(t *testing.T) {
	repo := buildDefaultTrie()
	_ = repo.Insert("test", 3)
	assert.Error(t, repo.Insert("test", 3))
}

func TestRepository_Insert_Selector(t *testing.T) {
	assert.Error(t, buildDefaultTrie().Insert("test/*", 3))
}

func TestRepository_Size_UniqueKeys(t *testing.T) {
	repo := buildDefaultTrie()
	for _, pattern := range []string{"aaa", "aa", "ab", "abc", "baa", "bba"} {
		_ = repo.Insert(pattern, 0)
	}
	expected := 6
	actual := repo.Size()
	assert.Equal(t, expected, actual)
}

func TestRepository_Size_NoSideEffectAfterInc(t *testing.T) {
	repo := buildTrieFromTokens(0,"aaa", "aa", "ab", "abc", "baa", "bba")
	_ = repo.Inc("*")
	_ = repo.Inc("aa*")
	_ = repo.Inc("a")
	_ = repo.Inc("ab")
	_ = repo.Inc("a*")
	expected := 6
	actual := repo.Size()
	assert.Equal(t, expected, actual)
}

func TestRepository_forceWalk_NonExistingKey(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern"
	node, patternExists := repo.forceWalk(key.New(pattern, repo.converter, repo.validator))
	expected := rune(pattern[len(pattern)-1])
	actual := node.symbol
	assert.Equal(t, expected, actual)
	assert.False(t, patternExists)
}

func TestRepository_forceWalk_NonExistingSelector(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern/*"
	_ = repo.Insert(pattern[0:len(pattern)-2], 1)
	node, patternExists := repo.forceWalk(key.New(pattern, repo.converter, repo.validator))
	expected := rune(pattern[len(pattern)-2])
	actual := node.symbol
	assert.Equal(t, expected, actual)
	assert.False(t, patternExists)
}

func TestRepository_forceWalk_ExistingKey(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern"
	_ = repo.Insert(pattern, 100)
	node, patternExists := repo.forceWalk(key.New(pattern, repo.converter, repo.validator))
	expected := rune(pattern[len(pattern)-1])
	actual := node.symbol
	assert.Equal(t, expected, actual)
	assert.True(t, patternExists)
}

func TestRepository_forceWalk_ExistingSelector(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern/*"
	_ = repo.Insert("pattern/", 10)
	_ = repo.Inc(pattern)
	node, patternExists := repo.forceWalk(key.New(pattern, repo.converter, repo.validator))
	expected := rune(pattern[len(pattern)-2])
	actual := node.symbol
	assert.Equal(t, expected, actual)
	assert.True(t, patternExists)
}

func TestRepository_lazyWalk_NonExistingKey(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern"
	node, patternExists, completePath := repo.lazyWalk(key.New(pattern, repo.converter, repo.validator))
	expected := repo.root
	actual := node
	assert.Equal(t, expected, actual)
	assert.False(t, patternExists)
	assert.False(t, completePath)
}

func TestRepository_lazyWalk_NonExistingSelector(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern/*"
	_ = repo.Insert(pattern[0:2], 1)
	node, patternExists, completePath := repo.lazyWalk(key.New(pattern, repo.converter, repo.validator))
	expected := node.symbol
	actual := rune(pattern[1])
	assert.Equal(t, expected, actual)
	assert.False(t, patternExists)
	assert.False(t, completePath)
}

func TestRepository_lazyWalk_ExistingKey(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern"
	_ = repo.Insert(pattern, 100)
	node, patternExists, completePath := repo.lazyWalk(key.New(pattern, repo.converter, repo.validator))
	expected := rune(pattern[len(pattern)-1])
	actual := node.symbol
	assert.Equal(t, expected, actual)
	assert.True(t, patternExists)
	assert.True(t, completePath)
}

func TestRepository_lazyWalk_ExistingSelector(t *testing.T) {
	repo := buildDefaultTrie()
	pattern := "pattern/*"
	_ = repo.Insert(pattern[0:len(pattern)-1], 100)
	node, patternExists, completePath := repo.lazyWalk(key.New(pattern, repo.converter, repo.validator))
	expected := string(pattern[len(pattern)-2])
	actual := string(node.symbol)
	assert.Equal(t, expected, actual)
	assert.True(t, patternExists)
	assert.True(t, completePath)
}
