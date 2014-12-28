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

// CList is a skip list built for high concurrency
type CList struct {
	sync.RWMutex
	MaxLevel int
	level    int
	length   int
	header   *CNode
	footer   *CNode
}

// var _ SkipList = (*CList)(nil)

// NewDupeList initializes a new skiplist with
// max level of 32 or 2^32 elements that allows duplicates
func NewCList() *CList {
	return NewCListWithLevel(ListMaxLevel)
}

// NewDupeListWithLevel initializes a new skiplist with a custom
// max level. Level is defaulted to 32 to allow
// for 2^32 max elements
func NewCListWithLevel(level int) *CList {
	return &CList{
		MaxLevel: level,
		header:   &CNode{forward: make([]*CNode, level)},
		level:    0,
	}
}

// Search will look for a node by the key passed in
// and return a Node if found otherwise nil
func (cl *CList) Search(key int) *CNode {
	x := cl.header
	for i := cl.level; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		return x
	}
	return nil
}
