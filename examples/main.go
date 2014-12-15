package main

import (
	"fmt"

	"github.com/mtchavez/skiplist"
)

func main() {
	list := skiplist.NewList()
	found := list.Search(3)
	fmt.Printf("List has 3? %+v\n", found)
	list.Insert(10, []byte("ten"))
	list.Insert(3, []byte("three"))
	list.Insert(2, []byte("two"))
	list.Insert(2, []byte("two two"))
	list.Insert(12, []byte("twelve"))

	found = list.Search(3)
	fmt.Printf("List has 3? %+v\n", found)
	found = list.Search(2)
	fmt.Printf("List has 2? %+v\n", found)
	for i := list.Iterator(); i.Next(); {
		fmt.Println(i.Key(), i.Val())
	}
}
