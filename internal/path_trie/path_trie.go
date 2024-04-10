package path_trie

import (
	"strings"
)

type PathTrieNode struct {
	Children PathTrieMap `json:"children,omitempty"`

	// Number of nodes in this subtree with non-null vals.  Includes this node.
	Count int `json:"count"`

	// Name of the path segment corresponding to this node.  E.g. if this node
	// represents /v1/foo/bar, the name would be "bar" (and the prefix would
	// be "/v1/foo/bar").
	Name string `json:"name"`

	// The prefix includes the node's name and uniquely identifies the node in
	// the tree.
	Prefix string `json:"prefix"`

	// Payload for the node.
	Val interface{} `json:"val"`
}

type PathTrieMap map[string]*PathTrieNode

type PathTrie struct {
	Trie          PathTrieMap `json:"trie"`
	PathSeparator string `json:"path_separator"`
}

type ValueMergeFunc func(existing, new *interface{})

// Create a PathTrie with "/" as the path separator.
func New() PathTrie {
	return PathTrie{
		Trie:          make(PathTrieMap, 0),
		PathSeparator: "/",
	}
}

// Create a PathTrie with a user-supplied path separator.
func NewWithPathSeparator(pathSeparator string) PathTrie {
	return PathTrie{
		Trie:          make(PathTrieMap, 0),
		PathSeparator: pathSeparator,
	}
}

// Insert val at path, with path segments separated by PathSeparator.
// Returns true if a new node was created, false if an existing node
// was overwritten.
//
// If path starts with PathSeparator, the first PathSeparator is disregarded,
// e.g. "/foo/bar" does not cause the key in the top-level map to be "".
func (pt *PathTrie) Insert(path string, val interface{}) bool {
	return pt.InsertMerge(path, val, func (existing, new *interface{}) {
		*existing = *new
	})
}

// Insert val at path, with path segments separated by PathSeparator.
// Returns true if a new node was created, false if an existing node
// was overwritten.
//
// The merge function is responsible for updating the existing value
// with the new value.
func (pt *PathTrie) InsertMerge(path string, val interface{}, merge ValueMergeFunc) bool {
	trie := pt.Trie
	nodesToUpdate := make([]*PathTrieNode, 0)
	createdNewNode := true
	segments := strings.Split(path, pt.PathSeparator)

	// Traverse the Trie along path, inserting nodes where necessary.
	for idx, segment := range segments {
		isLastElt := idx == len(segments) - 1
		if node, ok := trie[segment]; ok {
			if isLastElt {
				// If this is the last path segment, then this is the node to update.
				merge(&node.Val, &val)
				createdNewNode = false
			} else {
				// Otherwise, continue descending.
				nodesToUpdate = append(nodesToUpdate, node)
				trie = node.Children
			}
		} else {
			node := PathTrieNode{
				Children: make(PathTrieMap, 0),
				Count: 1,
				Name: segment,
				Prefix: strings.Join(segments[:idx + 1], pt.PathSeparator),
			}
			if isLastElt {
				node.Val = val
			}
			trie[segment] = &node
			trie = node.Children
		}
	}

	// If a new node was created, increment the count of its existing
	// ancestors.
	if createdNewNode {
		for _, node := range nodesToUpdate {
			node.Count += 1
		}
	}

	return createdNewNode
}

type mergeVisitor struct {
	DefaultVisitor
	target *PathTrie
	merge ValueMergeFunc
}

func (v *mergeVisitor) VisitNode(key string, node *PathTrieNode) bool {
	if node.Val == nil {
		return true
	}

	v.target.InsertMerge(node.Prefix, node.Val, v.merge)
	return true
}

// Merge another PathTrie into this one.  Values along the same paths
// are merged using the merge function.
func (pt *PathTrie) Merge(other *PathTrie, merge ValueMergeFunc) {
	vis := mergeVisitor{
		target:         pt,
		merge:          merge,
	}
	other.Apply(&vis)
}
