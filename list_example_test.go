package skiplist

import (
	"fmt"
	"reflect"
)

func ExampleNewList() {
	l := NewList()
	fmt.Println("Max level:", l.MaxLevel)
	// Output:
	// Max level: 32
}

func ExampleNewListWithLevel() {
	l := NewListWithLevel(200)
	fmt.Println("Max level:", l.MaxLevel)
	// Output:
	// Max level: 200
}

func ExampleList_Insert() {
	l := NewList()
	l.Insert(1, []byte("example 1"))
	l.Insert(2, []byte("example 2"))
	l.Insert(3, []byte("example 3"))
	l.Insert(4, []byte("example 4"))
	fmt.Println("Size:", l.Size())
	// Output:
	// Size: 4
}

func ExampleList_Search() {
	l := NewList()
	l.Insert(1, []byte("example 1"))
	l.Insert(2, []byte("example 2"))
	l.Insert(3, []byte("example 3"))
	l.Insert(4, []byte("example 4"))

	// Not Found
	notFound := l.Search(45)
	found := l.Search(4)

	fmt.Printf("Searched for 45 and got %+v\n", notFound)
	fmt.Printf("Searched for 4 and got '%+v'\n", string(reflect.ValueOf(found.Value()).Bytes()))
	// Output:
	// Searched for 45 and got <nil>
	// Searched for 4 and got 'example 4'
}

func ExampleList_Delete() {
	l := NewList()
	l.Insert(1, []byte("example 1"))
	l.Insert(2, []byte("example 2"))
	l.Insert(3, []byte("example 3"))
	l.Insert(4, []byte("example 4"))

	found := l.Search(4)
	fmt.Printf("Searched for 4 and got '%+v'\n", string(reflect.ValueOf(found.Value()).Bytes()))

	// Delete node
	fmt.Printf("Deleted key 4? %+v\n", l.Delete(4))

	notFound := l.Search(4)
	fmt.Printf("Searched for deleted key of 4 and got %+v\n", notFound)
	// Output:
	// Searched for 4 and got 'example 4'
	// Deleted key 4? true
	// Searched for deleted key of 4 and got <nil>
}
