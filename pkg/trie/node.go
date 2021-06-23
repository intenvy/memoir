package trie

import "github.com/intenvy/memoir/pkg/key"

type trieNode struct {
	children     map[rune]*trieNode
	symbol       rune
	value        int
	root         bool
	endOfKey     bool
	pathFromRoot string
}

func newTrieNode(symbol rune) *trieNode {
	return &trieNode{
		symbol:   symbol,
		children: make(map[rune]*trieNode),
	}
}

func (tn *trieNode) forceInitChild(symbol rune) {
	if _, hasChild := tn.children[symbol]; !hasChild {
		tn.children[symbol] = newTrieNode(symbol)
	}
}

func (tn *trieNode) noOfChildren() int {
	return len(tn.children)
}

func (tn *trieNode) hasChildren() bool {
	return tn.noOfChildren() > 0
}

func (tn *trieNode) valueOfSelectorChild() int {
	if _, hasSelectorChild := tn.children[key.SelectorChar]; hasSelectorChild {
		return tn.children[key.SelectorChar].value
	}
	return 0
}
