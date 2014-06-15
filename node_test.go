package skiplist

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode(2, 2, "Node two")
	if len(n.forward) != 2 {
		t.Errorf("Forward length should be set to 2")
	}
	if n.key != 2 {
		t.Errorf("Key should be set to 2")
	}
	if n.val != "Node two" {
		t.Errorf("Val should be set to 'Node two'")
	}
}

func TestNextWithoutForwardNodes(t *testing.T) {
	n := NewNode(0, 2, "Node two")
	if n.next() != nil {
		t.Errorf("Should not have found forward nodes")
	}
}

func TestNextWithForwardNodes(t *testing.T) {
	n := NewNode(2, 2, "Node two")
	n2 := NewNode(2, 3, "Node three")
	n.forward[0] = n2
	if n.next() != n2 {
		t.Errorf("Should have found n2 as a forward node")
	}
}

func TestPrev(t *testing.T) {
	n := NewNode(2, 2, "Node two")
	if n.prev() != nil {
		t.Errorf("Should not have found a previous node")
	}
	n2 := NewNode(2, 3, "Node three")
	n.backward = n2
	if n.prev() != n2 {
		t.Errorf("Should have found n2 as the previous node")
	}
}
