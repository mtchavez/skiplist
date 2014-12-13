package skiplist

func (l *List) Split(splitKey int) *List {
	level := ListMaxLevel
	if l.level > level {
		level = l.level
	}
	newList := NewListWithLevel(level)
	newList.length = 0
	x := l.header
	for i := len(x.forward) - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].key < splitKey {
			x = x.forward[i]
		}
		if x.forward[i] != nil {
			newList.header.forward[i] = x.forward[i]
			x.forward[i] = nil
			newList.length++
		}
	}
	for l.header.forward[l.level] == nil && l.level > 0 {
		l.level--
	}

	for newList.header.forward[newList.level] == nil && newList.level > 0 {
		newList.level--
	}
	l.length = l.length - newList.length
	return newList
}
