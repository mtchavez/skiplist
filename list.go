package skiplist

import "math/rand"

type List struct {
	MaxLevel int
	header   *node
	footer   *node
	level    int
	length   int
}

const (
	ListMaxLevel = 32
	ListP        = 0.25
)

// Initialize a new skiplist with
// max level of 2^32
func New() *List {
	return NewWithLevel(ListMaxLevel)
}

// Initialize a new skiplist with a custom
// max level. Level is set as 2^(level)
// and is defaulted to 2^32
func NewWithLevel(level int) *List {
	return &List{
		MaxLevel: level,
		header:   &node{forward: make([]*node, level)},
	}
}

// Returns length of list
func (l *List) Len() int {
	return l.length
}

// skiplist Iterator used for convenience to traverse
// through skip list. Implements the Iterator interface.
//
// Example:
//		for i: = list.Iterator(); i.Next; {
//			fmt.Println(i.Key(), i.Val())
//		}
//
func (l *List) Iterator() Iterator {
	return &iterable{curr: l.header}
}

// Search for a node in the skiplist by key
// Returns node value, if present, and a boolean
// if the node was found or not
func (l *List) Search(searchKey int) (string, bool) {
	x := l.header
	for i := len(x.forward) - 1; i >= 0; i-- {
		for x.forward[i] != nil &&
			x.forward[i].key < searchKey {
			x = x.forward[i]
		}
	}

	x = x.forward[0]
	if x != nil && x.key == searchKey {
		return x.val, true
	} else {
		return "", false
	}
}

// Convenience min function to avoid importing math
func minVal(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Returns a random level used during inserting nodes
func (l *List) randomLevel() int {
	newLevel := 1
	for rand.Float64() >= ListP {
		newLevel += 1
	}
	return minVal(newLevel, l.MaxLevel)
}

// Insert a new node into the skiplist
// Note: Allows duplicate keys
func (l *List) Insert(newKey int, newVal string) {
	update := make([]*node, l.MaxLevel)
	x := l.header
	for i := len(x.forward) - 1; i >= 0; i-- {
		for x.forward[i] != nil &&
			(x.forward[i].key < newKey ||
				(x.forward[i].key == newKey && x.forward[i].val == newVal)) {
			x = x.forward[i]
		}

		if update != nil {
			update[i] = x
		}
	}
	newLevel := l.randomLevel()
	if newLevel > l.level {
		for i := l.level + 1; i <= newLevel; i++ {
			update = append(update, l.header)
		}
		l.level = newLevel
	}
	x = NewNode(newLevel+1, newKey, newVal)
	for i := 0; i <= newLevel; i++ {
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
	l.length++
	if l.footer == nil || l.footer.key < newKey {
		l.footer = x
	}
}
