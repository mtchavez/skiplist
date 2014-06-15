package main

import (
	"fmt"

	"github.com/mtchavez/skiplist"
)

func main() {
	list := skiplist.New()
	s, f := list.Search(3)
	fmt.Printf("Searching for 3: %+v %+v\n", s, f)
	list.Insert(10, "ten")
	list.Insert(3, "three")
	list.Insert(2, "two")
	list.Insert(2, "two two")
	list.Insert(12, "twelve")

	s, f = list.Search(3)
	fmt.Printf("Searching for 3: %+v %+v\n", s, f)
	s, f = list.Search(2)
	fmt.Printf("Searching for 2: %+v %+v\n", s, f)
	s, f = list.Search(10)
	fmt.Printf("Searching for 10: %+v %+v\n", s, f)
	s, f = list.Search(1)
	fmt.Printf("Searching for 1: %+v %+v\n", s, f)
	fmt.Println("Length: ", list.Len())
	for i := list.Iterator(); i.Next(); {
		fmt.Println(i.Key(), i.Val())
	}
}
