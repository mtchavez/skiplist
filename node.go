package skiplist

// Internal skiplist node
//
// Implements a node for a skiplist which is a linked list
// node containing a slice of forward referencing nodes,
// the previous node and the key and value of the node
type Node struct {
	forward  []*Node
	backward *Node
	key      int
	val      []byte
}

// NewNode takes a level used for the forward slice
// referencing linked nodes as well as the key and
// value of the node
func NewNode(level, key int, val []byte) *Node {
	return &Node{
		forward: make([]*Node, level),
		key:     key,
		val:     val,
	}
}

// Returns next node or nil if next node is not present
func (n *Node) next() *Node {
	if len(n.forward) == 0 {
		return nil
	}
	return n.forward[0]
}

// Returns the previous node
func (n *Node) prev() *Node {
	return n.backward
}
