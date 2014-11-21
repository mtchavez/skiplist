package skiplist

// Internal skiplist node
//
// Implements a node for a skiplist which is a linked list
// node containing a slice of forward referencing nodes,
// the previous node and the key and value of the node
type node struct {
	forward  []*node
	backward *node
	key      int
	val      string
}

// NewNode takes a level used for the forward slice
// referencing linked nodes as well as the key and
// value of the node
func NewNode(level, key int, val string) *node {
	return &node{
		forward: make([]*node, level),
		key:     key,
		val:     val,
	}
}

// Returns next node or nil if next node is not present
func (n *node) next() *node {
	if len(n.forward) == 0 {
		return nil
	}
	return n.forward[0]
}

// Returns the previous node
func (n *node) prev() *node {
	return n.backward
}
