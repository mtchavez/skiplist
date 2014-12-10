package skiplist

import "testing"

func TestSplit(t *testing.T) {
	list1 := New()
	for i := 10; i > 0; i-- {
		list1.Insert(i, []byte{})
	}
	list2 := list1.Split(8)

	if list1.length != 7 {
		t.Errorf("Expected list 1 to be split into 7 nodes but has %+v", list1.length)
	}
	for i := 1; i < list1.length; i++ {
		if list1.Search(i) == nil {
			t.Errorf("Expected list1 to have node for key %+v", i)
		}
	}

	if list2.length != 3 {
		t.Errorf("Expected list 2 to be split into 3 nodes but has %+v", list2.length)
	}
	for i := 1; i < list2.length; i++ {
		if list2.Search(i+7) == nil {
			t.Errorf("Expected list2 to have node for key %+v", i+7)
		}
	}
}
