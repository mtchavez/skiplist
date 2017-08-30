package skiplist

import (
	"math/rand"
	"sync"
)

const (
	// ListMaxLevel is the Skiplist can have
	ListMaxLevel = 32
	// ListP is the P value for the SkipList
	ListP = 0.5
)

// SkipList interface defining the methods
// needed for a skip list
type SkipList interface {
	Search(key int) *Node
	Delete(key int) bool
	Insert(key int, val []byte) *Node
	Iterator() Iterator
}

// List is a basic skip list implementation
// with search, delete, and insert functionality
// and uses a mutex to allow for multi-threaded interaction
type List struct {
	sync.RWMutex
	MaxLevel int
	level    int
	length   int
	header   *Node
	footer   *Node
}

var _ SkipList = (*List)(nil)

// Returns a random level used during inserting nodes
func randomLevel(maxLevel int) int {
	newLevel := 1
	for rand.Float64() >= ListP && newLevel < maxLevel && newLevel < ListMaxLevel-1 {
		newLevel++
	}
	return newLevel
}

// NewList initializes a new skiplist with
// max level of 32 or 2^32 elements
func NewList() *List {
	return NewListWithLevel(ListMaxLevel)
}

// NewListWithLevel initializes a new skiplist with a custom
// max level. Level is defaulted to 32 to allow
// for 2^32 max elements
func NewListWithLevel(level int) *List {
	return &List{
		MaxLevel: level,
		header:   &Node{forward: make([]*Node, level)},
		level:    0,
	}
}

// Iterator returns an iterable from the current
// head of the skiplist.
func (l *List) Iterator() Iterator {
	return &iterable{curr: l.header}
}

// Size returns the length of the list
func (l *List) Size() int {
	return l.length
}

// Search for a node in the skip list by the key
// will return a Node if found or nil if not found
func (l *List) Search(key int) *Node {
	l.RLock()
	defer l.RUnlock()
	x := l.header
	for i := l.level; i >= 0; i-- {
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

// Insert a new node into the skip list providing a
// integer key and a byte array value. Will return
// the inserted Node
func (l *List) Insert(key int, val []byte) *Node {
	update := make([]*Node, l.MaxLevel)
	x := l.header
	var alreadyChecked *Node
	l.Lock()
	defer l.Unlock()

	for i := l.level; i >= 0; i-- {
		for x.forward[i] != nil &&
			alreadyChecked != x.forward[i] &&
			x.forward[i].key < key {
			x = x.forward[i]
		}
		alreadyChecked = x.forward[i]
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.key == key {
		x.val = val
		return x
	}
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

// Delete will delete a node for the provided key
// will return a true/false if Node was deleted or not
func (l *List) Delete(key int) bool {
	update := make([]*Node, l.MaxLevel)
	x := l.header
	var alreadyChecked *Node
	l.Lock()
	defer l.Unlock()

	for i := l.level; i >= 0; i-- {
		for x.forward[i] != nil &&
			alreadyChecked != x.forward[i] &&
			x.forward[i].key < key {
			x = x.forward[i]
		}
		alreadyChecked = x.forward[i]
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
