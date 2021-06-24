package trie

import (
	"fmt"
	"github.com/intenvy/memoir/pkg"
	"github.com/intenvy/memoir/pkg/key"
	"sync"
)

type Repository struct {
	rw        sync.RWMutex
	root      *trieNode
	size      int
	converter key.Converter
	validator key.Validator
}

func New() *Repository {
	return &Repository{
		size:      0,
		root:      &trieNode{root: true, children: make(map[rune]*trieNode)},
		converter: key.NewConverterPipeline(),
		validator: key.NewValidatorPipeline(),
	}
}

var _ pkg.KeyValueRepository = (*Repository)(nil)

func (t *Repository) AddConverter(converter key.Converter) *Repository {
	t.converter = converter
	return t
}

func (t *Repository) AddValidator(validator key.Validator) *Repository {
	t.validator = validator
	return t
}

// walks the trie along the entry characters
// if a node is not present, it force creates it
// if the entry is a selector, then it only walks
// until one rune is left
// returns the node it ended up on
func (t *Repository) forceWalk(entry *key.Key) (lastNode *trieNode, pathExists bool) {
	var (
		iter       = t.root
		isSelector = entry.IsSelector()
		pathSize   = entry.Size()
	)
	for idx, symbol := range *entry {
		if isSelector && idx == pathSize-1 {
			break
		}
		iter.forceInitChild(symbol)
		iter = iter.children[symbol]
	}
	return iter, iter.endOfKey
}

// walks the trie along the entry characters
// if a node is not present, it returns
// if the entry is a selector, then it only walks until one rune is left
// returns the last node of the walk
// and if the path exists or not
// and if the whole path has been walked or not
func (t *Repository) lazyWalk(entry *key.Key) (lastNode *trieNode, pathExists bool, completeWalk bool) {
	var (
		iter       = t.root
		isSelector = entry.IsSelector()
		pathSize   = entry.Size()
	)
	for idx, symbol := range *entry {
		if isSelector && idx == pathSize-1 {
			break
		}
		if _, hasChild := iter.children[symbol]; !hasChild {
			return iter, false, false
		}
		iter = iter.children[symbol]
	}
	return iter, iter.endOfKey, true
}

func (t *Repository) Insert(pattern string, value int) error {
	entry := key.New(pattern, t.converter, t.validator)
	if entry.IsSelector() {
		return key.NewErrSelectorKeyNotAllowed(*entry)
	}
	// entry is a raw key
	t.rw.Lock()
	defer t.rw.Unlock()
	node, keyExists := t.forceWalk(entry)
	if !keyExists {
		// entry is a new key
		t.size++
		node.endOfKey = true
		node.pathFromRoot = string(*entry)
		node.value = value
		return nil
	}
	// the key already existed
	return key.NewErrKeyAlreadyExist(*entry)
}

func dfsFillMap(tn *trieNode, out map[string]int) {
	if tn.hasChildren() {
		for _, node := range tn.children {
			if node.endOfKey {
				out[node.pathFromRoot] = node.value
			}
			dfsFillMap(node, out)
		}
	}
}

func (t *Repository) GetMap(pattern string) map[string]int {
	entry := key.New(pattern, t.converter, t.validator)
	results := make(map[string]int)
	t.rw.RLock()
	defer t.rw.RUnlock()
	node, pathExists, completeWalk := t.lazyWalk(entry)
	if !completeWalk {
		return make(map[string]int)
	}
	if pathExists {
		results[node.pathFromRoot] = node.value
	}
	if entry.IsSelector() {
		dfsFillMap(node, results)
	}
	return results
}

func dfsGetValue(tn *trieNode, carry int) int {
	result := 0
	carry += tn.valueOfSelectorChild()
	if tn.endOfKey {
		result += tn.value + carry
	}
	for _, child := range tn.children {
		if child.symbol != key.SelectorChar {
			result += dfsGetValue(child, carry)
		}
	}
	return result
}

func (t *Repository) GetValue(pattern string) int {
	entry := key.New(pattern, t.converter, t.validator)
	t.rw.RLock()
	defer t.rw.RUnlock()
	node, _, completeWalk := t.lazyWalk(entry)
	if !completeWalk {
		return 0
	}
	if entry.IsSelector() {
		return dfsGetValue(node, 0)
	} else {
		return node.valueOfSelectorChild() + node.value
	}
}

func (t *Repository) Inc(pattern string) error {
	entry := key.New(pattern, t.converter, t.validator)
	t.rw.Lock()
	defer t.rw.Unlock()
	node, pathExists, completeWalk := t.lazyWalk(entry)
	if !completeWalk {
		return key.NewErrKeyNotFound(*entry)
	}
	if entry.IsSelector() {
		//       / * increment here
		//  node - child
		//       \ child
		node.forceInitChild(key.SelectorChar)
		child := node.children[key.SelectorChar]
		child.value++
		child.endOfKey = true
		child.pathFromRoot = pattern
		if !pathExists {
			// no child, no path
			return key.NewErrKeyNotFound(*entry)
		}
		return nil
	} else if !pathExists {
		return key.NewErrKeyNotFound(*entry)
	} else {
		node.value++
	}
	return nil
}

func (t *Repository) Contains(pattern string) bool {
	entry := key.New(pattern, t.converter, t.validator)
	t.rw.RLock()
	defer t.rw.RUnlock()
	node, pathExists, completeWalk := t.lazyWalk(entry)
	if !completeWalk {
		return false
	}
	if pathExists {
		return true
	}
	if entry.IsSelector() && node.hasChildren() {
		//        / *
		//  node  - child
		//        \ child
		return true
	}
	return false
}

func (t *Repository) Size() int {
	t.rw.RLock()
	defer t.rw.RUnlock()
	return t.size
}

func (t *Repository) Print() {
	printTrie(t.root, "", true)
}

func printTrie(node *trieNode, indent string, isLastChild bool) {
	fmt.Print(indent)
	if isLastChild {
		fmt.Print(`┗━`)
		indent += "  "
	} else {
		fmt.Print("┟━")
		indent += "┃ "
	}
	fmt.Println(" "+string(node.symbol)+":", node.value)

	i := 0
	for _, child := range node.children {
		printTrie(child, indent, i == len(node.children)-1)
		i++
	}
}
