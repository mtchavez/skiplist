package skiplist

import (
	"bytes"
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

func TestSearchNotFound(t *testing.T) {
	l := New()
	found := l.Search(35)
	if found != nil {
		t.Errorf("Should not have found a value in an empty List")
	}
}

func TestSearchFound(t *testing.T) {
	l := New()
	l.Insert(35, []byte("My value"))
	found := l.Search(35)
	if found == nil {
		t.Errorf("Should have found a node for 35")
	}
	if !bytes.Equal(found.val, []byte("My value")) {
		t.Errorf("Value should have been 'My value'")
	}
}

func TestInsert(t *testing.T) {
	l := New()
	if l.length != 0 {
		t.Errorf("Should have 0 length before inserting nodes")
	}
	l.Insert(2, []byte("Node two"))
	if l.length != 1 {
		t.Errorf("Length should be 1 after inserting once")
	}
	l.Insert(3, []byte("Node three"))
	if l.length != 2 {
		t.Errorf("Length should be 2 after inserting twice")
	}
}

func TestInsertUpdateLevel(t *testing.T) {
	l := New()
	l.level = 0
	l.Insert(2, []byte("Node two"))
	if l.level == 0 {
		t.Errorf("Should have updated List level")
	}
	if l.length != 1 {
		t.Errorf("Length should be 1 after inserting once")
	}
}

func TestDelete(t *testing.T) {
	l := New()
	l.Insert(1, []byte("one"))
	l.Insert(2, []byte("two"))
	l.Insert(3, []byte("three"))

	l.Delete(1)
	if l.length != 2 {
		t.Errorf("Length should be 2 after one delete but got %v", l.length)
	}
	x := l.Search(1)
	if x != nil {
		t.Errorf("Expected to not find deleted node but got %+v", x)
	}

	l.Delete(3)
	if l.length != 1 {
		t.Errorf("Length should be 1 after two deletes but got %v", l.length)
	}
	x = l.Search(3)
	if x != nil {
		t.Errorf("Expected to not find deleted node but got %+v", x)
	}
}
