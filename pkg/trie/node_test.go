package trie

import (
	"github.com/intenvy/memoir/pkg/key"
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildSampleShallowNode() *trieNode {
	node := newTrieNode('t')
	node.forceInitChild('a')
	node.forceInitChild('b')
	return node
}

func Test_newTrieNode(t *testing.T) {
	expected := trieNode{
		children:     make(map[rune]*trieNode, 0),
		symbol:       'e',
		value:        0,
		root:         false,
		endOfKey:     false,
		pathFromRoot: "",
	}
	actual := *newTrieNode('e')
	assert.Equal(t, expected, actual)
}

func Test_trieNode_forceInitChild(t *testing.T) {
	expected := map[rune]*trieNode{'a': newTrieNode('a'), 'b': newTrieNode('b')}
	actual := buildSampleShallowNode().children
	assert.Equal(t, expected, actual)
}

func Test_trieNode_hasChildren_True(t *testing.T) {
	assert.True(t, buildSampleShallowNode().hasChildren())
}

func Test_trieNode_hasChildren_False(t *testing.T) {
	assert.False(t, newTrieNode('a').hasChildren())
}

func Test_trieNode_noOfChildren(t *testing.T) {
	expected := 2
	actual := buildSampleShallowNode().noOfChildren()
	assert.Equal(t, expected, actual)
}

func Test_trieNode_valueOfSelectorChild_NoSelectorChild(t *testing.T) {
	expected := 0
	actual := buildSampleShallowNode().valueOfSelectorChild()
	assert.Equal(t, expected, actual)
}

func Test_trieNode_valueOfSelectorChild_WithSelectorChild(t *testing.T) {
	node := buildSampleShallowNode()
	node.forceInitChild(key.SelectorChar)
	node.children[key.SelectorChar].value = 10
	expected := 10
	actual := node.valueOfSelectorChild()
	assert.Equal(t, expected, actual)
}
