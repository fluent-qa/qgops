package path_trie

// PathTrie visitor.  Each method is applied to each PathTrie node and
// non-nil value.  Returning false aborts the traversal.
type Visitor interface {
	VisitNode(key string, node *PathTrieNode) bool
	VisitVal(val *interface{}) bool
}

// Default visitor that performs a nop traversal.  Override to implement
// a subset of the visitor methods.
type DefaultVisitor struct{}
func (*DefaultVisitor) VisitNode(key string, node *PathTrieNode) bool {
	return true
}
func (*DefaultVisitor) VisitVal(val *interface{}) bool {
	return true
}

// Apply visitor to this PathTrie.  Returns true if traversal finishes,
// false otherwise.
func (pt *PathTrie) Apply(visitor Visitor) bool {
	for k, v := range pt.Trie {
		if !v.apply(k, visitor) {
			return false
		}
	}
	return true
}

func (node *PathTrieNode) apply(key string, visitor Visitor) bool {
	if !visitor.VisitNode(key, node) {
		return false
	}
	if !visitor.VisitVal(&node.Val) {
		return false
	}
	if node.Children == nil {
		return true
	}
	for k, v := range node.Children {
		if !v.apply(k, visitor) {
			return false
		}
	}
	return true
}
