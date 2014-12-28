package skiplist

import "sync"

// CNode is a node for the concurrent list
type CNode struct {
	sync.RWMutex
	forward     []*CNode
	key         int
	val         []byte
	marked      bool
	fullyLinked bool
	topLayer    int
}
