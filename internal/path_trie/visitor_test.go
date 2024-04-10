package path_trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyVisitor struct {
	DefaultVisitor
	prefixes []string
}

func (v *MyVisitor) VisitNode(key string, node *PathTrieNode) bool {
	v.prefixes = append(v.prefixes, node.Prefix)
	return true
}

func TestVisitor(t *testing.T) {
	val := 1
	tests := []struct {
		input    PathTrie
		expected []string
	}{
		{
			// "/foo",
			PathTrie{
				Trie: PathTrieMap{
					"foo": &PathTrieNode{
						Children: nil,
						Count:    1,
						Name:     "foo",
						Prefix:   "/foo",
						Val:      val,
					},
				},
				PathSeparator: "/",
			},
			[]string{"/foo"},
		},
		{
			// "foo",
			PathTrie{
				Trie: PathTrieMap{
					"foo": &PathTrieNode{
						Children: nil,
						Count:    1,
						Name:     "foo",
						Prefix:   "foo",
						Val:      val,
					},
				},
				PathSeparator: "/",
			},
			[]string{"foo"},
		},
		{
			// "/foo/{bar}",
			PathTrie{
				Trie: PathTrieMap{
					"foo": &PathTrieNode{
						Children: PathTrieMap{
							"{bar}": &PathTrieNode{
								Children: nil,
								Count:    1,
								Name:     "{bar}",
								Prefix:   "/foo/{bar}",
								Val:      val,
							},
						},
						Count:  1,
						Name:   "foo",
						Prefix: "/foo",
					},
				},
				PathSeparator: "/",
			},
			[]string{"/foo", "/foo/{bar}"},
		},
	}

	for _, test := range tests {
		vis := MyVisitor{}
		test.input.Apply(&vis)
		assert.Equal(t, test.expected, vis.prefixes)
	}
}