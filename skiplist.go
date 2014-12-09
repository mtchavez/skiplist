package skiplist

import "math/rand"

const (
	// ListMaxLevel is the Skiplist can have
	ListMaxLevel = 32
	// ListP is the P value for the SkipList
	ListP = 0.25
)

type SkipList interface {
	Search(key int) *Node
	Delete(key int) bool
	Insert(key int, val []byte) (bool, error)
}

type List struct {
	Length   int
	MaxLevel int
	level    int
	header   *Node
}

// New initializes a new skiplist with
// max level of 32 or 2^32 elements
func New() *List {
	return NewWithLevel(ListMaxLevel)
}

// NewWithLevel initializes a new skiplist with a custom
// max level. Level is defaulted to 32 to allow
// for 2^32 max elements
func NewWithLevel(level int) *List {
	return &List{
		MaxLevel: level,
		header:   &Node{forward: make([]*Node, level)},
	}
}

// Returns a random level used during inserting nodes
func (l *List) randomLevel() int {
	newLevel := 1
	for rand.Float64() >= ListP {
		newLevel++
	}
	if newLevel < l.MaxLevel {
		return newLevel
	}
	return l.MaxLevel
}
