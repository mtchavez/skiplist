package skiplist

// Split takes a key to split a list by
// All values less than the provided key
// will be in the new list which will be returned
func (l *List) Split(splitKey int) *List {
	level := ListMaxLevel
	if l.level > level {
		level = l.level
	}
	newList := NewListWithLevel(level)
	x := l.header
	for i := l.level; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < splitKey {
			x = x.forward[i]
		}
		if x.forward[i] != nil {
			newList.header.forward[i] = x.forward[i]
			x.forward[i] = nil
		}
	}
	for l.header.forward[l.level] == nil && l.level > 0 {
		l.level--
	}
	lx := l.header
	listLength := 0
	for lx.forward != nil && lx.forward[0] != nil {
		listLength++
		lx = lx.forward[0]
	}

	for newList.header.forward[newList.level] == nil && newList.level > 0 {
		newList.level--
	}
	newList.length = l.length - listLength
	l.length = listLength
	return newList
}
