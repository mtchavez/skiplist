package skiplist

import (
	"reflect"
	"sync"
)

// DupeList implements the SkipList interface
// but allows for duplicate keys to be inserted
type DupeList struct {
	sync.RWMutex
	MaxLevel int
	level    int
	length   int
	header   *Node
	footer   *Node
}

var _ SkipList = (*DupeList)(nil)

// NewDupeList initializes a new skiplist with
// max level of 32 or 2^32 elements that allows duplicates
func NewDupeList() *DupeList {
	return NewDupeListWithLevel(ListMaxLevel)
}

// NewDupeListWithLevel initializes a new skiplist with a custom
// max level. Level is defaulted to 32 to allow
// for 2^32 max elements
func NewDupeListWithLevel(level int) *DupeList {
	return &DupeList{
		MaxLevel: level,
		header:   &Node{forward: make([]*Node, level)},
		level:    0,
	}
}

// Iterator returns an iterable from current list header
func (l *DupeList) Iterator() Iterator {
	return &iterable{curr: l.header}
}

// Search by key to find the matching node in the list
// For duplicate keys this will always return the last
// inserted node for the key
func (l *DupeList) Search(key int) *Node {
	l.RLock()
	defer l.RUnlock()
	x := l.header
	for i := l.level; i >= 0; i-- {
		for i < len(x.forward) && x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		return x
	}
	return nil
}

// SearchKeyVal allows you to search for nodes by the key and
// value of the node which is useful for nodes with duplicate keys
func (l *DupeList) SearchKeyVal(key int, val []byte) *Node {
	l.RLock()
	defer l.RUnlock()
	x := l.header
	for i := l.level; i >= 0; i-- {
		for i < len(x.forward) && x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil {
		for x != nil && x.key == key {
			if reflect.DeepEqual(x.val, val) {
				return x
			}
			if x.forward != nil {
				x = x.forward[0]
			} else {
				x = nil
			}
		}
	}
	return nil
}

// Insert a node into the list given a key and a byte array value
func (l *DupeList) Insert(key int, val interface{}) *Node {
	update := make([]*Node, l.MaxLevel)
	x := l.header
	l.Lock()
	defer l.Unlock()

	for i := l.level; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	newLevel := randomLevel(l.MaxLevel)
	if newLevel > l.level {
		for i := l.level + 1; i < newLevel; i++ {
			update[i] = l.header
		}
		l.level = newLevel
	}
	x = NewNode(newLevel, key, val)
	for i := 0; i < newLevel; i++ {
		x.forward[i] = update[i].forward[i]
		update[i].forward[i] = x
	}
	x.backward = nil
	if update[0] != l.header {
		x.backward = update[0]
	}
	if x.forward[0] != nil {
		x.forward[0].backward = x
	}
	if l.footer == nil || l.footer.key < key {
		l.footer = x
	}
	l.length++
	return x
}

// Delete a node by the key provided
func (l *DupeList) Delete(key int) bool {
	update := make([]*Node, l.MaxLevel)
	x := l.header
	l.Lock()
	defer l.Unlock()

	for i := l.level; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < key {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		for i := 0; i < l.level; i++ {
			if update[i].forward[i] != x {
				break
			}
			update[i].forward[i] = x.forward[i]
		}
		for l.level > 1 && len(l.header.forward) > l.level && l.header.forward[l.level-1] == nil {
			l.level--
		}
		l.length--
		return true
	}
	return false
}
