package skiplist

type DupeList struct {
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

func (l *DupeList) Iterator() Iterator {
	return &iterable{curr: l.header}
}

func (l *DupeList) Search(key int) *Node {
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

func (l *DupeList) Insert(key int, val []byte) *Node {
	update := make([]*Node, l.MaxLevel)
	x := l.header
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

func (l *DupeList) Delete(key int) bool {
	update := make([]*Node, l.MaxLevel)
	x := l.header
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
