package path_trie

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func wrap(m PathTrieMap, count int) PathTrieMap {
	return PathTrieMap{
		"": &PathTrieNode{
			Children: m,
			Count:    count,
			Name:     "",
			Prefix:   "",
		},
	}
}

func TestInsert(t *testing.T) {
	val := 1
	emptyMap := make(PathTrieMap, 0)
	tests := []struct {
		input    []string
		expected PathTrie
	}{
		{
			[]string{"/foo"},
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: emptyMap,
						Count:    1,
						Name:     "foo",
						Prefix:   "/foo",
						Val:      val,
					},
				}, 1),
				PathSeparator: "/",
			},
		},
		{
			[]string{"/foo/"},
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: PathTrieMap{
							"": &PathTrieNode{
								Children: emptyMap,
								Count: 1,
								Name: "",
								Prefix: "/foo/",
								Val: val,
							},
						},
						Count:    1,
						Name:     "foo",
						Prefix:   "/foo",
					},
				}, 1),
				PathSeparator: "/",
			},
		},
		{
			[]string{"foo"},
			PathTrie{
				Trie: PathTrieMap{
					"foo": &PathTrieNode{
						Children: emptyMap,
						Count:    1,
						Name:     "foo",
						Prefix:   "foo",
						Val:      val,
					},
				},
				PathSeparator: "/",
			},
		},
		{
			[]string{"/foo/{bar}"},
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: PathTrieMap{
							"{bar}": &PathTrieNode{
								Children: emptyMap,
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
				}, 1),
				PathSeparator: "/",
			},
		},
		{
			[]string{"/foo", "/foo/{bar}"},
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: PathTrieMap{
							"{bar}": &PathTrieNode{
								Children: emptyMap,
								Count:    1,
								Name:     "{bar}",
								Prefix:   "/foo/{bar}",
								Val:      val,
							},
						},
						Count:  2,
						Name:   "foo",
						Prefix: "/foo",
						Val:    val,
					},
				}, 2),
				PathSeparator: "/",
			},
		},
	}

	for _, test := range tests {
		trie := New()
		for _, s := range test.input {
			nodeAdded := trie.Insert(s, val)
			assert.True(t, nodeAdded, strings.Join(test.input, ", "))
		}
		assert.Equal(t, test.expected, trie, strings.Join(test.input, ", "))
	}
}

func TestInsertMerge(t *testing.T) {
	emptyMap := make(PathTrieMap, 0)
	val := 1
	tests := []struct {
		input    string
		expected PathTrie
	}{
		{
			"/foo",
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: emptyMap,
						Count:    1,
						Name:     "foo",
						Prefix:   "/foo",
						Val:      val * 2,
					},
				}, 1),
				PathSeparator: "/",
			}},
		{
			"foo",
			PathTrie{
				Trie: PathTrieMap{
					"foo": &PathTrieNode{
						Children: emptyMap,
						Count:    1,
						Name:     "foo",
						Prefix:   "foo",
						Val:      val * 2,
					},
				},
				PathSeparator: "/",
			}},
		{
			"/foo/{bar}",
			PathTrie{
				Trie: wrap(PathTrieMap{
					"foo": &PathTrieNode{
						Children: PathTrieMap{
							"{bar}": &PathTrieNode{
								Children: emptyMap,
								Count:    1,
								Name:     "{bar}",
								Prefix:   "/foo/{bar}",
								Val:      val * 2,
							},
						},
						Count:  1,
						Name:   "foo",
						Prefix: "/foo",
					},
				}, 1),
				PathSeparator: "/",
			}},
	}

	for _, test := range tests {
		trie := New()
		merge := func(old, new *interface{}) {
			left := (*old).(int)
			right := (*new).(int)
			*old = left + right
		}

		nodeAdded := trie.InsertMerge(test.input, val, merge)
		assert.True(t, nodeAdded, test.input)

		nodeAdded = trie.InsertMerge(test.input, val, merge)
		assert.False(t, nodeAdded, test.input)

		assert.Equal(t, test.expected, trie, test.input)
	}
}

func TestMerge(t *testing.T) {
	tests := [][]string{
		{
			"/foo/bar",
			"/foo/xxx/yyy",
		},
		{
			"/foo/bar",
			"/foo/bar",
		},
	}

	for _, test := range tests {
		left := New()
		left.Insert(test[0], 1)

		right := New()
		right.Insert(test[1], 1)

		expected := New()
		expected.Insert(test[0], 1)
		expected.Insert(test[1], 1)

		left.Merge(&right, func(left, right *interface{}) {
			*left = *right
		})

		assert.Equal(t, expected, left)
	}
}
