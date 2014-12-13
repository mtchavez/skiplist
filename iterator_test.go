package skiplist

import (
	"bytes"
	"testing"
)

func TestNext(t *testing.T) {
	l := NewList()
	i := l.Iterator()
	if i.Next() {
		t.Errorf("Should not have nodes to iterate for new list")
	}
	l.Insert(2, []byte("Node two"))
	i.Next()
	l.Insert(3, []byte("Node three"))
	i.Next()
	if !bytes.Equal(i.Val(), []byte("Node three")) {
		t.Errorf("Didn't iterate to next node")
	}
	if i.Key() != 3 {
		t.Errorf("Got wrong next node")
	}
}

func TestPrevious(t *testing.T) {
	l := NewList()
	i := l.Iterator()
	if i.Prev() {
		t.Errorf("Should not find a val if no node")
	}
	l.Insert(2, []byte("Node two"))
	i.Next()

	l.Insert(3, []byte("Node three"))
	i.Next()

	if !i.Prev() {
		t.Errorf("Should have a previous node")
	}
}

func TestValueNoCurrent(t *testing.T) {
	i := &iterable{}
	if i.Val() != nil {
		t.Errorf("Should not find a val if no node")
	}
}

func TestKeyNoCurrent(t *testing.T) {
	i := &iterable{}
	if i.Key() != 0 {
		t.Errorf("Should not find a key if no node")
	}
}
