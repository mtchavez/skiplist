package skiplist

import "testing"

func TestNext(t *testing.T) {
	l := New()
	i := l.Iterator()
	if i.Next() {
		t.Errorf("Should not have nodes to iterate for new list")
	}
	l.Insert(2, "Node two")
	i.Next()
	l.Insert(3, "Node three")
	i.Next()
	if i.Val() != "Node three" {
		t.Errorf("Didn't iterate to next node")
	}
	if i.Key() != 3 {
		t.Errorf("Got wrong next node")
	}
}

func TestPrevious(t *testing.T) {
	l := New()
	i := l.Iterator()
	if i.Prev() {
		t.Errorf("Should not find a val if no node")
	}
	l.Insert(2, "Node two")
	i.Next()

	l.Insert(3, "Node three")
	i.Next()

	if !i.Prev() {
		t.Errorf("Should have a previous node")
	}
}

func TestValueNoCurrent(t *testing.T) {
	i := &iterable{}
	if i.Val() != "" {
		t.Errorf("Should not find a val if no node")
	}
}

func TestKeyNoCurrent(t *testing.T) {
	i := &iterable{}
	if i.Key() != 0 {
		t.Errorf("Should not find a key if no node")
	}
}
