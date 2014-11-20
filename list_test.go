package skiplist

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	if l.MaxLevel != ListMaxLevel {
		t.Errorf("Default max level should be %d", ListMaxLevel)
	}
	if l.level != 0 {
		t.Errorf("Default level should be 0")
	}
	if l.header == nil {
		t.Errorf("Header should exist for new List")
	}
	if len(l.header.forward) != ListMaxLevel {
		t.Errorf("Header forward size should be size of list max level")
	}
}

func TestNewWithLevel(t *testing.T) {
	l := NewWithLevel(64)
	if l.MaxLevel == ListMaxLevel {
		t.Errorf("Should not be Default max level")
	}
	if l.MaxLevel != 64 {
		t.Errorf("Should be able to set List level")
	}
}

func TestLen(t *testing.T) {
	l := New()
	if l.Len() != 0 {
		t.Errorf("Default list lenght should be 0")
	}
	l.Insert(2, "two")
	if l.Len() != 1 {
		t.Errorf("Should increment list size to 1")
	}
}

func TestIterator(t *testing.T) {
	l := New()
	i := l.Iterator()
	if i.Prev() {
		t.Errorf("Iterator should start with List header")
	}
}

func TestSearchNotFound(t *testing.T) {
	l := New()
	_, found := l.Search(35)
	if found {
		t.Errorf("Should not have found a value in an empty List")
	}
}

func TestSearchFound(t *testing.T) {
	l := New()
	l.Insert(35, "My value")
	val, found := l.Search(35)
	if !found {
		t.Errorf("Should have found a value for 35")
	}
	if val != "My value" {
		t.Errorf("Value should have been 'My value'")
	}
}

func TestSearchKeyValNotFound(t *testing.T) {
	l := New()
	found := l.SearchKeyVal(35, "Nope")
	if found {
		t.Errorf("Should not have found a value in an empty List")
	}
}

func TestSearchKeyValFound(t *testing.T) {
	l := New()
	l.Insert(35, "My value")
	found := l.SearchKeyVal(35, "My not value")
	if found {
		t.Errorf("Value should match when searching")
	}

	found = l.SearchKeyVal(35, "My value")
	if !found {
		t.Errorf("Should have found a value for 35")
	}
}

func TestSearchKeyValMultipleValues(t *testing.T) {
	l := New()
	l.Insert(35, "My value")
	l.Insert(35, "Another value")

	found := l.SearchKeyVal(35, "My value")
	if !found {
		t.Errorf("Should have found 'My Value' for key 35")
	}
}

func TestMinVal(t *testing.T) {
	if minVal(2, 3) != 2 {
		t.Errorf("Min value should have been 2")
	}
}

func TestRandomLevel(t *testing.T) {
	l := New()
	l.MaxLevel = 10
	if l.randomLevel() == 10 {
		t.Errorf("Should chose the minimum")
	}
	l.MaxLevel = 0
	if l.randomLevel() != 0 {
		t.Errorf("Should not change level if already minimum")
	}
}

func TestInsert(t *testing.T) {
	l := New()
	if l.length != 0 {
		t.Errorf("Should have 0 length before inserting nodes")
	}
	l.Insert(2, "Node two")
	if l.length != 1 {
		t.Errorf("Length should be 1 after inserting once")
	}
	l.Insert(3, "Node three")
	if l.length != 2 {
		t.Errorf("Length should be 2 after inserting twice")
	}
}

func TestInsertUpdateLevel(t *testing.T) {
	l := New()
	l.level = 0
	l.Insert(2, "Node two")
	if l.level == 0 {
		t.Errorf("Should have updated List level")
	}
	if l.length != 1 {
		t.Errorf("Length should be 1 after inserting once")
	}
}

func TestInsertDupeKeys(t *testing.T) {
	l := New()
	l.Insert(2, "Node two")
	if l.length != 1 {
		t.Errorf("Length should be 1 after inserting once")
	}
	l.Insert(2, "Node three")
	if l.length != 2 {
		t.Errorf("Length should be 2 after inserting twice")
	}
}
