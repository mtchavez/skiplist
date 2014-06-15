package skiplist

type Iterator interface {
	Next() (ok bool)
	Prev() (ok bool)
	Val() string
	Key() int
}

type iterable struct {
	curr *node
	key  int
	val  string
}

func (i *iterable) Next() bool {
	next := i.curr.next()
	if next == nil {
		return false
	}
	i.curr = next
	return true
}

func (i *iterable) Prev() bool {
	prev := i.curr.prev()
	if prev == nil {
		return false
	}
	i.curr = prev
	return true
}

func (i *iterable) Val() string {
	if i.curr == nil {
		return ""
	}
	return i.curr.val
}

func (i *iterable) Key() int {
	if i.curr == nil {
		return 0
	}
	return i.curr.key
}
