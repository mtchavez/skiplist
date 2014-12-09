package skiplist

import "testing"

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
